2024-03-06T07:19:14.400Z	info	service@v0.96.0/telemetry.go:55	Setting up own telemetry...
2024-03-06T07:19:14.400Z	info	service@v0.96.0/telemetry.go:97	Serving metrics	{"address": ":8888", "level": "Basic"}
2024-03-06T07:19:14.401Z	info	exporter@v0.96.0/exporter.go:275	Development component. May change in the future.	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:19:14.401Z	info	exporter@v0.96.0/exporter.go:275	Development component. May change in the future.	{"kind": "exporter", "data_type": "traces", "name": "debug"}
2024-03-06T07:19:14.404Z	info	exporter@v0.96.0/exporter.go:275	Development component. May change in the future.	{"kind": "exporter", "data_type": "logs", "name": "debug"}
2024-03-06T07:19:14.418Z	warn	jaegerreceiver@v0.96.0/factory.go:49	jaeger receiver will deprecate Thrift-gen and replace it with Proto-gen to be compatbible to jaeger 1.42.0 and higher. See https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/18485 for more details.	{"kind": "receiver", "name": "jaeger", "data_type": "traces"}
2024-03-06T07:19:14.419Z	info	service@v0.96.0/service.go:143	Starting otelcol...	{"Version": "0.96.0", "NumCPU": 6}
2024-03-06T07:19:14.419Z	info	extensions/extensions.go:34	Starting extensions...
2024-03-06T07:19:14.419Z	info	extensions/extensions.go:37	Extension is starting...	{"kind": "extension", "name": "zpages"}
2024-03-06T07:19:14.419Z	info	zpagesextension@v0.96.0/zpagesextension.go:53	Registered zPages span processor on tracer provider	{"kind": "extension", "name": "zpages"}
2024-03-06T07:19:14.419Z	info	zpagesextension@v0.96.0/zpagesextension.go:63	Registered Host's zPages	{"kind": "extension", "name": "zpages"}
2024-03-06T07:19:14.420Z	info	zpagesextension@v0.96.0/zpagesextension.go:75	Starting zPages extension	{"kind": "extension", "name": "zpages", "config": {"TCPAddr":{"Endpoint":"0.0.0.0:55679","DialerConfig":{"Timeout":0}}}}
2024-03-06T07:19:14.420Z	info	extensions/extensions.go:52	Extension started.	{"kind": "extension", "name": "zpages"}
2024-03-06T07:19:14.420Z	info	extensions/extensions.go:37	Extension is starting...	{"kind": "extension", "name": "pprof"}
2024-03-06T07:19:14.420Z	info	pprofextension@v0.96.0/pprofextension.go:60	Starting net/http/pprof server	{"kind": "extension", "name": "pprof", "config": {"TCPAddr":{"Endpoint":"0.0.0.0:1777","DialerConfig":{"Timeout":0}},"BlockProfileFraction":0,"MutexProfileFraction":0,"SaveToFile":""}}
2024-03-06T07:19:14.420Z	info	extensions/extensions.go:52	Extension started.	{"kind": "extension", "name": "pprof"}
2024-03-06T07:19:14.420Z	info	extensions/extensions.go:37	Extension is starting...	{"kind": "extension", "name": "health_check"}
2024-03-06T07:19:14.420Z	info	healthcheckextension@v0.96.0/healthcheckextension.go:35	Starting health_check extension	{"kind": "extension", "name": "health_check", "config": {"Endpoint":"0.0.0.0:13133","TLSSetting":null,"CORS":null,"Auth":null,"MaxRequestBodySize":0,"IncludeMetadata":false,"ResponseHeaders":null,"Path":"/","ResponseBody":null,"CheckCollectorPipeline":{"Enabled":false,"Interval":"5m","ExporterFailureThreshold":5}}}
2024-03-06T07:19:14.421Z	warn	internal@v0.96.0/warning.go:42	Using the 0.0.0.0 address exposes this server to every network interface, which may facilitate Denial of Service attacks. Enable the feature gate to change the default and remove this warning.	{"kind": "extension", "name": "health_check", "documentation": "https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/security-best-practices.md#safeguards-against-denial-of-service-attacks", "feature gate ID": "component.UseLocalHostAsDefaultHost"}
2024-03-06T07:19:14.421Z	info	extensions/extensions.go:52	Extension started.	{"kind": "extension", "name": "health_check"}
2024-03-06T07:19:14.421Z	warn	internal@v0.96.0/warning.go:42	Using the 0.0.0.0 address exposes this server to every network interface, which may facilitate Denial of Service attacks. Enable the feature gate to change the default and remove this warning.	{"kind": "receiver", "name": "opencensus", "data_type": "metrics", "documentation": "https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/security-best-practices.md#safeguards-against-denial-of-service-attacks", "feature gate ID": "component.UseLocalHostAsDefaultHost"}
2024-03-06T07:19:14.424Z	info	prometheusreceiver@v0.96.0/metrics_receiver.go:231	Scrape job added	{"kind": "receiver", "name": "prometheus", "data_type": "metrics", "jobName": "otel-collector"}
2024-03-06T07:19:14.424Z	warn	internal@v0.96.0/warning.go:42	Using the 0.0.0.0 address exposes this server to every network interface, which may facilitate Denial of Service attacks. Enable the feature gate to change the default and remove this warning.	{"kind": "receiver", "name": "zipkin", "data_type": "traces", "documentation": "https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/security-best-practices.md#safeguards-against-denial-of-service-attacks", "feature gate ID": "component.UseLocalHostAsDefaultHost"}
2024-03-06T07:19:14.426Z	warn	internal@v0.96.0/warning.go:42	Using the 0.0.0.0 address exposes this server to every network interface, which may facilitate Denial of Service attacks. Enable the feature gate to change the default and remove this warning.	{"kind": "receiver", "name": "otlp", "data_type": "metrics", "documentation": "https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/security-best-practices.md#safeguards-against-denial-of-service-attacks", "feature gate ID": "component.UseLocalHostAsDefaultHost"}
2024-03-06T07:19:14.426Z	info	otlpreceiver@v0.96.0/otlp.go:102	Starting GRPC server	{"kind": "receiver", "name": "otlp", "data_type": "metrics", "endpoint": "0.0.0.0:4317"}
2024-03-06T07:19:14.426Z	warn	internal@v0.96.0/warning.go:42	Using the 0.0.0.0 address exposes this server to every network interface, which may facilitate Denial of Service attacks. Enable the feature gate to change the default and remove this warning.	{"kind": "receiver", "name": "otlp", "data_type": "metrics", "documentation": "https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/security-best-practices.md#safeguards-against-denial-of-service-attacks", "feature gate ID": "component.UseLocalHostAsDefaultHost"}
2024-03-06T07:19:14.426Z	info	otlpreceiver@v0.96.0/otlp.go:152	Starting HTTP server	{"kind": "receiver", "name": "otlp", "data_type": "metrics", "endpoint": "0.0.0.0:4318"}
2024-03-06T07:19:14.426Z	warn	internal@v0.96.0/warning.go:42	Using the 0.0.0.0 address exposes this server to every network interface, which may facilitate Denial of Service attacks. Enable the feature gate to change the default and remove this warning.	{"kind": "receiver", "name": "jaeger", "data_type": "traces", "documentation": "https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/security-best-practices.md#safeguards-against-denial-of-service-attacks", "feature gate ID": "component.UseLocalHostAsDefaultHost"}
2024-03-06T07:19:14.426Z	warn	internal@v0.96.0/warning.go:42	Using the 0.0.0.0 address exposes this server to every network interface, which may facilitate Denial of Service attacks. Enable the feature gate to change the default and remove this warning.	{"kind": "receiver", "name": "jaeger", "data_type": "traces", "documentation": "https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/security-best-practices.md#safeguards-against-denial-of-service-attacks", "feature gate ID": "component.UseLocalHostAsDefaultHost"}
2024-03-06T07:19:14.426Z	info	healthcheck/handler.go:132	Health Check state change	{"kind": "extension", "name": "health_check", "status": "ready"}
2024-03-06T07:19:14.427Z	info	service@v0.96.0/service.go:169	Everything is ready. Begin running and processing data.
2024-03-06T07:19:14.427Z	warn	localhostgate/featuregate.go:63	The default endpoints for all servers in components will change to use localhost instead of 0.0.0.0 in a future version. Use the feature gate to preview the new default.	{"feature gate ID": "component.UseLocalHostAsDefaultHost"}
2024-03-06T07:19:14.427Z	info	prometheusreceiver@v0.96.0/metrics_receiver.go:240	Starting discovery manager	{"kind": "receiver", "name": "prometheus", "data_type": "metrics"}
2024-03-06T07:19:14.428Z	info	prometheusreceiver@v0.96.0/metrics_receiver.go:282	Starting scrape manager	{"kind": "receiver", "name": "prometheus", "data_type": "metrics"}
2024-03-06T07:19:27.031Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 12, "data points": 12}
2024-03-06T07:19:27.032Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:26.909 +0000 UTC
Value: 9104480.000000
Metric #1
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:19:26.909 +0000 UTC
Value: 15306768.000000
Metric #2
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:19:26.909 +0000 UTC
Value: 12.491654
Metric #3
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:26.909 +0000 UTC
Value: 1.000000
Metric #4
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:26.909 +0000 UTC
Value: 0.003056
Metric #5
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:26.909 +0000 UTC
Value: 68333568.000000
Metric #6
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:26.909 +0000 UTC
Value: 24483080.000000
Metric #7
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:26.909 +0000 UTC
Value: 1.000000
Metric #8
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:26.909 +0000 UTC
Value: 8.000000
Metric #9
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:26.909 +0000 UTC
Value: 8.000000
Metric #10
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:26.909 +0000 UTC
Value: 8.000000
Metric #11
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:19:26.909 +0000 UTC
Value: 0.300000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:19:37.060Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 18, "data points": 18}
2024-03-06T07:19:37.060Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 69734400.000000
Metric #1
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Count: 1
Sum: 12.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #2
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 1.000000
Metric #3
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 38.000000
Metric #4
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 38.000000
Metric #5
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 16443344.000000
Metric #6
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 1.000000
Metric #7
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 12.000000
Metric #8
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 0.003420
Metric #9
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 0.000000
Metric #10
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 1.000000
Metric #11
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 12.000000
Metric #12
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 0.360000
Metric #13
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 10241056.000000
Metric #14
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 24483080.000000
Metric #15
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 22.492162
Metric #16
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 0.000000
Metric #17
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:36.909 +0000 UTC
Value: 38.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:19:47.092Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 18, "data points": 18}
2024-03-06T07:19:47.093Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 38.000000
Metric #1
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 30.000000
Metric #2
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 2.000000
Metric #3
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 1.000000
Metric #4
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 0.002819
Metric #5
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 38.000000
Metric #6
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 17533808.000000
Metric #7
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 1.000000
Metric #8
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 38.000000
Metric #9
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 0.000000
Metric #10
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 0.440000
Metric #11
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 70754304.000000
Metric #12
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Count: 2
Sum: 30.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 0
Buckets #1, Count: 2
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #13
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 30.000000
Metric #14
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 0.000000
Metric #15
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 11331520.000000
Metric #16
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 24483080.000000
Metric #17
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:19:46.909 +0000 UTC
Value: 32.492706
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:19:56.925Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 18, "data points": 18}
2024-03-06T07:19:56.925Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 1.000000
Metric #1
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 0.500000
Metric #2
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 12397600.000000
Metric #3
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 18599888.000000
Metric #4
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Count: 3
Sum: 48.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 0
Buckets #1, Count: 3
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #5
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 0.000000
Metric #6
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 3.000000
Metric #7
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 0.000000
Metric #8
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 48.000000
Metric #9
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 71311360.000000
Metric #10
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 1.000000
Metric #11
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 48.000000
Metric #12
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 38.000000
Metric #13
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 38.000000
Metric #14
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 28677384.000000
Metric #15
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 42.491938
Metric #16
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 0.001997
Metric #17
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:19:56.909 +0000 UTC
Value: 38.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:20:06.956Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 18, "data points": 18}
2024-03-06T07:20:06.956Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 66.000000
Metric #1
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 66.000000
Metric #2
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 0.000000
Metric #3
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 71548928.000000
Metric #4
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 28677384.000000
Metric #5
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 52.493013
Metric #6
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 1.000000
Metric #7
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 38.000000
Metric #8
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 0.570000
Metric #9
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 12620464.000000
Metric #10
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Count: 4
Sum: 66.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 0
Buckets #1, Count: 4
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #11
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 4.000000
Metric #12
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 1.000000
Metric #13
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 0.003074
Metric #14
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 38.000000
Metric #15
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 0.000000
Metric #16
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 18822752.000000
Metric #17
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:06.909 +0000 UTC
Value: 38.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:20:16.988Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 18, "data points": 18}
2024-03-06T07:20:16.988Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 5.000000
Metric #1
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 0.001844
Metric #2
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Count: 5
Sum: 84.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 0
Buckets #1, Count: 5
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #3
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 84.000000
Metric #4
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 0.000000
Metric #5
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 38.000000
Metric #6
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 0.000000
Metric #7
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 84.000000
Metric #8
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 28677384.000000
Metric #9
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 62.492917
Metric #10
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 0.630000
Metric #11
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 71839744.000000
Metric #12
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 19045376.000000
Metric #13
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 38.000000
Metric #14
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 12843088.000000
Metric #15
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 1.000000
Metric #16
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 1.000000
Metric #17
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:16.909 +0000 UTC
Value: 38.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:20:27.021Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 18, "data points": 18}
2024-03-06T07:20:27.021Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 13084136.000000
Metric #1
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 28677384.000000
Metric #2
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Count: 6
Sum: 102.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 0
Buckets #1, Count: 6
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #3
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 38.000000
Metric #4
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 6.000000
Metric #5
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 0.001861
Metric #6
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 38.000000
Metric #7
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 0.000000
Metric #8
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 102.000000
Metric #9
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 19286424.000000
Metric #10
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 1.000000
Metric #11
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 102.000000
Metric #12
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 38.000000
Metric #13
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 0.700000
Metric #14
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 71839744.000000
Metric #15
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 72.493069
Metric #16
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 0.000000
Metric #17
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:26.909 +0000 UTC
Value: 1.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:20:37.052Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 18, "data points": 18}
2024-03-06T07:20:37.053Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 0.000000
Metric #1
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 38.000000
Metric #2
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 0.000000
Metric #3
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 120.000000
Metric #4
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 19509176.000000
Metric #5
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Count: 7
Sum: 120.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 0
Buckets #1, Count: 7
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #6
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 7.000000
Metric #7
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 120.000000
Metric #8
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 0.001811
Metric #9
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 38.000000
Metric #10
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 0.770000
Metric #11
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 73641984.000000
Metric #12
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 28677384.000000
Metric #13
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 82.492315
Metric #14
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 13306888.000000
Metric #15
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 1.000000
Metric #16
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 1.000000
Metric #17
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:36.909 +0000 UTC
Value: 38.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:20:47.084Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 18, "data points": 18}
2024-03-06T07:20:47.084Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 38.000000
Metric #1
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 19731752.000000
Metric #2
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 138.000000
Metric #3
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 1.000000
Metric #4
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 1.000000
Metric #5
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 0.000000
Metric #6
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 0.002127
Metric #7
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 38.000000
Metric #8
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 0.000000
Metric #9
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 73641984.000000
Metric #10
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Count: 8
Sum: 138.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 0
Buckets #1, Count: 8
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #11
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 8.000000
Metric #12
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 0.850000
Metric #13
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 13529464.000000
Metric #14
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 28677384.000000
Metric #15
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 138.000000
Metric #16
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 92.493948
Metric #17
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:46.912 +0000 UTC
Value: 38.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:20:53.128Z	info	TracesExporter	{"kind": "exporter", "data_type": "traces", "name": "debug", "resource spans": 1, "spans": 6}
2024-03-06T07:20:53.128Z	info	ResourceSpans #0
Resource SchemaURL: https://opentelemetry.io/schemas/1.4.0
Resource attributes:
     -> service.name: Str(telemetrygen)
