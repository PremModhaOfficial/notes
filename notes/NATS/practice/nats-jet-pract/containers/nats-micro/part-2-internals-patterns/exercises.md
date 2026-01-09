# Part 2 Exercises: Internals & Patterns

Complete these exercises after reading the [README](./README.md).

---

## Exercise 1: Error Handling Patterns (Guided)

**Goal**: Implement proper error responses using `req.Error()`.

### Setup

```bash
docker-compose up -d
```

### Task

Create a `ValidationService` (v1.0.0) with endpoint `validate.user` that:

1. Accepts JSON: `{"email": "...", "age": N}`
2. Returns validation errors using `req.Error()` with proper codes:
   - `400` - Invalid JSON
   - `422` - Validation failed (invalid email format, age < 0)
   - `500` - Internal error (simulate with specific input)

3. On success, returns: `{"valid": true, "message": "User data is valid"}`

### Hints

<details>
<summary>Click for hint: Error response format</summary>

```go
handler := micro.HandlerFunc(func(req micro.Request) {
    var user UserRequest
    if err := json.Unmarshal(req.Data(), &user); err != nil {
        req.Error("400", "Invalid JSON format", []byte(err.Error()))
        return
    }
    
    if !isValidEmail(user.Email) {
        req.Error("422", "Invalid email format", nil)
        return
    }
    
    if user.Age < 0 {
        req.Error("422", "Age must be non-negative", nil)
        return
    }
    
    req.RespondJSON(map[string]any{
        "valid":   true,
        "message": "User data is valid",
    })
})
```

</details>

### Verification

```bash
# Valid request
nats req validate.user '{"email": "test@example.com", "age": 25}'
# Response: {"valid": true, "message": "User data is valid"}

# Invalid JSON
nats req validate.user 'not json'
# Error 400: Invalid JSON format

# Validation failure
nats req validate.user '{"email": "invalid", "age": -5}'
# Error 422: Invalid email format

# Check stats show error counts
nats req '$SRV.STATS.ValidationService' ''
```

---

## Exercise 2: Custom StatsHandler (Guided)

**Goal**: Implement custom metrics collection using StatsHandler.

### Task

Create a service that tracks custom business metrics:

1. Service `OrderService` with endpoint `orders.create`
2. Custom StatsHandler that tracks:
   - Total order value processed
   - Orders by category (electronics, clothing, food)
   - Average order value

3. Expose metrics on a custom endpoint `orders.metrics`

### Hints

<details>
<summary>Click for hint: StatsHandler implementation</summary>

```go
type OrderMetrics struct {
    mu              sync.Mutex
    TotalValue      float64
    OrdersByCategory map[string]int
    OrderCount      int
}

var metrics = &OrderMetrics{
    OrdersByCategory: make(map[string]int),
}

func statsHandler(endpoint micro.Endpoint, req micro.Request) {
    // Parse the order to extract metrics
    var order Order
    if err := json.Unmarshal(req.Data(), &order); err == nil {
        metrics.mu.Lock()
        metrics.TotalValue += order.Value
        metrics.OrdersByCategory[order.Category]++
        metrics.OrderCount++
        metrics.mu.Unlock()
    }
}

// In service config:
config := micro.Config{
    Name:         "OrderService",
    Version:      "1.0.0",
    StatsHandler: statsHandler,
}
```

</details>

### Verification

```bash
# Create some orders
nats req orders.create '{"category": "electronics", "value": 599.99}'
nats req orders.create '{"category": "clothing", "value": 49.99}'
nats req orders.create '{"category": "electronics", "value": 1299.99}'

# Check custom metrics
nats req orders.metrics ''
# Expected:
# {
#   "total_value": 1949.97,
#   "orders_by_category": {"electronics": 2, "clothing": 1},
#   "average_value": 649.99
# }
```

---

## Exercise 3: Graceful Shutdown (Guided)

**Goal**: Implement proper shutdown with DoneHandler.

### Task

Create a service that:

1. Opens a "database connection" (simulated with a channel)
2. Processes requests that "write to database"
3. On shutdown:
   - Stops accepting new requests
   - Flushes pending writes
   - Closes the connection
   - Logs cleanup completion

### Hints

<details>
<summary>Click for hint: DoneHandler pattern</summary>

```go
// Simulated database
type FakeDB struct {
    pending chan string
    done    chan struct{}
}

func newFakeDB() *FakeDB {
    db := &FakeDB{
        pending: make(chan string, 100),
        done:    make(chan struct{}),
    }
    // Background writer
    go func() {
        for {
            select {
            case data := <-db.pending:
                fmt.Printf("Writing to DB: %s\n", data)
                time.Sleep(100 * time.Millisecond)
            case <-db.done:
                // Flush remaining
                for len(db.pending) > 0 {
                    data := <-db.pending
                    fmt.Printf("Flushing: %s\n", data)
                }
                fmt.Println("Database closed")
                return
            }
        }
    }()
    return db
}

func (db *FakeDB) Close() {
    close(db.done)
}

// Service setup
db := newFakeDB()

srv, _ := micro.AddService(nc, micro.Config{
    Name:    "DBWriter",
    Version: "1.0.0",
    DoneHandler: func(srv micro.Service) {
        fmt.Println("Service stopping, flushing database...")
        db.Close()
        fmt.Println("Cleanup complete")
    },
})
```

</details>

### Verification

```bash
# Send several requests
for i in {1..10}; do nats req write "Data $i" & done

# While processing, stop the service (Ctrl+C)
# Expected output:
# Writing to DB: Data 1
# Writing to DB: Data 2
# ...
# Service stopping, flushing database...
# Flushing: Data 8
# Flushing: Data 9
# Flushing: Data 10
# Database closed
# Cleanup complete
```

---

