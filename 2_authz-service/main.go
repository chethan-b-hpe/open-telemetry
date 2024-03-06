package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	// Create and configure the OTLP exporter to send traces to the collector
	exporter, err := otlptracegrpc.New(context.Background(), otlptracegrpc.WithEndpointURL("http://localhost:4317/api/traces"))
	if err != nil {
		log.Fatalf("failed to create OTLP exporter: %v", err)
	}
	defer exporter.Shutdown(context.Background())

	// Create a new trace provider with the exporter
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes("", semconv.ServiceNameKey.String("authz-service"))))
	otel.SetTracerProvider(provider)

	// Create a new Gin router
	r := gin.Default()

	// Define route handlers
	r.GET("/verify", AuthzHandler)

	// Start HTTP server
	log.Info("Server started on :5001")
	if err := http.ListenAndServe(":5001", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func AuthzHandler(c *gin.Context) {

	// Extract the context from the incoming request
	// propagator := otel.GetTextMapPropagator()
	// parentCtx := propagator.Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
	traceID := c.Request.Header.Get("TraceID")
	spanID := c.Request.Header.Get("SpanID")
	fmt.Printf("Parent-TraceID: %s\n", traceID)
	fmt.Printf("SpanID 1: %s\n", spanID)
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
	tracer := otel.GetTracerProvider().Tracer("authz-handler")

	// Start a new span with the extracted context
	_, span := tracer.Start(ctx, "AuthzHandler")
	defer span.End()

	currentSpan := trace.SpanFromContext(ctx)
	currentTraceID := currentSpan.SpanContext().TraceID()
	currentSpanID := currentSpan.SpanContext().SpanID()
	// Print the extracted information
	fmt.Printf("Current Trace ID: %s\n", currentTraceID)
	fmt.Printf("Current Span ID: %s\n", currentSpanID)
	span.AddEvent("handling the /verify request in authz-service")

	// Add an attribute to the span
	span.SetAttributes(semconv.HTTPMethodKey.String("GET"))

	span.AddEvent("verified the request in authz-service")

	// Set some attributes on the span
	span.SetAttributes(semconv.ServiceNameKey.String("authz-service"))

	// Respond with "user verified successfully!"
	c.String(http.StatusOK, "user verified successfully!")
}