ScopeSpans #0
ScopeSpans SchemaURL: 
InstrumentationScope telemetrygen 
Span #0
    Trace ID       : 7568d34678fc64970881d936cc3f03f6
    Parent ID      : 5a769294a07806db
    ID             : ce6a7feef625fb5f
    Name           : okey-dokey-0
    Kind           : Server
    Start time     : 2024-03-06 07:20:53.044178 +0000 UTC
    End time       : 2024-03-06 07:20:53.044301 +0000 UTC
    Status code    : Unset
    Status message : 
Attributes:
     -> net.peer.ip: Str(1.2.3.4)
     -> peer.service: Str(telemetrygen-client)
Span #1
    Trace ID       : 7568d34678fc64970881d936cc3f03f6
    Parent ID      : 
    ID             : 5a769294a07806db
    Name           : lets-go
    Kind           : Client
    Start time     : 2024-03-06 07:20:53.044178 +0000 UTC
    End time       : 2024-03-06 07:20:53.044301 +0000 UTC
    Status code    : Unset
    Status message : 
Attributes:
     -> net.peer.ip: Str(1.2.3.4)
     -> peer.service: Str(telemetrygen-server)
Span #2
    Trace ID       : 00ea3f0adfc70b260ab439a9a357e4c0
    Parent ID      : b6af8b580405845b
    ID             : 5b3c1ba8b1531c96
    Name           : okey-dokey-0
    Kind           : Server
    Start time     : 2024-03-06 07:20:53.044298 +0000 UTC
    End time       : 2024-03-06 07:20:53.044421 +0000 UTC
    Status code    : Unset
    Status message : 
