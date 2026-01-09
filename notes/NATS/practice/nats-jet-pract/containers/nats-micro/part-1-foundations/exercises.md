# Part 1 Exercises: Micro Service Foundations

Complete these exercises after reading the [README](./README.md).

---

## Exercise 1: Basic Echo Service (Guided)

**Goal**: Create your first micro service that echoes requests.

### Setup

```bash
docker-compose up -d
```

### Task

Write a Go program that:

1. Connects to NATS
2. Creates a service named `EchoService` version `1.0.0`
3. Adds an endpoint on subject `echo` that returns the request data
4. Keeps running until interrupted (Ctrl+C)

### Hints

<details>
<summary>Click for hint 1: Service creation</summary>

```go
srv, err := micro.AddService(nc, micro.Config{
    Name:        "EchoService",
    Version:     "1.0.0",
    Description: "Echoes back requests",
})
```

</details>

<details>
<summary>Click for hint 2: Adding endpoint</summary>

```go
srv.AddEndpoint("echo", micro.HandlerFunc(func(req micro.Request) {
    req.Respond(req.Data())
}))
```

</details>

### Verification

```bash
# Test the service
nats req echo "Hello, NATS Micro!"
# Expected: "Hello, NATS Micro!"

# Check discovery
nats req '$SRV.PING' '' --replies=1

# Get service info
nats req '$SRV.INFO.EchoService' ''
```

---

## Exercise 2: Calculator with Groups (Guided)

**Goal**: Create a calculator service with organized endpoint groups.

### Task

Create a service `Calculator` (v1.0.0) with:

1. Group `math` containing:
   - Endpoint `add` - adds two numbers from JSON: `{"a": 5, "b": 3}` → `{"result": 8}`
   - Endpoint `multiply` - multiplies two numbers
   - Endpoint `divide` - divides (handle divide by zero!)

2. Group `convert` containing:
   - Endpoint `celsius-to-fahrenheit` - converts temperature
   - Endpoint `fahrenheit-to-celsius` - converts temperature

### Hints

<details>
<summary>Click for hint: Creating groups</summary>

```go
mathGroup := srv.AddGroup("math")
mathGroup.AddEndpoint("add", addHandler)

convertGroup := srv.AddGroup("convert")
convertGroup.AddEndpoint("celsius-to-fahrenheit", c2fHandler)
```

</details>

<details>
<summary>Click for hint: Parsing JSON request</summary>

```go
type MathRequest struct {
    A float64 `json:"a"`
    B float64 `json:"b"`
}

handler := micro.HandlerFunc(func(req micro.Request) {
    var r MathRequest
    if err := json.Unmarshal(req.Data(), &r); err != nil {
        req.Error("400", "Invalid JSON", nil)
        return
    }
    result := r.A + r.B
    req.RespondJSON(map[string]float64{"result": result})
})
```

</details>

### Verification

```bash
# Test math endpoints
nats req math.add '{"a": 10, "b": 5}'
# Expected: {"result": 15}

nats req math.divide '{"a": 10, "b": 0}'
# Expected: Error response

# Test convert endpoints  
nats req convert.celsius-to-fahrenheit '{"value": 100}'
# Expected: {"result": 212}

# Check all endpoints registered
nats req '$SRV.INFO.Calculator' ''
```

---

## Exercise 3: Load Balancing Demo (Guided)

**Goal**: Observe queue groups distributing requests across instances.

### Task

1. Create a service `Worker` with a single endpoint `process`
2. The handler should:
   - Print a unique worker ID (use a random string or process ID)
   - Sleep for 1 second (simulate work)
   - Return `{"worker": "<id>", "processed": true}`

3. Run **3 instances** of this service in separate terminals
4. Send 10 requests rapidly
5. Observe which workers handle which requests

### Hints

<details>
<summary>Click for hint: Unique worker ID</summary>

```go
import "github.com/nats-io/nuid"

workerID := nuid.Next()[:8]  // Short unique ID

handler := micro.HandlerFunc(func(req micro.Request) {
    fmt.Printf("[Worker %s] Processing request\n", workerID)
    time.Sleep(1 * time.Second)
    req.RespondJSON(map[string]any{
        "worker":    workerID,
        "processed": true,
    })
})
```

