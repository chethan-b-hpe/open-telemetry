package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

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
		sdktrace.WithResource(resource.NewWithAttributes("", semconv.ServiceNameKey.String("ServiceA"))))
	otel.SetTracerProvider(provider)

	return provider
}

func main() {

	// var err error
	// password:= "vipin"

	ctx := context.Background()
	// get the jaeger provider
	// jagerProvider := jaegerProvider(ctx)
	// defer shutdownTraceProvider(ctx, jagerProvider)

	// get new relic provider
	newRelicProvider := newRelicProvider(ctx)
	defer shutdownTraceProvider(ctx, newRelicProvider)

	// Create a new Gin router
	r := gin.Default()

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
	userData := make(map[string]string)
	db, err := sql.Open("postgres", "postgres://vipin:vipin@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Get the tracer from the global provider
	tracer := otel.Tracer("user-database")
	// Start a span
	ctx, span := tracer.Start(context.Background(), "DB-Transactions")
	defer ctx.Done()
	defer span.End()

	// Start a new transaction to trace
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Create
	_, err = db.Exec("INSERT INTO users(name, email) VALUES($1, $2)", "Sheru1", "sheru@example.com")
	if err != nil {
		log.Printf("Error creating user: %v", err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// Start a new transaction to trace
	tx1, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Read
	rows, err := db.Query("SELECT name, email FROM users")
	if err != nil {
		log.Printf("Error reading users: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var name, email string
		err = rows.Scan(&name, &email)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
		}
		fmt.Println(name, email)
		userData[name] = email
	}

	err = tx1.Commit()
	if err != nil {
		log.Fatal(err)
	}

	span.AddEvent("user fetch done")

	return userData, nil
}

// UserService is the handler for the /users route
func UserService(ctx context.Context) (map[string]string, error) {
	// Get the tracer from the global provider
	tracer := otel.Tracer("user-service")
	// Start a span
	ctx, span := tracer.Start(ctx, "UserService")
	defer span.End()

	span.AddEvent("user verification done")
	userdetails, err := UserDatabase(ctx)
	if err != nil {
		span.RecordError(err)
		return userdetails, err
	}
	return userdetails, nil

}

// HelloHandler is the handler for the /hello route
func UserHandler(c *gin.Context) {
	log.Info("Got a request to get /users")
	// Get the tracer from the global provider
	tracer := otel.GetTracerProvider().Tracer("user-handler")
	// Start a span
	ctx, span := tracer.Start(c.Request.Context(), "UserHandler")
	defer span.End()
	span.AddEvent("handling get /users request")

	// Call the authz service
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:5001/verify", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		span.RecordError(err)
		log.Printf("Error calling authz service: %v", err)
		c.String(http.StatusInternalServerError, "Error calling authz service: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Info("Authz service response: ", resp.Status)

	// Add an attribute to the span
	span.SetAttributes(semconv.HTTPMethodKey.String("GET"))
	userdetails, err := UserService(ctx)
	if err != nil {
		span.RecordError(err)
		log.Printf("Error calling user service: %v", err)
		c.String(http.StatusInternalServerError, "Error calling user service: %v", err)
		return
	}
	span.AddEvent("getting user details")

	// Set some attributes on the span
	// span.SetAttributes(semconv.ServiceNameKey.String("sample-service"))

	// Respond with "Hello, World!"
	// return user json response

	c.JSON(http.StatusOK, userdetails)

}