Attributes:
     -> net.peer.ip: Str(1.2.3.4)
     -> peer.service: Str(telemetrygen-client)
Span #3
    Trace ID       : 00ea3f0adfc70b260ab439a9a357e4c0
    Parent ID      : 
    ID             : b6af8b580405845b
    Name           : lets-go
    Kind           : Client
    Start time     : 2024-03-06 07:20:53.044298 +0000 UTC
    End time       : 2024-03-06 07:20:53.044421 +0000 UTC
    Status code    : Unset
    Status message : 
Attributes:
     -> net.peer.ip: Str(1.2.3.4)
     -> peer.service: Str(telemetrygen-server)
Span #4
    Trace ID       : 3d0ff444a206da51eb0f1065696e565d
    Parent ID      : 618ed68584632eca
    ID             : b5bc843780fa1940
    Name           : okey-dokey-0
    Kind           : Server
    Start time     : 2024-03-06 07:20:53.044305 +0000 UTC
    End time       : 2024-03-06 07:20:53.044428 +0000 UTC
    Status code    : Unset
    Status message : 
Attributes:
     -> net.peer.ip: Str(1.2.3.4)
     -> peer.service: Str(telemetrygen-client)
Span #5
    Trace ID       : 3d0ff444a206da51eb0f1065696e565d
    Parent ID      : 
    ID             : 618ed68584632eca
    Name           : lets-go
    Kind           : Client
    Start time     : 2024-03-06 07:20:53.044305 +0000 UTC
    End time       : 2024-03-06 07:20:53.044428 +0000 UTC
    Status code    : Unset
    Status message : 