## Exercise 4: JetStream Integration (Guided)

**Goal**: Combine Micro service with JetStream consumer.

### Task

Create an `OrderProcessor` service that:

1. Exposes `orders.submit` endpoint (receives order, publishes to JetStream)
2. Runs a background JetStream consumer that processes orders
3. Exposes `orders.status` endpoint to check processing status
4. Uses proper correlation IDs

### Architecture

```
Client                 OrderProcessor Service
  │                           │
  │── orders.submit ─────────►│──► JetStream ORDERS stream
  │◄─── {order_id: "xyz"} ────│
  │                           │
  │                    Consumer│◄── Pull from ORDERS
  │                           │    (background goroutine)
  │                           │
  │── orders.status ─────────►│
  │   {order_id: "xyz"}       │
  │◄─── {status: "completed"} │
```

### Hints

<details>
<summary>Click for hint: Structure</summary>

```go
type OrderStatus struct {
    ID        string
    Status    string
    CreatedAt time.Time
    UpdatedAt time.Time
}

var statusStore = sync.Map{}

// Submit handler - publishes to JetStream
submitHandler := micro.HandlerFunc(func(req micro.Request) {
    orderID := nuid.Next()
    
    // Store initial status
    statusStore.Store(orderID, &OrderStatus{
        ID:        orderID,
        Status:    "pending",
        CreatedAt: time.Now(),
    })
    
    // Publish to JetStream
    js.Publish(ctx, "ORDERS.new", req.Data(), 
        jetstream.WithMsgID(orderID))
    
    req.RespondJSON(map[string]string{"order_id": orderID})
})

// Background consumer
go func() {
    consumer, _ := js.CreateOrUpdateConsumer(ctx, "ORDERS", ...)
    msgs, _ := consumer.Consume(func(msg jetstream.Msg) {
        orderID := msg.Headers().Get("Nats-Msg-Id")
        
        // Process order...
        time.Sleep(2 * time.Second)
        
        // Update status
        if status, ok := statusStore.Load(orderID); ok {
            s := status.(*OrderStatus)
            s.Status = "completed"
            s.UpdatedAt = time.Now()
        }
        
        msg.Ack()
    })
}()
```

</details>

### Verification

```bash
# Create stream
nats stream add ORDERS --subjects="ORDERS.>" --storage=memory

# Submit order
nats req orders.submit '{"item": "Widget", "qty": 5}'
# Response: {"order_id": "abc123"}

# Check status (immediately)
nats req orders.status '{"order_id": "abc123"}'
# Response: {"status": "pending"}

# Wait 3 seconds, check again
nats req orders.status '{"order_id": "abc123"}'
# Response: {"status": "completed"}
```

---

## Stretch Challenge 1: Circuit Breaker Pattern

**Goal**: Implement a circuit breaker for external service calls.

### Requirements

1. Service `ProxyService` with endpoint `external.call`
2. Wraps calls to an external (simulated) API
3. Circuit breaker states:
   - **Closed**: Normal operation
   - **Open**: Fail fast (after 5 consecutive failures)
   - **Half-Open**: Test with single request (after 30s cooldown)

4. Expose `circuit.status` endpoint showing current state

### Success Criteria

```bash
# Normal operation
nats req external.call '{"endpoint": "healthy"}'
# Response: {"data": "..."}

# After 5 failures
nats req external.call '{"endpoint": "failing"}'
# (5 times)
# Response: Error "Circuit open - service unavailable"

# After 30s cooldown
nats req external.call '{"endpoint": "healthy"}'
# Response: {"data": "...", "circuit": "half-open -> closed"}
```

---

## Stretch Challenge 2: Service Mesh Observability

**Goal**: Build a distributed tracing system.

### Requirements

1. Create 3 services: `Gateway`, `UserService`, `OrderService`
2. Gateway calls UserService and OrderService
3. Implement trace propagation via headers:
   ```
   X-Trace-ID: unique trace ID
   X-Span-ID: current span ID
   X-Parent-Span-ID: parent span ID
   ```

4. Each service logs spans with timing
5. Create a `tracing.collect` endpoint that returns full trace

### Expected Trace Output

```json
{
  "trace_id": "abc123",
  "spans": [
    {
      "span_id": "span1",
      "service": "Gateway",
      "operation": "process_request",
      "duration_ms": 150,
      "children": ["span2", "span3"]
    },
    {
      "span_id": "span2", 
      "parent": "span1",
      "service": "UserService",
      "operation": "get_user",
      "duration_ms": 45
    },
    {
      "span_id": "span3",
      "parent": "span1", 
      "service": "OrderService",
      "operation": "get_orders",
      "duration_ms": 80
    }
  ]
}
```

---

## Debug Reference Card

Quick commands for debugging micro services:

```bash
# Discover all services
nats req '$SRV.PING' '' --replies=0 --timeout=1s

# Get service info
nats req '$SRV.INFO.ServiceName' ''

# Get service stats
nats req '$SRV.STATS.ServiceName' ''

# Target specific instance
nats req '$SRV.STATS.ServiceName.instance-id' ''

# Test endpoint directly
nats req 'endpoint.subject' 'request data'

# Monitor service subjects
nats sub '$SRV.>'
```

---

## Solutions

Solutions are available in the [examples/](./examples/) directory:

- `01-error-handling/main.go` - Exercise 1
- `02-stats-handler/main.go` - Exercise 2
- `03-graceful-shutdown/main.go` - Exercise 3
- `04-jetstream-integration/main.go` - Exercise 4

---

## Next Steps

After completing these exercises:

1. Review the [nats-consumer module](../../nats-consumer/README.md) for JetStream patterns
2. Explore production deployment patterns in the NATS documentation
3. Cross-reference with [RAFT modules](../../nats/modules/) for clustering
