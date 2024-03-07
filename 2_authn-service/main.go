package main

import (
	"context"
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
		sdktrace.WithResource(resource.NewWithAttributes("", semconv.ServiceNameKey.String("authn-service"))))
	otel.SetTracerProvider(provider)

	// Create a new Gin router
	r := gin.Default()

	// Define route handlers
	r.GET("/verify", AuthnHandler)

	// Start HTTP server
	log.Info("Server started on :5001")
	if err := http.ListenAndServe(":5001", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func AuthnHandler(c *gin.Context) {
	log.Info("Got a /verify request in authn-service")
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
	tracer := otel.GetTracerProvider().Tracer("authn-handler")
	// Start a new span with the extracted context
	_, span := tracer.Start(ctx, "AuthnHandler")
	defer span.End()
	currentSpan := trace.SpanFromContext(ctx)
	currentTraceID := currentSpan.SpanContext().TraceID()
	currentSpanID := currentSpan.SpanContext().SpanID()
	// Print the extracted information
	log.Infof("Current Trace ID: %s\n", currentTraceID)
	log.Infof("Current Span ID: %s\n", currentSpanID)
	span.AddEvent("Got a request to verify the request in authn-service")
	// Add an attribute to the span
	span.SetAttributes(semconv.HTTPMethodKey.String("GET"))
	span.SetAttributes(semconv.HTTPURLKey.String("/users"))
	span.SetAttributes(semconv.ServiceNameKey.String("authn-service"))
	// Add an event to the span
	span.AddEvent("Succefully verified the request in authn-service")
	// Return a response
	c.JSON(http.StatusOK, nil)
}