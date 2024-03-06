# Instrument golang transactions with opentelementry and newrelic

## Prerequisites

Go 1.18+ (1.20 is used in this example)
NewRelic Account

## Setting up Environment Varables

    export OTEL_SERVICE_NAME=vipin-nr-service
    export OTEL_EXPORTER_OTLP_HEADERS=api-key=5ff9356f375ebf68cd1d3879450a9ec1FFFFNRAL
    export OTEL_EXPORTER_OTLP_ENDPOINT='https://otlp.nr-data.net'
    export OTEL_ATTRIBUTE_VALUE_LENGTH_LIMIT=4095