Attributes:
     -> net.peer.ip: Str(1.2.3.4)
     -> peer.service: Str(telemetrygen-server)
	{"kind": "exporter", "data_type": "traces", "name": "debug"}
2024-03-06T07:20:57.114Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:20:57.114Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 10.000000
Metric #1
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 6.000000
Metric #2
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 0.940000
Metric #4
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 0.000000
Metric #5
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 1.000000
Metric #6
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 156.000000
Metric #7
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 132.000000
Metric #8
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 0.000000
Metric #9
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 28677384.000000
Metric #10
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #11
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 132.000000
Metric #12
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 0.000000
Metric #13
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 74637312.000000
Metric #14
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 20719592.000000
Metric #15
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #16
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #17
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #18
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 132.000000
Metric #19
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 6.000000
Metric #20
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Count: 10
Sum: 162.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #21
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 102.492866
Metric #22
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 1.000000
Metric #23
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 156.000000
Metric #24
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 0.000000
Metric #25
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 0.004074
Metric #26
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:20:56.909 +0000 UTC
Value: 12215472.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:21:06.954Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:21:06.955Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 28677384.000000
Metric #1
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Count: 11
Sum: 189.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 1
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #2
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 11.000000
Metric #3
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 0.000000
Metric #4
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #5
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #6
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 132.000000
Metric #7
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 13618968.000000
Metric #8
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 1.000000
Metric #9
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 6.000000
Metric #10
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 0.011166
Metric #11
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 132.000000
Metric #12
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 183.000000
Metric #13
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 0.000000
Metric #14
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #15
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 1.000000
Metric #16
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 0.000000
Metric #17
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 183.000000
Metric #18
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #19
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 132.000000
Metric #20
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 75460608.000000
Metric #21
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 1.020000
Metric #22
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 112.494129
Metric #23
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #24
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 6.000000
Metric #25
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 22123088.000000
Metric #26
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:06.912 +0000 UTC
Value: 0.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:21:16.990Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:21:16.990Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 76800000.000000
Metric #1
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 0.000000
Metric #2
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #4
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 132.000000
Metric #5
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 0.000000
Metric #6
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 23519384.000000
Metric #7
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 1.000000
Metric #8
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 0.003572
Metric #9
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 15015264.000000
Metric #10
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Count: 12
Sum: 216.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 2
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #11
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 6.000000
Metric #12
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 0.000000
Metric #13
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #14
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 12.000000
Metric #15
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #16
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 6.000000
Metric #17
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 210.000000
Metric #18
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #19
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 1.000000
Metric #20
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 1.100000
Metric #21
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 28677384.000000
Metric #22
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 132.000000
Metric #23
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 132.000000
Metric #24
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 0.000000
Metric #25
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 210.000000
Metric #26
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:16.912 +0000 UTC
Value: 122.493709
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:21:27.028Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:21:27.029Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 237.000000
Metric #1
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 1.180000
Metric #2
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 132.000000
Metric #4
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 132.493924
Metric #5
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 13.000000
Metric #6
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 0.000000
Metric #7
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 132.000000
Metric #8
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 6.000000
Metric #9
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #10
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 0.000000
Metric #11
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 28677384.000000
Metric #12
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Count: 13
Sum: 243.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 3
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #13
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #14
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 1.000000
Metric #15
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 237.000000
Metric #16
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #17
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 0.004704
Metric #18
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 6.000000
Metric #19
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 77606912.000000
Metric #20
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 16290336.000000
Metric #21
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #22
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 0.000000
Metric #23
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 0.000000
Metric #24
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 1.000000
Metric #25
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 132.000000
Metric #26
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:26.912 +0000 UTC
Value: 24794456.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:21:37.062Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:21:37.062Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 32871688.000000
Metric #1
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #2
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 264.000000
Metric #3
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 1.250000
Metric #4
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #5
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 1.000000
Metric #6
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 132.000000
Metric #7
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 142.494271
Metric #8
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Count: 14
Sum: 270.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 4
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #9
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 1.000000
Metric #10
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 264.000000
Metric #11
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 78446592.000000
Metric #12
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 14.000000
Metric #13
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 0.000000
Metric #14
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #15
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #16
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #17
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 0.000000
Metric #18
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 132.000000
Metric #19
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 132.000000
Metric #20
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 26154088.000000
Metric #21
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 0.000000
Metric #22
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 0.003114
Metric #23
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 0.000000
Metric #24
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 6.000000
Metric #25
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 17649968.000000
Metric #26
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:36.912 +0000 UTC
Value: 6.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:21:47.090Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:21:47.090Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 32871688.000000
Metric #1
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 0.000000
Metric #2
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #3
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #4
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 18116808.000000
Metric #5
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 152.494666
Metric #6
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 1.000000
Metric #7
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #8
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 0.000000
Metric #9
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #10
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 6.000000
Metric #11
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 78663680.000000
Metric #12
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 132.000000
Metric #13
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 132.000000
Metric #14
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 291.000000
Metric #15
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 15.000000
Metric #16
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 6.000000
Metric #17
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 291.000000
Metric #18
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 1.000000
Metric #19
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 0.000000
Metric #20
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #21
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 132.000000
Metric #22
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 0.000000
Metric #23
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 1.330000
Metric #24
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 26620928.000000
Metric #25
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Count: 15
Sum: 297.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 5
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #26
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:46.912 +0000 UTC
Value: 0.004722
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:21:56.923Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:21:56.923Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 0.000000
Metric #1
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 1.410000
Metric #2
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 162.493846
Metric #3
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 132.000000
Metric #4
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 0.000000
Metric #5
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 132.000000
Metric #6
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 1.000000
Metric #7
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 132.000000
Metric #8
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 78901248.000000
Metric #9
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #10
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #11
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 27078096.000000
Metric #12
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 18573976.000000
Metric #13
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 16.000000
Metric #14
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 6.000000
Metric #15
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 318.000000
Metric #16
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 318.000000
Metric #17
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 0.000000
Metric #18
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 1.000000
Metric #19
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 32871688.000000
Metric #20
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #21
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 0.003338
Metric #22
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Count: 16
Sum: 324.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 6
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #23
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 0.000000
Metric #24
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #25
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #26
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:21:56.912 +0000 UTC
Value: 6.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:22:06.950Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:22:06.951Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #1
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 132.000000
Metric #2
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 0.000000
Metric #3
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 345.000000
Metric #4
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 0.000000
Metric #5
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #6
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 0.000000
Metric #7
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 1.000000
Metric #8
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 19031112.000000
Metric #9
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 32871688.000000
Metric #10
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 0.000000
Metric #11
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 132.000000
Metric #12
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 1.000000
Metric #13
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 345.000000
Metric #14
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 172.494027
Metric #15
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Count: 17
Sum: 351.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 7
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #16
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 6.000000
Metric #17
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #18
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #19
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #20
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 6.000000
Metric #21
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 27535232.000000
Metric #22
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 0.003641
Metric #23
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 132.000000
Metric #24
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 1.470000
Metric #25
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 79400960.000000
Metric #26
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:06.912 +0000 UTC
Value: 17.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:22:16.983Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:22:16.983Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 372.000000
Metric #1
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 0.000000
Metric #2
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 372.000000
Metric #3
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 1.000000
Metric #4
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #5
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 6.000000
Metric #6
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 27993688.000000
Metric #7
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 19489568.000000
Metric #8
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 0.000000
Metric #9
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 1.540000
Metric #10
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 79654912.000000
Metric #11
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #12
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Count: 18
Sum: 378.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 8
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #13
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 18.000000
Metric #14
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 6.000000
Metric #15
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #16
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 0.003741
Metric #17
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 132.000000
Metric #18
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 0.000000
Metric #19
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 182.495405
Metric #20
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 33133832.000000
Metric #21
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #22
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 0.000000
Metric #23
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 132.000000
Metric #24
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 132.000000
Metric #25
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #26
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:16.913 +0000 UTC
Value: 1.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:22:27.013Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:22:27.014Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 6.000000
Metric #1
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 399.000000
Metric #2
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 0.000000
Metric #3
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 1.000000
Metric #4
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #5
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #6
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 1.000000
Metric #7
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 132.000000
Metric #8
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 0.000000
Metric #9
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 0.000000
Metric #10
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 399.000000
Metric #11
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Count: 19
Sum: 405.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 9
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #12
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #13
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #14
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 132.000000
Metric #15
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 79917056.000000
Metric #16
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 28452040.000000
Metric #17
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 33133832.000000
Metric #18
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 192.494853
Metric #19
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 19.000000
Metric #20
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 6.000000
Metric #21
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 0.003083
Metric #22
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 19947920.000000
Metric #23
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 0.000000
Metric #24
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #25
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 1.620000
Metric #26
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:26.913 +0000 UTC
Value: 132.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:22:37.045Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:22:37.046Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 0.003355
Metric #1
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 132.000000
Metric #2
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 132.000000
Metric #3
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 426.000000
Metric #4
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 6.000000
Metric #5
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 0.000000
Metric #6
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 20626760.000000
Metric #7
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #8
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #9
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 80662528.000000
Metric #10
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 426.000000
Metric #11
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #12
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 1.000000
Metric #13
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 0.000000
Metric #14
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 6.000000
Metric #15
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 29130880.000000
Metric #16
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Count: 20
Sum: 432.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 10
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #17
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 1.000000
Metric #18
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 0.000000
Metric #19
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 20.000000
Metric #20
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #21
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 33133832.000000
Metric #22
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #23
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 132.000000
Metric #24
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 0.000000
Metric #25
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 1.690000
Metric #26
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:36.913 +0000 UTC
Value: 202.494886
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:22:47.080Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:22:47.081Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 81162240.000000
Metric #1
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 29588376.000000
Metric #2
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 21.000000
Metric #3
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #4
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #5
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #6
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 132.000000
Metric #7
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 1.000000
Metric #8
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 6.000000
Metric #9
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Count: 21
Sum: 459.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 11
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #10
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 453.000000
Metric #11
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 0.000000
Metric #12
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 453.000000
Metric #13
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 33133832.000000
Metric #14
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #15
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 0.005965
Metric #16
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 132.000000
Metric #17
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 212.495083
Metric #18
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #19
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 132.000000
Metric #20
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 0.000000
Metric #21
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 1.000000
Metric #22
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 0.000000
Metric #23
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 6.000000
Metric #24
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 1.770000
Metric #25
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 21084256.000000
Metric #26
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:46.913 +0000 UTC
Value: 0.000000
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
2024-03-06T07:22:57.119Z	info	MetricsExporter	{"kind": "exporter", "data_type": "metrics", "name": "debug", "resource metrics": 1, "metrics": 27, "data points": 27}
2024-03-06T07:22:57.119Z	info	ResourceMetrics #0
Resource SchemaURL: 
Resource attributes:
     -> service.name: Str(otel-collector)
     -> service.instance.id: Str(0.0.0.0:8888)
     -> net.host.port: Str(8888)
     -> http.scheme: Str(http)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