</details>

### Verification

```bash
# Terminal 1, 2, 3: Run workers
go run main.go

# Terminal 4: Send requests
for i in {1..10}; do nats req process "Job $i" & done; wait

# Expected: Requests distributed across workers
# Each worker should handle ~3-4 requests
```

---

## Exercise 4: Discovery Client (Guided)

**Goal**: Build a client that discovers and lists all services.

### Task

Write a Go program that:

1. Sends a request to `$SRV.PING` with a 1-second timeout
2. Collects all responses
3. For each service found, requests its stats from `$SRV.STATS.<name>.<id>`
4. Prints a summary table

### Expected Output

```
┌─────────────────┬──────────┬──────────────┬──────────┬────────┐
│ Service         │ Version  │ ID           │ Requests │ Errors │
├─────────────────┼──────────┼──────────────┼──────────┼────────┤
│ EchoService     │ 1.0.0    │ x3Yu...K7    │      150 │      0 │
│ Calculator      │ 1.0.0    │ y4Zv...L8    │       45 │      2 │
│ Worker          │ 1.0.0    │ z5Aw...M9    │       10 │      0 │
└─────────────────┴──────────┴──────────────┴──────────┴────────┘
```

### Hints

<details>
<summary>Click for hint: Collecting multiple responses</summary>

```go
inbox := nats.NewInbox()
sub, _ := nc.SubscribeSync(inbox)
nc.PublishRequest("$SRV.PING", inbox, nil)

// Collect responses for 1 second
timeout := time.After(1 * time.Second)
var responses []PingResponse
for {
    select {
    case <-timeout:
        goto done
    default:
        msg, err := sub.NextMsg(100 * time.Millisecond)
        if err != nil {
            continue
        }
        var ping PingResponse
        json.Unmarshal(msg.Data, &ping)
        responses = append(responses, ping)
    }
}
done:
```

</details>

---

## Stretch Challenge 1: Request Timeout Handler

**Goal**: Implement a service that handles slow operations gracefully.

### Requirements

1. Service `SlowAPI` with endpoint `fetch`
2. Request includes a timeout: `{"url": "...", "timeout_ms": 5000}`
3. If processing exceeds timeout, return error before client times out
4. Include processing time in response

### Success Criteria

```bash
# Fast response
nats req fetch '{"url": "http://fast.api", "timeout_ms": 5000}'
# {"data": "...", "duration_ms": 150}

# Timeout response
nats req fetch '{"url": "http://slow.api", "timeout_ms": 100}'
# Error: "Request timeout after 100ms"
```

---

## Stretch Challenge 2: Service Health Check

**Goal**: Create a health check aggregator.

### Requirements

1. Create multiple services (EchoService, Calculator, Worker)
2. Each service should have a `health` endpoint returning:
   ```json
   {
     "status": "healthy",
     "uptime_seconds": 3600,
     "memory_mb": 45.2
   }
   ```

3. Create a `HealthAggregator` service that:
   - Periodically pings all services (every 10s)
   - Maintains a map of service health
   - Exposes `status` endpoint returning overall system health

### Output Format

```json
{
  "overall": "healthy",
  "services": {
    "EchoService": {"status": "healthy", "instances": 2},
    "Calculator": {"status": "healthy", "instances": 1},
    "Worker": {"status": "degraded", "instances": 1, "expected": 3}
  },
  "checked_at": "2024-01-01T12:00:00Z"
}
```

---

## Solutions

Solutions are available in the [examples/](./examples/) directory:

- `01-basic-service/main.go` - Exercise 1
- `02-endpoints-groups/main.go` - Exercise 2
- `03-queue-groups/main.go` - Exercise 3
- `04-discovery/main.go` - Exercise 4

---

## Next Steps

After completing these exercises:

1. Review your solutions against the examples
2. Proceed to [Part 2: Internals & Patterns](../part-2-internals-patterns/README.md)
3. Explore how Micro integrates with JetStream consumers
