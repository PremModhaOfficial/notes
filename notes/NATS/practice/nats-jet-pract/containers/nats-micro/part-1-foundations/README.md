# Part 1: Micro Service Foundations

**Time**: ~1 hour  
**Goal**: Understand the NATS micro service framework basics  
**Prerequisites**: Basic NATS pub/sub, request/reply pattern

---

## 1. What is the Micro Framework?

The NATS micro package provides a **standardized way** to build services that:
- Register themselves for **automatic discovery**
- Expose **statistics** and **health info**
- Handle **request/reply** patterns consistently

```
┌─────────────────────────────────────────────────────────────────┐
│                         NATS SERVER                             │
│                                                                 │
│  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐         │
│  │  EchoSvc    │    │  MathSvc    │    │  UserSvc    │         │
│  │  (v1.0.0)   │    │  (v2.1.0)   │    │  (v1.2.0)   │         │
│  └──────┬──────┘    └──────┬──────┘    └──────┬──────┘         │
│         │                  │                  │                 │
│         └──────────────────┼──────────────────┘                 │
│                            │                                    │
│                     $SRV.PING ───► All services respond         │
│                     $SRV.INFO ───► Service details              │
│                     $SRV.STATS ──► Request counts, errors       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

> **🎓 Feynman Box**
> 
> *"A way for services to register with a phone book and answer calls. Anyone can dial the directory ($SRV) to find who's available."*

### Without Micro vs With Micro

| Aspect | Manual Request/Reply | Micro Framework |
|--------|---------------------|-----------------|
| Discovery | Custom implementation | Built-in $SRV protocol |
| Monitoring | Custom metrics | Auto stats collection |
| Versioning | Manual tracking | First-class support |
| Error handling | Ad-hoc | Standardized headers |

---

## 2. Service Configuration

### Basic Service

```go
srv, err := micro.AddService(nc, micro.Config{
    Name:        "EchoService",
    Version:     "1.0.0",
    Description: "Echoes back the request",
})
```

### Config Fields

| Field | Required | Description |
|-------|----------|-------------|
| Name | Yes | Service name (used in discovery) |
| Version | Yes | Semantic version (e.g., "1.0.0") |
| Description | No | Human-readable description |
| Endpoint | No | Default endpoint config |
| QueueGroup | No | Override default "q" |
| DoneHandler | No | Called when service stops |
| ErrorHandler | No | Called on NATS errors |

### Unique ID

Each service instance gets a unique ID (using `nuid`):

```
Service: EchoService
ID: x3Yuiq7g7MoxhXdxk7i4K7  ← Auto-generated
```

> **🎓 Feynman Box**
> 
> *"Your restaurant's Yelp profile: name, version, description, and a unique ID so customers can find you."*

---

## 3. Handlers & Requests

### Handler Interface

```go
type Handler interface {
    Handle(Request)
}

// Most common: use HandlerFunc
handler := micro.HandlerFunc(func(req micro.Request) {
    req.Respond(req.Data())
})
```

### Request Interface

```go
type Request interface {
    Respond(data []byte) error           // Send response
    RespondJSON(v any) error             // Send JSON response
    Error(code, desc string, data []byte) error  // Send error
    
    Data() []byte                        // Request payload
    Headers() nats.Header                // Request headers
    Subject() string                     // Subject received on
    Reply() string                       // Reply subject
}
```

### Simple Example

```go
echoHandler := micro.HandlerFunc(func(req micro.Request) {
    req.Respond(req.Data())  // Echo back the data
})

srv, _ := micro.AddService(nc, micro.Config{
    Name:    "EchoService",
    Version: "1.0.0",
    Endpoint: &micro.EndpointConfig{
        Subject: "echo",
        Handler: echoHandler,
    },
})
```

> **🎓 Feynman Box**
> 
> *"Menu items that customers can order. Each handler is a dish - customer sends order, you send back food."*

**See**: [examples/01-basic-service](./examples/01-basic-service/main.go)

---

## 4. Endpoints & Groups

### Adding Endpoints

```go
srv, _ := micro.AddService(nc, micro.Config{
    Name:    "Calculator",
    Version: "1.0.0",
})

// Add individual endpoints
srv.AddEndpoint("add", addHandler)       // Subject: "add"
srv.AddEndpoint("subtract", subHandler)  // Subject: "subtract"
```

### Using Groups

Groups add a **subject prefix** to organize endpoints:

```go
// Create a group
math := srv.AddGroup("math")