ScopeMetrics #0
ScopeMetrics SchemaURL: 
InstrumentationScope otelcol/prometheusreceiver 0.96.0
Metric #0
Descriptor:
     -> Name: otelcol_exporter_send_failed_metric_points
     -> Description: Number of metric points in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 0.000000
Metric #1
Descriptor:
     -> Name: scrape_samples_scraped
     -> Description: The number of samples the target exposed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 132.000000
Metric #2
Descriptor:
     -> Name: scrape_series_added
     -> Description: The approximate number of new series in this scrape
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 132.000000
Metric #3
Descriptor:
     -> Name: otelcol_receiver_accepted_metric_points
     -> Description: Number of metric points successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 480.000000
Metric #4
Descriptor:
     -> Name: otelcol_rpc_server_responses_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #5
Descriptor:
     -> Name: up
     -> Description: The scraping was successful
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 1.000000
Metric #6
Descriptor:
     -> Name: scrape_duration_seconds
     -> Description: Duration of the scrape
     -> Unit: s
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 0.005252
Metric #7
Descriptor:
     -> Name: otelcol_process_cpu_seconds
     -> Description: Total CPU user and system time in seconds
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 1.870000
Metric #8
Descriptor:
     -> Name: otelcol_process_memory_rss
     -> Description: Total physical memory (resident set size)
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 81424384.000000
Metric #9
Descriptor:
     -> Name: otelcol_processor_batch_batch_send_size
     -> Description: Number of units in the batch
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Count: 22
Sum: 486.000000
ExplicitBounds #0: 10.000000
ExplicitBounds #1: 25.000000
ExplicitBounds #2: 50.000000
ExplicitBounds #3: 75.000000
ExplicitBounds #4: 100.000000
ExplicitBounds #5: 250.000000
ExplicitBounds #6: 500.000000
ExplicitBounds #7: 750.000000
ExplicitBounds #8: 1000.000000
ExplicitBounds #9: 2000.000000
ExplicitBounds #10: 3000.000000
ExplicitBounds #11: 4000.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 6000.000000
ExplicitBounds #14: 7000.000000
ExplicitBounds #15: 8000.000000
ExplicitBounds #16: 9000.000000
ExplicitBounds #17: 10000.000000
ExplicitBounds #18: 20000.000000
ExplicitBounds #19: 30000.000000
ExplicitBounds #20: 50000.000000
ExplicitBounds #21: 100000.000000
Buckets #0, Count: 1
Buckets #1, Count: 9
Buckets #2, Count: 12
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Buckets #16, Count: 0
Buckets #17, Count: 0
Buckets #18, Count: 0
Buckets #19, Count: 0
Buckets #20, Count: 0
Buckets #21, Count: 0
Buckets #22, Count: 0
Metric #10
Descriptor:
     -> Name: otelcol_receiver_accepted_spans
     -> Description: Number of spans successfully pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 6.000000
