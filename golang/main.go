package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
)

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
		sdktrace.WithResource(resource.NewWithAttributes("", semconv.ServiceNameKey.String("ServiceA"))))
	otel.SetTracerProvider(provider)

	return exporter
}

// HelloHandler is the handler for the /hello route
func HelloHandler(c *gin.Context) {
	// Get the tracer from the global provider
	// Start a span
	span := trace.SpanFromContext(c)
	ctx := trace.ContextWithSpan(c, span)
	defer span.End()
	span.AddEvent("handling the request")
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:5001/", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		span.RecordError(err)
		c.String(http.StatusInternalServerError, "Error calling Service A: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Service B response: %v", resp.Status)

	// Respond with "Hello, World!"
	c.String(http.StatusOK, "Hello, World!")
}
func main() {
	// Create and configure the OTLP exporter to send traces to the collector
	expoter := initExporter()
	defer expoter.Shutdown(context.Background())

	// Create a new Gin router
	r := gin.Default()

	// Define route handlers
	r.GET("/hello", HelloHandler)

	// Start HTTP server
	fmt.Println("Server started on :5000")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
