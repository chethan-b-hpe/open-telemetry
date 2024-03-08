package main

import (
	"context"
	"errors"
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
		sdktrace.WithResource(resource.NewWithAttributes("", semconv.ServiceNameKey.String("App1"))))
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
		sdktrace.WithResource(resource.NewWithAttributes("", semconv.ServiceNameKey.String("App1"))))
	otel.SetTracerProvider(provider)

	return provider
}

func main() {

	// Get the environment variable
	if len(os.Args) < 2 {
		log.Error("URL argument is required")
		os.Exit(1)
	}

	URL = os.Args[1]
	log.Info("Received URL: ", URL)

	TraceProvider = os.Args[2]
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
	log.Info("Server started on :5000")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// Database is the handler for the / route
func Database(ctx context.Context) (map[string]string, error) {
	// Get the tracer from the global provider
	tracer := otel.Tracer("app1-database-layer")
	// Start a span
	ctx, span := tracer.Start(ctx, "App1-DatabaseLayer")
	defer ctx.Done()
	defer span.End()

	span.AddEvent("Fetching records from database")
	// Simulate a database call
	time.Sleep(100 * time.Millisecond)
	span.AddEvent("Successfully fetched records from database")
	return map[string]string{
		"user1": "hpe-user1",
		"user2": "hpe-user2",
	}, nil
}

// Service is the handler for the / route
func ServiceLayer(ctx context.Context) (map[string]string, error) {
	// Get the tracer from the global provider
	tracer := otel.Tracer("app1-service-layer")
	// Start a span
	ctx, span := tracer.Start(ctx, "App1-ServiceLayer")
	defer span.End()

	currentSpan := trace.SpanFromContext(ctx)
	currentTraceID := currentSpan.SpanContext().TraceID()
	currentSpanID := currentSpan.SpanContext().SpanID()
	// Print the extracted information
	log.Infof("Current Trace ID: %s\n", currentTraceID)
	log.Infof("Current Span ID: %s\n", currentSpanID)
	// Inject the trace context into the HTTP request headers
	span.AddEvent("Calling app2 service")
	// Call the app2 service
	req, _ := http.NewRequestWithContext(ctx, "GET", URL, nil)
	req = req.WithContext(ctx)
	req.Header.Set("TraceID", currentTraceID.String())
	req.Header.Set("SpanID", currentSpanID.String())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Failed to call app2 service")
		span.RecordError(errors.New("Failed to call app2 service"))
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(500))
		span.SetStatus(codes.Error, "Failed to call app2 service")
		return nil, err
	}
	defer resp.Body.Close()
	log.Info("app2 service response: ", resp.Status)
	if resp.StatusCode != http.StatusOK {
		log.Error("Invalid Request")
		span.RecordError(errors.New("Invalid Request"))
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(400))
		span.SetStatus(codes.Error, "Invalid Request")
		return nil, errors.New("Invalid Request")
	}
	response, err := Database(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(500))
		span.SetStatus(codes.Error, err.Error())
		return response, err
	}
	return response, nil

}

// HelloHandler is the handler for the /hello route
func HandlerLayer(c *gin.Context) {
	log.Info("Got a get request")
	// Get the tracer from the global provider
	tracer := otel.GetTracerProvider().Tracer("app1-handler-layer")
	// Start a span
	ctx, span := tracer.Start(context.Background(), "App1-HandlerLayer")
	defer span.End()
	span.AddEvent("Got a get request")
	// Add an attribute to the span
	span.SetAttributes(semconv.HTTPMethodKey.String("GET"))
	span.SetAttributes(semconv.ServiceNameKey.String("App1"))
	// Call the service layer
	resp, err := ServiceLayer(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(500))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error calling app2 service"})
		return
	}
	span.AddEvent("Successfully processed the request in app1")
	span.AddEvent("Successfully fetched records")
	span.SetStatus(codes.Ok, "Successfully fetched records")
	c.JSON(http.StatusOK, resp)
}