Metric #11
Descriptor:
     -> Name: otelcol_rpc_server_requests_per_rpc
     -> Description: Measures the number of messages received per RPC. Should be 1 for all non-streaming RPCs.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Count: 1
Sum: 1.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #12
Descriptor:
     -> Name: otelcol_rpc_server_response_size
     -> Description: Measures size of RPC response messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Count: 1
Sum: 2.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #13
Descriptor:
     -> Name: otelcol_process_runtime_total_alloc_bytes
     -> Description: Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 30049032.000000
Metric #14
Descriptor:
     -> Name: otelcol_rpc_server_request_size
     -> Description: Measures size of RPC request messages (uncompressed).
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Count: 1
Sum: 900.000000
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 0
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 1
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #15
Descriptor:
     -> Name: otelcol_exporter_sent_spans
     -> Description: Number of spans successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 6.000000
Metric #16
Descriptor:
     -> Name: otelcol_process_runtime_total_sys_memory_bytes
     -> Description: Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 33133832.000000
Metric #17
Descriptor:
     -> Name: otelcol_processor_batch_timeout_trigger_send
     -> Description: Number of times the batch was sent due to a timeout trigger
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> processor: Str(batch)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 22.000000
Metric #18
Descriptor:
     -> Name: otelcol_receiver_refused_metric_points
     -> Description: Number of metric points that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(prometheus)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(http)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 0.000000
