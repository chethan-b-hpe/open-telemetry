package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

var URL string
var TraceProvider string

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

func jaegerProvider(ctx context.Context) *sdktrace.TracerProvider {
	// Create and configure the OTLP exporter to send traces to the collector
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create OTLP exporter: %v", err)
	}

	// Create a new trace provider with the exporter
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes("", semconv.ServiceNameKey.String("App2"))))
	otel.SetTracerProvider(provider)

	return provider
}

func opsrampProvider(ctx context.Context) *sdktrace.TracerProvider {
	// Create and configure the OTLP exporter to send traces to the collector
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create OTLP exporter: %v", err)
	}

	// Create a new trace provider with the exporter
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes("", semconv.ServiceNameKey.String("App2"))))
	otel.SetTracerProvider(provider)

	return provider
}

func main() {

	// Get the environment variable
	if len(os.Args) < 2 {
		log.Error("TraceProvider not provide")
		os.Exit(1)
	}

	TraceProvider = os.Args[1]
	log.Info("Received TraceProvider: ", TraceProvider)

	// check if the trace provider is newrelic or jaeger or opsramp
	var traceProvider *sdktrace.TracerProvider
	ctx := context.Background()

	// switch case to check the trace provider
	switch TraceProvider {
	case "newrelic":
		traceProvider = newRelicProvider(ctx)
	case "jaeger":
		traceProvider = jaegerProvider(ctx)
	case "opsramp":
		traceProvider = opsrampProvider(ctx)
	default:
		traceProvider = jaegerProvider(ctx)
	}
	defer shutdownTraceProvider(ctx, traceProvider)

	// Create a new Gin router
	r := gin.Default()

	// Define route handlers
	r.GET("/", HandlerLayer)

	// Start HTTP server
	log.Info("Server started on :5001")
	if err := http.ListenAndServe(":5001", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func DatabaseLayer(ctx context.Context) (string, error) {
	log.Info("Got a request in database layer")
	// Get the tracer from the global provider
	response := "Successfully fetched records from database"
	// Return the response
	return response, nil
}

func ServiceLayer(ctx context.Context) (string, error) {
	log.Info("Got a request in service layer")
	// Get the tracer from the global provider
	tracer := otel.Tracer("app2-service-layer")
	// Start a span
	ctxx, span := tracer.Start(ctx, "App2-ServiceLayer")
	defer span.End()
	span.AddEvent("Got a request in service layer")
	time.Sleep(100 * time.Millisecond)
	span.AddEvent("Calling database layer")
	response, err := DatabaseLayer(ctxx)
	if err != nil {
		span.RecordError(err)
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(500))
		span.SetStatus(codes.Error, err.Error())
		return response, err
	}
	return response, nil
}

func HandlerLayer(c *gin.Context) {
	log.Info("Got a verify request in app2")
	// Extract the trace ID and span ID from the incoming request
	traceID := c.Request.Header.Get("TraceID")
	spanID := c.Request.Header.Get("SpanID")
	// Convert the trace ID and span ID from strings to their appropriate types
	tID, _ := trace.TraceIDFromHex(traceID)
	sID, _ := trace.SpanIDFromHex(spanID)
	// Create a new span context with the extracted trace ID and span
	sc := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    tID,
		SpanID:     sID,
		TraceFlags: trace.FlagsSampled,
		Remote:     true,
	})
	ctx := trace.ContextWithRemoteSpanContext(context.Background(), sc)
	// Get the tracer from the global provider
	tracer := otel.GetTracerProvider().Tracer("app2-handler-layer")
	// Start a new span with the extracted context
	ctxx, span := tracer.Start(ctx, "App2-HandlerLayer")
	defer span.End()
	currentSpan := trace.SpanFromContext(ctx)
	currentTraceID := currentSpan.SpanContext().TraceID()
	currentSpanID := currentSpan.SpanContext().SpanID()
	// Print the extracted information
	log.Infof("Current Trace ID: %s\n", currentTraceID)
	log.Infof("Current Span ID: %s\n", currentSpanID)
	span.AddEvent("Got a request in app2 service")
	// call service layer
	span.AddEvent("Calling service layer")
	_, err := ServiceLayer(ctxx)
	if err != nil {
		log.Error("Error in service layer: ", err)
	}
	// Add an attribute to the span
	span.SetAttributes(semconv.HTTPMethodKey.String("GET"))
	span.SetAttributes(semconv.ServiceNameKey.String("App2"))
	// Return a response
	c.JSON(http.StatusOK, nil)
}
