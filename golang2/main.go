package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/semconv/v1.19.0/httpconv"
	"go.opentelemetry.io/otel/trace"
)

type HttpWrapper struct {
	operation            string
	serverName           string
	handler              http.Handler
	httpServerDuration   metric.Float64Histogram
	fibonacciInvocations metric.Int64Counter
}

// newRelicProvider creates a new Relic provider
func newRelicProvider(ctx context.Context) *sdktrace.TracerProvider {
	var exp sdktrace.SpanExporter
	var err error

	exp, err = otlptracehttp.New(ctx)
	if err != nil {
		panic(err)
	}

	// Instantiate a default resource with environment variables
	r := resource.Default()

	// Create trace provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)

	// Set global trace provider
	otel.SetTracerProvider(tp)

	// Set trace propagator
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		))

	return tp
}

func shutdownTraceProvider(
	ctx context.Context,
	tp *sdktrace.TracerProvider,
) {
	// Do not make the application hang when it is shutdown.
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	if err := tp.Shutdown(ctx); err != nil {
		panic(err)
	}
}

func initExporter() *otlptrace.Exporter {
	// Create and configure the OTLP exporter to send traces to the collector
	exporter, err := otlptracegrpc.New(context.Background(), otlptracegrpc.WithEndpointURL("http://localhost:4317/api/traces"))
	if err != nil {
		log.Fatalf("failed to create OTLP exporter: %v", err)
		return nil
	}
	// Create a new trace provider with the exporter
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes("", semconv.ServiceNameKey.String("ServiceB"))))
	otel.SetTracerProvider(provider)

	return exporter
}

// HelloHandler is the handler for the /hello route
func Handler(c *gin.Context) {

	r := c.Request
	ctx := r.Context()

	// Set up trace attributes
	startSpanAttributes := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindServer),
		trace.WithAttributes(httpconv.ServerRequest("Vipin'sServer", r)...),
		trace.WithAttributes(semconv.NetHostName("Vipin'sServer")),
	}

	// Start the server span
	ctx, span := otel.GetTracerProvider().
		Tracer("Vipin-Nuwas").
		Start(ctx, r.Method+" /hello", startSpanAttributes...)
	defer span.End()

	endSpanAttributes := []attribute.KeyValue{semconv.HTTPStatusCode(200)}
	span.SetAttributes(endSpanAttributes...)

	// ------------------------------------

	// Create response writer wrapper
	// rww := NewResponseWriterWrapper(w)
	// h.handler.ServeHTTP(rww, r.WithContext(ctx))

	// // Set up metric attributes
	// httpServerMetricAttributes := httpconv.ServerRequest(h.serverName, r)
	// fibonacciInvocationMetricAttributes := []attribute.KeyValue{}

	// if rww.statusCode > 0 {
	// 	// Add status code to metric attributes
	// 	httpServerMetricAttributes = append(
	// 		httpServerMetricAttributes,
	// 		semconv.HTTPStatusCode(rww.statusCode),
	// 	)
	// 	if rww.statusCode == 200 {
	// 		fibonacciInvocationMetricAttributes = append(
	// 			fibonacciInvocationMetricAttributes,
	// 			attribute.Bool("fibonacci.valid.n", true),
	// 		)
	// 	} else {
	// 		fibonacciInvocationMetricAttributes = append(
	// 			fibonacciInvocationMetricAttributes,
	// 			attribute.Bool("fibonacci.valid.n", false),
	// 		)
	// 	}

	// 	// Add status code to span attributes
	// 	endSpanAttributes := []attribute.KeyValue{semconv.HTTPStatusCode(rww.statusCode)}
	// 	span.SetAttributes(endSpanAttributes...)
	// }

	// // Use floating point division here for higher precision (instead of Millisecond method).
	// elapsedTime := float64(time.Since(requestStartTime)) / float64(time.Millisecond)

	// h.fibonacciInvocations.Add(ctx, 1, metric.WithAttributes(fibonacciInvocationMetricAttributes...))
	// h.httpServerDuration.Record(ctx, elapsedTime, metric.WithAttributes(httpServerMetricAttributes...))

	// ------------------------------------
	// Get the tracer from the global provider
	// tracer := otel.GetTracerProvider().Tracer("serviceB")

	// Start a span
	// _, span := tracer.Start(c.Request.Context(), "HelloHandler")

	// Simulate some work
	time.Sleep(time.Second)

	span.AddEvent("handling the request in Service B")
	// Respond with "Hello, World!"
	c.String(http.StatusOK, "Hello from Service B!")
}
func main() {
	// Create and configure the OTLP exporter to send traces to the collector
	//expoter := initExporter()
	// defer expoter.Shutdown(context.Background())

	ctx := context.Background()
	tp := newRelicProvider(ctx)
	defer shutdownTraceProvider(ctx, tp)

	// Create a new Gin router
	r := gin.Default()

	// Define route handlers
	r.GET("/hello", Handler)

	// Start HTTP server
	fmt.Println("Server started on :5001")
	if err := http.ListenAndServe(":5001", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
