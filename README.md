# TCP Retransmission Monitoring System

A lightweight observability tool that tracks TCP packet retransmissions between microservices running inside a Kubernetes cluster, and exposes the metrics to Prometheus for monitoring and alerting.

---

## Overview

In distributed systems, TCP retransmissions are a key signal of network degradation, packet loss, or congestion between services. This project captures per-pod retransmission counters at the kernel level and surfaces them as Prometheus metrics, giving you real-time visibility into the health of your inter-service network communication.

## Architecture

```
┌─────────────────────────────────────────────────┐
│                Kubernetes Cluster               │
│                                                 │
│  ┌──────────┐   TCP    ┌──────────┐             │
│  │  Pod A   │◄────────►│  Pod B   │             │
│  └──────────┘          └──────────┘             │
│        │                     │                  │
│        └──────────┬──────────┘                  │
│                   ▼                             │
│         ┌─────────────────┐                     │
│         │  Monitoring     │                     │
│         │  Agent (Go)     │  ← reads kernel     │
│         │                 │    TCP stats        │
│         └────────┬────────┘                     │
│                  │ /metrics                     │
│                  ▼                              │
│         ┌─────────────────┐                     │
│         │   Prometheus    │                     │
│         └─────────────────┘                     │
└─────────────────────────────────────────────────┘
```

## Features

- Tracks TCP retransmission counts per microservice / pod pair
- Exposes metrics in Prometheus format via an HTTP `/metrics` endpoint
- Runs inside Kubernetes alongside your existing workloads
- Low overhead — reads directly from kernel network statistics
- Written in Go for minimal resource consumption

## Prerequisites

| Dependency   | Version       |
|--------------|---------------|
| Go           | 1.18+         |
| Ubuntu       | 16.04+        |
| Docker       | 20.10+        |
| Kubernetes   | 1.20+         |
| Prometheus   | 2.x           |

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/saeed-moradi-98/TCP_retransmission_monitoring_system.git
cd TCP_retransmission_monitoring_system
```

### 2. Build the binary

```bash
go build -o tcp-monitor .
```

### 3. Build the Docker image

```bash
docker build -t tcp-retransmission-monitor:latest .
```

### 4. Deploy to Kubernetes

Apply the provided manifests to deploy the monitoring agent as a DaemonSet (runs on every node):

```bash
kubectl apply -f k8s/
```

### 5. Configure Prometheus

Add the following scrape config to your `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: 'tcp-retransmission-monitor'
    kubernetes_sd_configs:
      - role: pod
    relabel_configs:
      - source_labels: [__meta_kubernetes_pod_label_app]
        action: keep
        regex: tcp-retransmission-monitor
```

## Metrics Reference

| Metric Name | Type | Description |
|---|---|---|
| `tcp_retransmission_total` | Counter | Total TCP retransmitted packets per pod |
| `tcp_retransmission_rate` | Gauge | Current retransmission rate (packets/sec) |

Metrics are exposed at `http://<node-ip>:<port>/metrics`.

## Configuration

The agent can be configured via environment variables:

| Variable | Default | Description |
|---|---|---|
| `SCRAPE_INTERVAL` | `15s` | How often to poll kernel TCP stats |
| `METRICS_PORT` | `9090` | Port to expose `/metrics` on |
| `LOG_LEVEL` | `info` | Log verbosity (`debug`, `info`, `warn`, `error`) |

## Project Structure

```
.
├── main.go              # Entry point
├── collector/           # Prometheus metric collectors
├── k8s/                 # Kubernetes manifests (Deployment, Service, RBAC)
├── Dockerfile
└── README.md
```

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Commit your changes: `git commit -m 'Add my feature'`
4. Push to the branch: `git push origin feature/my-feature`
5. Open a Pull Request

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
