
# NATS Monitoring and Observability Guide (Hour 5)

This guide provides a comprehensive overview of how to set up a robust monitoring and observability pipeline for your NATS cluster. After building a battle-tested cluster in Hour 4, this is the final step to prepare for a production deployment.

## 1. Observability Strategy for NATS

A successful NATS deployment relies on a proactive observability strategy. Our approach is built on three pillars:

- **Metrics:** Time-series data collected from the NATS server via Prometheus. This gives us quantitative insights into performance, load, and health.
- **Logging:** Structured logs provide detailed, event-driven context for debugging and understanding server behavior.
- **Tracing:** (Advanced) While not covered in this guide, distributed tracing can be implemented to track individual requests as they traverse the NATS ecosystem and client applications.

We will focus on a combination of Prometheus for metrics and Grafana for visualization, which is a powerful, open-source, and industry-standard stack for monitoring.

## 2. Prometheus Configuration and Setup

NATS exposes a Prometheus metrics endpoint out-of-the-box. We just need to configure Prometheus to scrape it.

### Enabling the Monitoring Endpoint in NATS

First, ensure your NATS server configuration includes the monitoring endpoint. It's typically enabled by default on port `8222`. If you need to explicitly configure it, add the following to your `nats-server.conf`:

```conf
# HTTP/Websocket monitoring port
http: 8222
```

### Prometheus Configuration

Next, create a `prometheus.yml` file to instruct Prometheus where to find your NATS servers. For our 3-node cluster, the configuration will list all nodes.

**File: `monitoring/prometheus.yml`**

```yaml
# See monitoring/prometheus.yml for the full configuration
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'nats'
    static_configs:
      - targets: ['nats-1:8222', 'nats-2:8222', 'nats-3:8222']
```

This configuration tells Prometheus to scrape the `/metrics` endpoint on each NATS node every 15 seconds.

To run Prometheus, you can use Docker:

```bash
docker run -d --name prometheus \
  -p 9090:9090 \
  -v $(pwd)/monitoring/prometheus.yml:/etc/prometheus/prometheus.yml \
  prom/prometheus
```

You can now access the Prometheus UI at `http://localhost:9090`.

## 3. Key Metrics and What They Mean

NATS exposes a wealth of metrics. Here are some of the most critical ones to monitor:

| Metric Name                  | Description                                                                 | Why it's important                                                              |
| ---------------------------- | --------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `nats_connections`           | Total number of active connections to the server.                           | Sudden drops can indicate client disconnects or network issues.                 |
| `nats_subscriptions`         | Total number of active subscriptions.                                       | A high number of subscriptions can impact performance.                          |
| `nats_in_msgs` / `nats_out_msgs` | Total number of incoming/outgoing messages.                               | Core indicator of message throughput.                                           |
| `nats_in_bytes` / `nats_out_bytes` | Total volume of incoming/outgoing data.                                   | Helps understand bandwidth usage and message payload sizes.                     |
| `nats_slow_consumers`        | Number of slow consumers detected.                                          | A critical indicator of clients that are not keeping up with message flow.      |
| `nats_js_memory`             | JetStream memory usage in bytes.                                            | Monitor to prevent out-of-memory errors.                                        |
| `nats_js_storage`            | JetStream storage usage in bytes.                                           | Monitor to prevent disk space exhaustion.                                       |
| `nats_js_ha_assets`          | Number of JetStream assets for which this server is the leader.             | In a cluster, ensures that leadership is balanced.                              |
| `nats_js_streams` / `nats_js_consumers` | Total number of JetStream streams and consumers.                 | Tracks the usage of JetStream resources.                                        |
| `nats_route_rtt_nanos`         | Round-trip time between clustered nodes.                                    | High RTT can indicate network latency issues between nodes, affecting cluster stability. |

## 4. Grafana Dashboard Creation

Grafana is the ideal tool to visualize the metrics collected by Prometheus. You can create a dashboard to get a single-pane-of-glass view of your cluster's health.

### Setup Grafana

Run Grafana using Docker and connect it to your Prometheus data source.

```bash
docker run -d --name grafana \
  -p 3000:3000 \
  grafana/grafana
```

1.  Access Grafana at `http://localhost:3000` (admin/admin).
2.  Add a new data source:
    -   **Type:** Prometheus
    -   **URL:** `http://<your_host_ip>:9090` (or `http://prometheus:9090` if using docker networking).
3.  Import the provided dashboard JSON.
    -   Navigate to `Dashboards` -> `Import`.
    -   Upload `monitoring/grafana-dashboard.json`.

A pre-built dashboard is an excellent starting point. The file `monitoring/grafana-dashboard.json` contains a basic dashboard definition to get you started.