Metric #19
Descriptor:
     -> Name: otelcol_exporter_send_failed_spans
     -> Description: Number of spans in failed attempts to send to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 0.000000
Metric #20
Descriptor:
     -> Name: otelcol_rpc_server_duration
     -> Description: Measures the duration of inbound RPC.
     -> Unit: 
     -> DataType: Histogram
     -> AggregationTemporality: Cumulative
HistogramDataPoints #0
Data point attributes:
     -> rpc_grpc_status_code: Str(0)
     -> rpc_method: Str(Export)
     -> rpc_service: Str(opentelemetry.proto.collector.trace.v1.TraceService)
     -> rpc_system: Str(grpc)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Count: 1
Sum: 4.216575
ExplicitBounds #0: 0.000000
ExplicitBounds #1: 5.000000
ExplicitBounds #2: 10.000000
ExplicitBounds #3: 25.000000
ExplicitBounds #4: 50.000000
ExplicitBounds #5: 75.000000
ExplicitBounds #6: 100.000000
ExplicitBounds #7: 250.000000
ExplicitBounds #8: 500.000000
ExplicitBounds #9: 750.000000
ExplicitBounds #10: 1000.000000
ExplicitBounds #11: 2500.000000
ExplicitBounds #12: 5000.000000
ExplicitBounds #13: 7500.000000
ExplicitBounds #14: 10000.000000
Buckets #0, Count: 0
Buckets #1, Count: 1
Buckets #2, Count: 0
Buckets #3, Count: 0
Buckets #4, Count: 0
Buckets #5, Count: 0
Buckets #6, Count: 0
Buckets #7, Count: 0
Buckets #8, Count: 0
Buckets #9, Count: 0
Buckets #10, Count: 0
Buckets #11, Count: 0
Buckets #12, Count: 0
Buckets #13, Count: 0
Buckets #14, Count: 0
Buckets #15, Count: 0
Metric #21
Descriptor:
     -> Name: scrape_samples_post_metric_relabeling
     -> Description: The number of samples remaining after metric relabeling was applied
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 132.000000
Metric #22
Descriptor:
     -> Name: otelcol_processor_batch_metadata_cardinality
     -> Description: Number of distinct metadata value combinations being processed
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 1.000000
Metric #23
Descriptor:
     -> Name: otelcol_receiver_refused_spans
     -> Description: Number of spans that could not be pushed into the pipeline.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> receiver: Str(otlp)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
     -> transport: Str(grpc)
StartTimestamp: 2024-03-06 07:20:56.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 0.000000
Metric #24
Descriptor:
     -> Name: otelcol_exporter_sent_metric_points
     -> Description: Number of metric points successfully sent to destination.
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> exporter: Str(debug)
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:36.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 480.000000
Metric #25
Descriptor:
     -> Name: otelcol_process_runtime_heap_alloc_bytes
     -> Description: Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')
     -> Unit: 
     -> DataType: Gauge
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 1970-01-01 00:00:00 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 12823656.000000
Metric #26
Descriptor:
     -> Name: otelcol_process_uptime
     -> Description: Uptime of the process
     -> Unit: 
     -> DataType: Sum
     -> IsMonotonic: true
     -> AggregationTemporality: Cumulative
NumberDataPoints #0
Data point attributes:
     -> service_instance_id: Str(faef3727-9eea-43f0-bd98-b5970279b9ed)
     -> service_name: Str(otelcol)
     -> service_version: Str(0.96.0)
StartTimestamp: 2024-03-06 07:19:26.909 +0000 UTC
Timestamp: 2024-03-06 07:22:56.914 +0000 UTC
Value: 222.496549
	{"kind": "exporter", "data_type": "metrics", "name": "debug"}
