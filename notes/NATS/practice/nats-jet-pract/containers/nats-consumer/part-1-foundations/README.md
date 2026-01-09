# Part 1: Consumer Foundations

**Time**: ~1 hour  
**Goal**: Understand consumer types, configuration, and basic usage patterns

---

## 1. What is a Consumer?

A consumer is a **stateful view** of a stream. It tracks:
- Which messages have been **delivered** to clients
- Which messages have been **acknowledged**
- Which messages need **redelivery**

```
┌─────────────────────────────────────────────────────────┐
│                      STREAM                             │
│  ┌─────┬─────┬─────┬─────┬─────┬─────┬─────┬─────┐     │
│  │ M1  │ M2  │ M3  │ M4  │ M5  │ M6  │ M7  │ M8  │     │
│  └─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┘     │
│         │                   ▲                           │
│         │                   │                           │
│         ▼                   │                           │
│  ┌──────────────────────────┴────────────────────┐     │
│  │              CONSUMER STATE                    │     │
│  │  • Delivered through: M5                       │     │
│  │  • Acknowledged through: M3                    │     │
│  │  • Pending: M4, M5                             │     │
│  │  • Redelivery queue: [M4]                      │     │
│  └───────────────────────────────────────────────┘     │
└─────────────────────────────────────────────────────────┘
```

> **🎓 Feynman Box**
> 
> *"A consumer is like a bookmark that remembers what you've read AND can re-show pages you didn't finish."*
> 
> Unlike a simple bookmark (which just marks a position), a consumer actively tracks whether you actually processed each page.

### Core NATS vs JetStream Consumers

| Aspect | Core NATS | JetStream Consumer |
|--------|-----------|-------------------|
| Delivery guarantee | At-most-once | At-least-once |
| Message replay | Not possible | Yes (from any point) |
| Acknowledgment | None | Explicit/All/None |
| State | Stateless | Stateful (server-side) |

---

## 2. Pull vs Push Consumers

### Pull Consumer (Recommended)

Client **requests** messages when ready.

```
Client                          Server
  │                               │
  │──── "Give me 10 messages" ───►│
  │                               │
  │◄─── Messages 1-10 ────────────│
  │                               │
  │──── ACK messages 1-10 ───────►│
  │                               │
  │──── "Give me 10 more" ───────►│
  │                               │
```

**When to use:**
- Horizontal scaling (multiple workers)
- Batch processing
- Backpressure control
- New projects (recommended default)

### Push Consumer

Server **pushes** messages to a delivery subject.

```
Server                          Client (subscribed to "deliver.orders")
  │                               │
  │──── Message 1 ───────────────►│
  │──── Message 2 ───────────────►│
  │──── Message 3 ───────────────►│
  │                               │
  │◄─── ACK 1, ACK 2, ACK 3 ──────│
  │                               │
```

**When to use:**
- Simple replay scenarios
- Real-time streaming to single client
- Legacy migration from Core NATS

> **🎓 Feynman Box**
> 
> *"Pull = you order food when you're hungry. Push = conveyor belt sushi that keeps coming whether you're ready or not."*

### Comparison Table

| Aspect | Pull | Push |
|--------|------|------|
| Flow control | Client-driven | Server-driven (needs FlowControl) |
| Scaling | Easy (add more pullers) | Complex (queue groups) |
| Backpressure | Natural (stop pulling) | Need MaxAckPending |
| Complexity | Simple | More config needed |

**See**: [examples/01-pull-consumer](./examples/01-pull-consumer/main.go), [examples/02-push-consumer](./examples/02-push-consumer/main.go)

---

## 3. Durable vs Ephemeral Consumers

### Durable Consumer

- Has an explicit **name**
- State **persists** across restarts
- Survives client disconnects
- Can be shared by multiple clients

```go
consumer, _ := js.CreateOrUpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
    Durable: "order-processor",  // <-- Named = Durable
    // ...
})
```

### Ephemeral Consumer

- System-generated name
- State in **memory only**
- Auto-deleted after `InactiveThreshold` (default: 5s)
- Single client use

```go
consumer, _ := js.CreateOrUpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
    // No Durable field = Ephemeral
    InactiveThreshold: 10 * time.Second,
})
```

> **🎓 Feynman Box**
> 
> *"Durable = permanent tattoo. Ephemeral = temporary henna that fades away."*

### When to Use Each