// Add endpoints to the group
math.AddEndpoint("add", addHandler)       // Subject: "math.add"
math.AddEndpoint("multiply", mulHandler)  // Subject: "math.multiply"

// Nested groups
advanced := math.AddGroup("advanced")
advanced.AddEndpoint("sqrt", sqrtHandler) // Subject: "math.advanced.sqrt"
```

```
Service: Calculator v1.0.0

Endpoints:
├── add                    (no group)
├── subtract               (no group)
└── math                   (group)
    ├── add               → math.add
    ├── multiply          → math.multiply
    └── advanced          (nested group)
        └── sqrt          → math.advanced.sqrt
```

> **🎓 Feynman Box**
> 
> *"Organizing your menu into sections: appetizers, mains, desserts. Groups = sections, Endpoints = dishes."*

**See**: [examples/02-endpoints-groups](./examples/02-endpoints-groups/main.go)

---

## 5. Queue Groups

Queue groups enable **load balancing** across service instances.

```
Request to "echo"
       │
       ▼
┌──────────────────────────────────────────┐
│           NATS Queue Group "q"           │
└──────────────────────────────────────────┘
       │
       ├────────────► Instance 1 (handles this one)
       │
       ├────────────► Instance 2 (idle)
       │
       └────────────► Instance 3 (idle)
```

### Default Behavior

All endpoints use queue group `"q"` by default.

### Customizing Queue Groups

```go
// Service level
srv, _ := micro.AddService(nc, micro.Config{
    Name:       "Worker",
    Version:    "1.0.0",
    QueueGroup: "workers",  // All endpoints use "workers"
})

// Endpoint level
srv.AddEndpoint("process", handler, 
    micro.WithEndpointQueueGroup("heavy-workers"))
```

### Disabling Queue Groups

For **broadcast** scenarios (all instances receive):

```go
srv.AddEndpoint("broadcast", handler,
    micro.WithEndpointQueueGroup(""))  // Empty = no queue group
```

> **🎓 Feynman Box**
> 
> *"Multiple waiters sharing the same section. Only one waiter handles each table's order - they don't all run over at once."*

**See**: [examples/03-queue-groups](./examples/03-queue-groups/main.go)

---

## 6. Discovery Protocol

### $SRV Subjects

The micro framework automatically subscribes to these control subjects:

```
$SRV.PING              ← All services respond
$SRV.PING.EchoService  ← Only EchoService instances respond
$SRV.PING.EchoService.x3Yuiq7g ← Only this specific instance

$SRV.INFO              ← Service configuration
$SRV.INFO.EchoService
$SRV.INFO.EchoService.x3Yuiq7g

$SRV.STATS             ← Request counts, errors, timing
$SRV.STATS.EchoService
$SRV.STATS.EchoService.x3Yuiq7g
```

### Discovery in Action

```bash
# Find all services
nats req '$SRV.PING' '' --replies=0 --timeout=1s

# Get info for specific service
nats req '$SRV.INFO.EchoService' ''

# Get stats
nats req '$SRV.STATS.EchoService' ''
```

### Response Examples

**PING Response:**
```json
{
  "name": "EchoService",
  "id": "x3Yuiq7g7MoxhXdxk7i4K7",
  "version": "1.0.0",
  "type": "io.nats.micro.v1.ping_response"
}
```

**STATS Response:**
```json
{
  "name": "EchoService",
  "id": "x3Yuiq7g7MoxhXdxk7i4K7",
  "version": "1.0.0",
  "started": "2024-01-01T00:00:00Z",
  "endpoints": [{
    "name": "echo",
    "subject": "echo",
    "num_requests": 150,
    "num_errors": 2,
    "average_processing_time": "1.5ms"
  }]
}
```

> **🎓 Feynman Box**
> 
> *"The broadcast system where services shout 'I'm here!' ($SRV.PING), 'Here's what I do!' ($SRV.INFO), and 'Here's how busy I am!' ($SRV.STATS)"*

**See**: [examples/04-discovery](./examples/04-discovery/main.go)

---

## Summary

| Concept | Key Point |
|---------|-----------|
| **Service** | Name + Version + Endpoints |
| **Handler** | Process request, send response |
| **Endpoint** | Subject + Handler |
| **Group** | Subject prefix for organization |
| **Queue Group** | Load balancing across instances |
| **$SRV Protocol** | Auto-discovery and monitoring |

---

## Next Steps

1. Run the examples in [./examples/](./examples/)
2. Complete the [exercises](./exercises.md)
3. Continue to [Part 2: Internals & Patterns](../part-2-internals-patterns/README.md)
