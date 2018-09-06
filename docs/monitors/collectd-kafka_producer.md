<!--- GENERATED BY gomplate from scripts/docs/monitor-page.md.tmpl --->

# collectd/kafka_producer

 Monitors a Java based Kafka producer using GenericJMX.

See the [integration documentation](https://github.com/signalfx/integrations/tree/master/collectd-kafka)
for more information.

This monitor has a set of [built in MBeans
configured](https://github.com/signalfx/signalfx-agent/tree/master/internal/monitors/collectd/kafkaproducer/mbeans.go)
for which it pulls metrics from the Kafka producer's JMX endpoint.

Sample YAML configuration:
```yaml
monitors:
  - type: collectd/kafka_producer
    host: localhost
    port: 8099
```

Note that this monitor requires Kafka v0.9.0.0 or above and collects metrics from the new producer API.


Monitor Type: `collectd/kafka_producer`

[Monitor Source Code](https://github.com/signalfx/signalfx-agent/tree/master/internal/monitors/collectd/kafkaproducer)

**Accepts Endpoints**: **Yes**

**Multiple Instances Allowed**: Yes

## Configuration

| Config option | Required | Type | Description |
| --- | --- | --- | --- |
| `host` | **yes** | `string` | Host to connect to -- JMX must be configured for remote access and accessible from the agent |
| `port` | **yes** | `integer` | JMX RMI port on the host |
| `name` | no | `string` |  |
| `serviceName` | no | `string` | This is how the service type is identified in the SignalFx UI so that you can get built-in content for it.  For custom JMX integrations, it can be set to whatever you like and metrics will get the special property `sf_hostHasService` set to this value. |
| `serviceURL` | no | `string` | The JMX connection string.  This is rendered as a Go template and has access to the other values in this config. NOTE: under normal circumstances it is not advised to set this string directly - setting the host and port as specified above is preferred. (**default:** `service:jmx:rmi:///jndi/rmi://{{.Host}}:{{.Port}}/jmxrmi`) |
| `instancePrefix` | no | `string` |  |
| `username` | no | `string` |  |
| `password` | no | `string` |  |
| `customDimensions` | no | `map of string` | Takes in key-values pairs of custom dimensions at the connection level. |
| `mBeansToCollect` | no | `list of string` | A list of the MBeans defined in `mBeanDefinitions` to actually collect. If not provided, then all defined MBeans will be collected. |
| `mBeansToOmit` | no | `list of string` | A list of the MBeans to omit. This will come handy in cases where only a few MBeans need to omitted from the default list |
| `mBeanDefinitions` | no | `map of object (see below)` | Specifies how to map JMX MBean values to metrics.  If using a specific service monitor such as cassandra, kafka, or activemq, they come pre-loaded with a set of mappings, and any that you add in this option will be merged with those.  See [collectd GenericJMX](https://collectd.org/documentation/manpages/collectd-java.5.shtml#genericjmx_plugin) for more details. |


The **nested** `mBeanDefinitions` config object has the following fields:

| Config option | Required | Type | Description |
| --- | --- | --- | --- |
| `objectName` | no | `string` |  |
| `instancePrefix` | no | `string` |  |
| `instanceFrom` | no | `list of string` |  |
| `values` | no | `list of object (see below)` |  |
| `dimensions` | no | `list of string` |  |


The **nested** `values` config object has the following fields:

| Config option | Required | Type | Description |
| --- | --- | --- | --- |
| `type` | no | `string` |  |
| `table` | no | `bool` |  (**default:** `false`) |
| `instancePrefix` | no | `string` |  |
| `instanceFrom` | no | `list of string` |  |
| `attribute` | no | `string` |  |




## Metrics

The following table lists the metrics available for this monitor. Metrics that are not marked as Custom are standard metrics and are monitored by default.

| Name | Type | Custom | Description |
| ---  | ---  | ---    | ---         |
| `kafka.producer.byte-rate` | gauge |  | Average number of bytes sent per second for a topic. This metric has client-id and topic dimensions. |
| `kafka.producer.compression-rate` | gauge |  | Average compression rate of record batches for a topic. This metric has client-id and topic dimensions. |
| `kafka.producer.io-wait-time-ns-avg` | gauge |  | Average length of time the I/O thread spent waiting for a socket ready for reads or writes in nanoseconds. This metric has client-id dimension. |
| `kafka.producer.outgoing-byte-rate` | gauge |  | Average number of outgoing bytes sent per second to all servers. This metric has client-id dimension. |
| `kafka.producer.record-error-rate` | gauge |  | Average per-second number of record sends that resulted in errors for a topic. This metric has client-id and topic dimensions. |
| `kafka.producer.record-retry-rate` | gauge |  | Average per-second number of retried record sends for a topic. This metric has client-id and topic dimensions. |
| `kafka.producer.record-send-rate` | gauge |  | Average number of records sent per second for a topic. This metric has client-id and topic dimensions. |
| `kafka.producer.request-latency-avg` | gauge |  | Average request latency in ms. Time it takes on average for the producer to get responses from the broker. This metric has client-id dimension. |
| `kafka.producer.request-rate` | gauge |  | Average number of requests sent per second. This metric has client-id dimension. |
| `kafka.producer.response-rate` | gauge |  | Average number of responses received per second. This metric has client-id dimension. |

To specify custom metrics you want to monitor, add a negated `metricsToExclude` to the monitor configuration, as shown in the code snippet below. The snippet lists all available custom metrics. You can copy and paste the snippet into your configuration file, then delete any custom metrics that you do not want to monitor. 
Note that some of the custom metrics require you to set a flag as well as add them to the list. Check the monitor configuration file to see if a flag is required for gathering additional metrics.
```yaml 
metricsToExclude:
  negated: true
```




