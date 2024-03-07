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
		sdktrace.WithResource(resource.NewWithAttributes("", semconv.ServiceNameKey.String("user-service"))))
	otel.SetTracerProvider(provider)

	return provider
}

func main() {

	ctx := context.Background()
	// get the jaeger provider
	jagerProvider := jaegerProvider(ctx)
	defer shutdownTraceProvider(ctx, jagerProvider)

	// get new relic provider
	// newRelicProvider := newRelicProvider(ctx)
	// defer shutdownTraceProvider(ctx, newRelicProvider)

	// Create a new Gin router
	r := gin.Default()

	// Get the environment variable
	if len(os.Args) < 2 {
		log.Error("URL argument is required")
		os.Exit(1)
	}

	URL = os.Args[1]
	log.Info("Received URL: ", URL)

	// Define route handlers
	r.GET("/users", UserHandler)

	// Start HTTP server
	log.Info("Server started on :5000")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// UserDatabase is the handler for the /users route
func UserDatabase(ctx context.Context) (map[string]string, error) {
	// Get the tracer from the global provider
	tracer := otel.Tracer("user-database")
	// Start a span
	ctx, span := tracer.Start(ctx, "UserDatabase")
	defer ctx.Done()
	defer span.End()

	span.AddEvent("Fetching user details from database")
	// Simulate a database call
	time.Sleep(100 * time.Millisecond)
	span.AddEvent("Successfully fetched user details from database")
	return map[string]string{
		"user1": "hpe-user1",
		"user2": "hpe-user2",
	}, nil
}

// UserService is the handler for the /users route
func UserService(ctx context.Context) (map[string]string, error) {
	// Get the tracer from the global provider
	tracer := otel.Tracer("user-service")
	// Start a span
	ctx, span := tracer.Start(ctx, "UserService")
	defer span.End()

	currentSpan := trace.SpanFromContext(ctx)
	currentTraceID := currentSpan.SpanContext().TraceID()
	currentSpanID := currentSpan.SpanContext().SpanID()
	// Print the extracted information
	log.Infof("Current Trace ID: %s\n", currentTraceID)
	log.Infof("Current Span ID: %s\n", currentSpanID)
	// Inject the trace context into the HTTP request headers
	span.AddEvent("Calling authn service")
	// Call the authz service
	req, _ := http.NewRequestWithContext(ctx, "GET", URL, nil)
	req = req.WithContext(ctx)
	req.Header.Set("TraceID", currentTraceID.String())
	req.Header.Set("SpanID", currentSpanID.String())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Failed to call authz service")
		span.RecordError(errors.New("Failed to call authz service"))
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(500))
		span.SetStatus(codes.Error, "Failed to call authz service")
		return nil, err
	}
	defer resp.Body.Close()
	log.Info("Authz service response: ", resp.Status)
	if resp.StatusCode != http.StatusOK {
		log.Error("Invalid Request")
		span.RecordError(errors.New("Invalid Request"))
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(400))
		span.SetStatus(codes.Error, "Invalid Request")
		return nil, errors.New("Invalid Request")
	}
	span.AddEvent("Successfully verified the request from authn service")
	userdetails, err := UserDatabase(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(500))
		span.SetStatus(codes.Error, err.Error())
		return userdetails, err
	}
	return userdetails, nil

}

// HelloHandler is the handler for the /hello route
func UserHandler(c *gin.Context) {
	log.Info("Got a request to get users")
	// Get the tracer from the global provider
	tracer := otel.GetTracerProvider().Tracer("user-handler")
	// Start a span
	ctx, span := tracer.Start(context.Background(), "UserHandler")
	defer span.End()
	span.AddEvent("Got a request to get users")
	// Add an attribute to the span
	span.SetAttributes(semconv.HTTPMethodKey.String("GET"))
	span.SetAttributes(semconv.HTTPURLKey.String("/users"))
	span.SetAttributes(semconv.ServiceNameKey.String("user-service"))
	// Call the user service
	userdetails, err := UserService(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(500))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error calling authn service"})
		return
	}
	span.AddEvent("Successfully verfied the request from authn service")
	span.AddEvent("user details fetched")
	span.SetStatus(codes.Ok, "User details fetched successfully")
	c.JSON(http.StatusOK, userdetails)
}
