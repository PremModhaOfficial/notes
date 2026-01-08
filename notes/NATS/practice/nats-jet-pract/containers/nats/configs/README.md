# NATS Cluster Docker Configuration

This directory contains Docker configuration files for setting up NATS clusters with JetStream support.

## Files Overview

- **NATS_docker-compose.single.yml** - Single node NATS setup for testing and development
- **NATS_docker-compose.yml** - 3-node NATS cluster for high availability
- **NATS_nats-1.conf** - Configuration for the primary NATS node
- **NATS_nats-2.conf** - Configuration for the secondary NATS node  
- **NATS_nats-3.conf** - Configuration for the tertiary NATS node

## Quick Start

### Single Node (Testing)
```bash
docker-compose -f configs/NATS_docker-compose.single.yml up -d
```
Access: Client port 4222, Monitoring port 8222

### 3-Node Cluster (Production-ready)
```bash
docker-compose -f configs/NATS_docker-compose.yml up -d
```
Access: 
- Node 1: Client 4222, Monitoring 8222
- Node 2: Client 4223, Monitoring 8223  
- Node 3: Client 4224, Monitoring 8224

## Features

- **JetStream Enabled** - Persistent messaging with stream storage
- **Cluster Networking** - Full mesh topology between nodes
- **Health Checks** - Automated health monitoring for all nodes
- **Volume Persistence** - Data survives container restarts
- **Monitoring Ready** - HTTP endpoints for metrics and status
- **Production Settings** - Optimized for learning with production patterns

## Connection Examples

```bash
# Connect to single node
nats-cli --server=nats://localhost:4222 pub test.subject "Hello NATS!"

# Connect to cluster (will auto-discover other nodes)
nats-cli --server=nats://localhost:4222,nats://localhost:4223,nats://localhost:4224 stream ls
```

## Monitoring

Access monitoring dashboards:
- Single node: http://localhost:8222
- Cluster nodes: http://localhost:8222, http://localhost:8223, http://localhost:8224

## Authentication

Default credentials (change for production):
- Admin user: `admin` / `password`  
- Client user: `client` / `password`