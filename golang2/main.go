package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
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
		sdktrace.WithResource(resource.NewWithAttributes("", semconv.ServiceNameKey.String("ServiceB"))))
	otel.SetTracerProvider(provider)

	return exporter
}

// HelloHandler is the handler for the /hello route
func Handler(c *gin.Context) {
	// Get the tracer from the global provider
	tracer := otel.GetTracerProvider().Tracer("serviceB")

	// Start a span
	_, span := tracer.Start(c.Request.Context(), "HelloHandler")

	// Simulate some work
	time.Sleep(time.Second)

	span.AddEvent("handling the request in Service B")
	// Respond with "Hello, World!"
	c.String(http.StatusOK, "Hello from Service B!")
}
func main() {
	// Create and configure the OTLP exporter to send traces to the collector
	expoter := initExporter()
	defer expoter.Shutdown(context.Background())

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