| Use Case | Choice |
|----------|--------|
| Production workers | Durable |
| One-time replay | Ephemeral |
| Debugging/testing | Ephemeral |
| Shared processing pool | Durable |
| Real-time dashboards | Ephemeral |

---

## 4. Delivery Policies

Where should the consumer **start reading** from?

```
Stream Timeline:
    ┌────────────────────────────────────────────────────┐
    │  M1   M2   M3   M4   M5   M6   M7   M8   M9   M10  │
    └────────────────────────────────────────────────────┘
         ▲                   ▲              ▲         ▲
         │                   │              │         │
    DeliverAll          ByStartSeq(5)   ByStartTime  DeliverNew
                                                  DeliverLast
```

| Policy | Starts From | Use Case |
|--------|-------------|----------|
| `DeliverAll` | First available message | Full replay, new consumer setup |
| `DeliverLast` | Most recent message | Tail -f behavior |
| `DeliverNew` | After consumer creation | Real-time only |
| `DeliverByStartSequence` | Specific sequence number | Resume from checkpoint |
| `DeliverByStartTime` | Specific timestamp | Time-based replay |
| `DeliverLastPerSubject` | Last per filtered subject | State snapshots |

> **🎓 Feynman Box**
> 
> *"Where do you want to start reading? Beginning of the book, last chapter, or a specific page number?"*

**See**: [examples/04-delivery-policies](./examples/04-delivery-policies/main.go)

---

## 5. ACK Policies

How should the server track acknowledgments?

### AckExplicit (Recommended)

Each message must be individually acknowledged.

```go
msg.Ack()  // Acknowledge this specific message
```

**Pros**: Fine-grained control, exactly what was processed is tracked
**Cons**: More network traffic

### AckAll

Acknowledging message N acknowledges all messages 1 through N.

```go
// After receiving M1, M2, M3, M4, M5
msg5.Ack()  // Acks M1, M2, M3, M4, M5
```

**Pros**: Efficient for batch processing
**Cons**: All-or-nothing, can't skip individual messages

### AckNone

No acknowledgments required. Server assumes delivery = processed.

```go
// No ack calls needed
```

**Pros**: Highest throughput
**Cons**: No redelivery on failure, at-most-once semantics

> **🎓 Feynman Box**
> 
> *"Sign for every package (Explicit), sign once for the whole batch (All), or no signature needed (None)."*

### Comparison

| Policy | Redelivery | Use Case |
|--------|------------|----------|
| Explicit | Yes, per message | Most applications |
| All | Yes, batch | Ordered batch processing |
| None | No | Metrics, logs, fire-and-forget |

**See**: [examples/03-ack-policies](./examples/03-ack-policies/main.go)

---

## 6. Filter Subjects

Server-side filtering of messages before delivery.

```
Stream subjects: orders.>
Stream contents: orders.us.new, orders.eu.new, orders.us.shipped, orders.eu.shipped

Consumer with filter "orders.us.>":
  → Only receives: orders.us.new, orders.us.shipped
```

### Single Filter

```go
consumer, _ := js.CreateOrUpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
    FilterSubject: "orders.us.>",
})
```

### Multiple Filters

```go
consumer, _ := js.CreateOrUpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
    FilterSubjects: []string{"orders.us.>", "orders.ca.>"},
})
```

> **🎓 Feynman Box**
> 
> *"A librarian who pre-filters books before handing them to you. You only see the categories you asked for."*

---

## 7. Key Configuration Options

### AckWait

How long before unacked message is redelivered.

```go
AckWait: 30 * time.Second,  // Default
```

**Too short**: Messages redelivered while still processing
**Too long**: Slow recovery from failures

### MaxDeliver

Maximum redelivery attempts per message.

```go
MaxDeliver: 3,  // Try 3 times, then give up
MaxDeliver: -1, // Infinite retries (default)
```

### MaxAckPending

Maximum unacknowledged messages in flight.

```go
MaxAckPending: 1000,  // Default
MaxAckPending: 100,   // Lower for slow processors
MaxAckPending: -1,    // Unlimited (careful!)
```

---

## Summary

| Concept | Recommendation |
|---------|----------------|
| Consumer type | **Pull** for most cases |
| Persistence | **Durable** for production |
| ACK policy | **Explicit** for reliability |
| Delivery | **DeliverAll** for new consumers |
| Filter | Use when processing subset |

---

## Next Steps

1. Run the examples in [./examples/](./examples/)
2. Complete the [exercises](./exercises.md)
3. Continue to [Part 2: Internals & Debugging](../part-2-internals-debugging/README.md)
