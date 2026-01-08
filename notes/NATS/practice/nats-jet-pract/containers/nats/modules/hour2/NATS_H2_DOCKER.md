# Hour 2: NATS Docker Setup Guide

Welcome to the hands-on portion of Hour 2! In this guide, you'll learn how to set up a NATS.io environment using Docker, progressing from a single node to a full 3-node cluster. This guide is designed to be a practical, step-by-step tutorial.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Step 1: Single NATS Node Setup](#step-1-single-nats-node-setup)
- [Step 2: 3-Node NATS Cluster Configuration](#step-2-3-node-nats-cluster-configuration)
- [Step 3: Docker Networking for NATS](#step-3-docker-networking-for-nats)
- [Step 4: Verification and Health Checks](#step-4-verification-and-health-checks)
- [Step 5: Troubleshooting Common Issues](#step-5-troubleshooting-common-issues)

---

### Prerequisites

Before we begin, ensure you have the following installed and running:

- [ ] **Docker:** `docker --version` (tested with Docker version 20.10.17)
- [ ] **Docker Compose:** `docker-compose --version` (tested with Docker Compose version 1.29.2)
- [ ] **Go:** `go version` (tested with go version go1.18.1 linux/amd64)

This guide assumes you have a basic understanding of Docker and have completed the theoretical components of [Hour 1](../hour1/NATS_H1_THEORY.md).

---

### Step 1: Single NATS Node Setup (Sanity Check)

First, let's start a single NATS server to ensure our Docker environment is working correctly.

- [ ] **Create a `docker-compose.yml` file:**

```yaml
# docker-compose.yml
version: '3.8'

services:
  nats:
    image: nats:2.9
    ports:
      - "4222:4222" # Client port
      - "8222:8222" # HTTP management port
    command: "-c /etc/nats/nats-server.conf"
    volumes:
      - ./nats-server.conf:/etc/nats/nats-server.conf
```

- [ ] **Create a `nats-server.conf` file:**

This configuration file enables the monitoring port.

```conf
# nats-server.conf
port: 4222
http_port: 8222
```

- [ ] **Start the NATS server:**

```bash
docker-compose up -d
```

- [ ] **Verify the single node is running:**

You can check the logs to see if the server started successfully.

```bash
docker-compose logs nats
```

You should see output similar to this, indicating the server is ready for connections:

```
[1] 2023/10/27 12:00:00.000000 [INF] Starting nats-server version 2.9.0
[1] 2023/10/27 12:00:00.000000 [INF] Git commit [commit-hash]
[1] 2023/10/27 12:00:00.000000 [INF] Listening for client connections on 0.0.0.0:4222
[1] 2023/10/27 12:00:00.000000 [INF] Server is ready
```

You can also inspect the server's monitoring endpoint:

```bash
curl localhost:8222/varz
```

This command should return a JSON object with server statistics.

- [ ] **Stop the single node before proceeding:**

```bash
docker-compose down
```

---

### Step 2: 3-Node NATS Cluster Configuration

Now, let's expand our setup to a 3-node cluster. This requires separate configuration files for each node and an updated `docker-compose.yml`.

- [ ] **Create configuration files for each node:**

**`nats-1.conf`**
```conf
# nats-1.conf
port: 4222
http_port: 8222

cluster {
  name: "my_cluster"
  listen: 0.0.0.0:6222
  routes = [
    nats-route://nats-2:6222,
    nats-route://nats-3:6222
  ]
}
```

**`nats-2.conf`**
```conf
# nats-2.conf
port: 4222
http_port: 8222

cluster {
  name: "my_cluster"
  listen: 0.0.0.0:6222
  routes = [
    nats-route://nats-1:6222,
    nats-route://nats-3:6222
  ]
}
```

**`nats-3.conf`**
```conf
# nats-3.conf
port: 4222
http_port: 8222

cluster {
  name: "my_cluster"
  listen: 0.0.0.0:6222
  routes = [
    nats-route://nats-1:6222,
    nats-route://nats-2:6222
  ]
}
```

**Explanation:**
- `cluster.name`: A unique name for your cluster. All nodes must share the same name.
- `cluster.listen`: The host and port where this node listens for connections from other nodes in the cluster.
- `cluster.routes`: An array of other nodes this node will try to connect to. We use the service names (`nats-1`, `nats-2`) defined in our `docker-compose.yml`.

- [ ] **Update `docker-compose.yml` for a 3-node cluster:**

```yaml
# docker-compose.yml
version: '3.8'

networks:
  nats_net:

services:
  nats-1:
    image: nats:2.9
    networks:
      - nats_net
    ports:
      - "4222:4222"
      - "8222:8222"
    command: "-c /etc/nats/nats-1.conf"
    volumes:
      - ./nats-1.conf:/etc/nats/nats-1.conf

  nats-2:
    image: nats:2.9
    networks:
      - nats_net
    ports:
      - "4223:4222"
      - "8223:8222"
    command: "-c /etc/nats/nats-2.conf"
    volumes:
      - ./nats-2.conf:/etc/nats/nats-2.conf

  nats-3:
    image: nats:2.9
    networks:
      - nats_net
    ports:
      - "4224:4222"
      - "8224:8222"
    command: "-c /etc/nats/nats-3.conf"
    volumes:
      - ./nats-3.conf:/etc/nats/nats-3.conf
```

- [ ] **Start the 3-node cluster:**

```bash
docker-compose up -d
```

---

### Step 3: Docker Networking for NATS

Understanding how Docker networking enables our NATS cluster is crucial.

- **Custom Network (`nats_net`):** We define a custom bridge network called `nats_net`. When services are attached to the same custom network, Docker's embedded DNS server allows them to resolve each other by their service name.

- **Service Resolution:**
  - In our configuration (`nats-1.conf`), the route `nats-route://nats-2:6222` works because `nats-1` can resolve the hostname `nats-2` to the internal IP address of the `nats-2` container.
  - This is why we don't need to worry about managing IP addresses manually.

- **Port Mapping:**
  - **Internal vs. External:** Inside the `nats_net` network, containers communicate directly on their internal ports (e.g., `6222` for clustering). The `ports` section in `docker-compose.yml` maps host ports to these container ports to allow external access (e.g., from your local machine).
  - **Avoiding Conflicts:** We map the internal `4222` and `8222` ports of `nats-2` and `nats-3` to different host ports (`4223`/`8223` and `4224`/`8224`) to avoid conflicts on the host machine.

---

### Step 4: Verification and Health Checks

With the cluster running, let's verify that the nodes have formed a cluster correctly.

- [ ] **Check the logs for each service:**

```bash
docker-compose logs nats-1
docker-compose logs nats-2
docker-compose logs nats-3
```
Look for messages indicating that routes have been established. For `nats-1`, you should see something like:
```
[1] 2023/10/27 12:15:00.000000 [INF] Listening for route connections on 0.0.0.0:6222
[1] 2023/10/27 12:15:01.000000 [INF] Established route to nats-2
[1] 2023/10/27 12:15:02.000000 [INF] Established route to nats-3
```

- [ ] **Check cluster status via the monitoring endpoint:**

The `/routesz` endpoint provides information about the cluster routes.

```bash
# Check nats-1
curl localhost:8222/routesz

# Check nats-2
curl localhost:8223/routesz

# Check nats-3
curl localhost:8224/routesz
```

The output for each node should show two routes listed, one for each of the other nodes in the cluster.

Example output from `curl localhost:8222/routesz`:
```json
{
  "server_id": "NA...",
  "now": "2023-10-27T12:20:00.123Z",
  "routes": [
    {
      "rid": 2,
      "remote_id": "NB...",
      "did_solicit": true,
      "is_configured": true,
      "ip": "172.19.0.4",
      "port": 6222,
      "rtt": "500µs",
      ...
    },
    {
      "rid": 3,
      "remote_id": "NC...",
      "did_solicit": true,
      "is_configured": true,
      "ip": "172.19.0.3",
      "port": 6222,
      "rtt": "600µs",
      ...
    }
  ],
  "num_routes": 2
}
```
A `num_routes` value of `2` on each node confirms a healthy 3-node cluster.

---

### Step 5: Troubleshooting Common Issues

If your cluster isn't forming correctly, here are some common issues to check:

1.  **"No Routes to Host" Error:**
    - **Cause:** One container cannot reach another. This is often a Docker networking issue.
    - **Solution:**
      - Ensure all services are on the same `networks` (`nats_net`).
      - Check for typos in the service names within the `.conf` files (e.g., `nats-route://nats-2:6222` must match the service name `nats-2`).
      - Use `docker network inspect nats_net` to see if all three containers are attached.

2.  **Cluster Name Mismatch:**
    - **Cause:** The `cluster.name` in the `.conf` files is not identical across all nodes.
    - **Solution:** Verify that the `cluster.name` is exactly the same in `nats-1.conf`, `nats-2.conf`, and `nats-3.conf`.

3.  **Port Conflicts on Host:**
    - **Cause:** Trying to map the same host port to multiple services.
    - **Solution:** Ensure you are using unique host ports for each service's client and monitoring ports, as shown in the example `docker-compose.yml`.

4.  **Incorrect Volume Mounts:**
    - **Cause:** The configuration files are not correctly mounted into the containers.
    - **Solution:** Double-check the `volumes` paths in `docker-compose.yml`. Ensure the local file names match the files you created (e.g., `./nats-1.conf`).

When you are finished, you can bring the entire cluster down with:
```bash
docker-compose down
```