**File: `monitoring/grafana-dashboard.json`**
```json
{
  "__inputs": [],
  "__requires": [],
  "annotations": {
    "list": []
  },
  "editable": true,
  "gnetId": null,
  "graphTooltip": 0,
  "id": null,
  "links": [],
  "panels": [
    {
      "title": "Connections",
      "type": "graph",
      "datasource": "Prometheus",
      "gridPos": { "h": 8, "w": 12, "x": 0, "y": 0 },
      "targets": [
        { "expr": "sum(nats_connections)" }
      ]
    },
    {
      "title": "Slow Consumers",
      "type": "graph",
      "datasource": "Prometheus",
      "gridPos": { "h": 8, "w": 12, "x": 12, "y": 0 },
      "targets": [
        { "expr": "sum(nats_slow_consumers)" }
      ]
    },
    {
      "title": "In/Out Messages",
      "type": "graph",
      "datasource": "Prometheus",
      "gridPos": { "h": 8, "w": 12, "x": 0, "y": 8 },
      "targets": [
        { "expr": "sum(rate(nats_in_msgs[5m]))", "legendFormat": "In" },
        { "expr": "sum(rate(nats_out_msgs[5m]))", "legendFormat": "Out" }
      ]
    }
  ],
  "schemaVersion": 16,
  "style": "dark",
  "tags": [],
  "templating": {
    "list": []
  },
  "time": {
    "from": "now-6h",
    "to": "now"
  },
  "timepicker": {
    "refresh_intervals": [ "5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d" ],
    "time_options": [ "5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d" ]
  },
  "timezone": "",
  "title": "NATS Cluster Monitoring",
  "version": 1
}
```

## 5. Alerting Rules and Thresholds

Alerting is crucial for proactively addressing issues. Prometheus Alertmanager can be configured to fire alerts based on rules.

Below are some recommended alerting rules to start with. These rules should be placed in a file like `monitoring/nats.rules.yml` and linked from your `prometheus.yml`.

**File: `monitoring/nats.rules.yml`**
```yaml
groups:
- name: nats_alerts
  rules:
  - alert: NatsServerDown
    expr: up{job="nats"} == 0
    for: 1m
    labels:
      severity: critical
    annotations:
      summary: "NATS server instance is down"
      description: "Instance {{ $labels.instance }} has been down for more than 1 minute."

  - alert: NatsSlowConsumers
    expr: sum(nats_slow_consumers) > 5
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "High number of slow consumers detected"
      description: "There are {{ $value }} slow consumers on the NATS cluster."

  - alert: NatsJetStreamMemoryHigh
    expr: (nats_js_memory / 1e9) > 1.5 # Example threshold: 1.5 GB
    for: 10m
    labels:
      severity: warning
    annotations:
      summary: "JetStream memory usage is high"
      description: "JetStream is using {{ $value }}GB of memory on {{ $labels.instance }}."

  - alert: NatsClusterNoLeader
    expr: nats_js_ha_assets == 0 and nats_js_streams > 0
    for: 2m
    labels:
      severity: critical
    annotations:
      summary: "NATS node has no RAFT leadership but streams exist"
      description: "Node {{ $labels.instance }} has no stream leadership, indicating a potential quorum issue."

  - alert: NatsRouteLatencyHigh
    expr: avg(nats_route_rtt_nanos) / 1e6 > 100 # > 100ms
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: "High latency between NATS cluster nodes"
      description: "Average route RTT is {{ $value }}ms, which may impact cluster performance."
```

## 6. Log Analysis and Troubleshooting

While metrics provide the "what," logs provide the "why." NATS server logs are your primary tool for deep troubleshooting.

### Log Aggregation

In a production environment, you should aggregate logs from all NATS nodes into a central logging platform (e.g., ELK stack, Loki, Splunk).

### Troubleshooting with Logs

- **Client Connectivity Issues:** Look for `Client connection created` and `Client connection closed` messages. Unexpected disconnects will be logged here.
- **Slow Consumers:** When a slow consumer is detected, NATS will log a message like `Slow consumer detected`. This log entry will include the client's name, IP, and the subject it's subscribed to.
- **JetStream Issues:** Look for logs related to RAFT (`[RAFT]`), stream operations, and storage. Errors like `file store: no space left on device` are critical.
- **Authentication/Authorization:** Failed connection attempts due to auth issues will be logged with messages like `Authentication timeout` or `Authorization violation`.

## 7. Performance Tuning with Metrics

Use your Grafana dashboard to guide performance tuning efforts:

- **High `nats_slow_consumers`:**
    -   Are your client applications processing messages efficiently?
    -   Can you increase the number of concurrent processors on the client-side?
    -   Is the message payload unusually large? Check `nats_in_bytes`.
- **High `nats_js_memory` or `nats_js_storage`:**
    -   Are your stream retention policies too aggressive? Consider `Limits` or `Interest` based retention.
    -   Are consumers acknowledging messages properly? Unacknowledged messages remain in the stream.
- **High `nats_route_rtt_nanos`:**
    -   Investigate network performance between your cluster nodes.
    -   Ensure nodes are geographically co-located or have low-latency links.

## 8. Production Monitoring Checklist

Before going to production, ensure you can answer "yes" to all of the following:

- [ ] Is a Prometheus instance scraping all NATS nodes?
- [ ] Is a Grafana dashboard set up with key metrics?
- [ ] Are alerting rules configured in Alertmanager for critical conditions (server down, slow consumers, etc.)?
- [ ] Is there a centralized logging solution in place?
- [ ] Has a performance baseline been established during normal load?
- [ ] Does the operations team know how to interpret the dashboards and logs to troubleshoot issues?
- [ ] Are the monitoring systems themselves monitored?

By following this guide, you can build a powerful observability solution that provides the confidence needed to run NATS in production.
