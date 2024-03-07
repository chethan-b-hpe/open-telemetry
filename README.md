# Instrument golang transactions with opentelementry and newrelic

## Prerequisites

Go 1.18+ (1.20 is used in this example)
NewRelic Account

## Setting up Environment Varables for Instrumentation with NewRelic

    export OTEL_SERVICE_NAME=vipin-nr-service
    export OTEL_EXPORTER_OTLP_HEADERS=api-key=5ff9356f375ebf68cd1d3879450a9ec1FFFFNRAL
    export OTEL_EXPORTER_OTLP_ENDPOINT='https://otlp.nr-data.net'
    export OTEL_ATTRIBUTE_VALUE_LENGTH_LIMIT=4095

## Setting up Jaeger for local Tests

        docker run --rm --name jaeger \
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
