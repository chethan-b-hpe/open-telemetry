# Instrument golang transactions with Opentelementry, Newrelic and Jaeger

## Prerequisites

- Go 1.18+ (1.20 is used in this example)
- NewRelic Account
- Jaeger instance running in background

## Set up Jaeger instance

- Run this command to start [Jaeger](https://www.jaegertracing.io/docs/1.54/getting-started/) instance 
```go
docker run --rm --name jaeger \
  -e JAEGER_DISABLED=true \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.54
```

## Set up Environment Varables for Newrelic

```go
export OTEL_SERVICE_NAME=vipin-nr-service
vipin - export OTEL_EXPORTER_OTLP_HEADERS=api-key=5ff9356f375ebf68cd1d3879450a9ec1FFFFNRAL
chandan - export OTEL_EXPORTER_OTLP_HEADERS=api-key=46c7e0df19b313a7ed3dfc78c7dea347FFFFNRAL
export OTEL_EXPORTER_OTLP_ENDPOINT='https://otlp.nr-data.net'
export OTEL_ATTRIBUTE_VALUE_LENGTH_LIMIT=4095
```

## Run services

#### Postive case:

```go
go run 1_user-service/main.go http://localhost:5001/verify
go run 2_authn-service/main.go
```

#### Negative case:
```go
go run 1_user-service/main.go http://localhost:5001/
go run 2_authn-service/main.go
```
