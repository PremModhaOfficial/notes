# Part 2 Exercises: Internals & Debugging

Complete these exercises after reading the [README](./README.md).

---

## Exercise 1: ACK Types in Action (Guided)

**Goal**: Use all 5 ACK types and observe their effects.

### Setup

```bash
nats stream add TASKS --subjects="tasks.>" --storage=memory --retention=limits

# Publish test tasks
for i in {1..5}; do nats pub tasks.process "Task $i"; done
```

### Task

Write a Go program that processes tasks with different ACK responses:

1. **Task 1**: `msg.Ack()` - processed successfully
2. **Task 2**: `msg.Nak()` - temporary failure, immediate retry
3. **Task 3**: `msg.NakWithDelay(5*time.Second)` - delayed retry
4. **Task 4**: `msg.InProgress()` then `msg.Ack()` - long processing
5. **Task 5**: `msg.Term()` - poison pill, stop retrying

### Hints

<details>
<summary>Click for hint: Handling each task</summary>

```go
for msg := range msgs.Messages() {
    taskNum := extractTaskNum(string(msg.Data()))
    
    switch taskNum {
    case 1:
        fmt.Println("Task 1: Success")
        msg.Ack()
    case 2:
        fmt.Println("Task 2: Nak (immediate retry)")
        msg.Nak()
        // Note: You'll see this task again immediately!
    case 3:
        fmt.Println("Task 3: Nak with 5s delay")
        msg.NakWithDelay(5 * time.Second)
    case 4:
        fmt.Println("Task 4: Working...")
        msg.InProgress()
        time.Sleep(3 * time.Second)
        msg.Ack()
    case 5:
        fmt.Println("Task 5: Terminated")
        msg.Term()
    }
}
```

</details>

### Verification

```bash
# Watch consumer state during execution
watch -n 1 'nats consumer info TASKS task-processor'

# Look for:
# - Redelivered count increasing for Task 2
# - Pending messages changing
# - AckFloor advancing
```

---

## Exercise 2: BackOff Retry Strategy (Guided)

**Goal**: Implement exponential backoff for failing messages.

### Setup

```bash
nats stream add WEBHOOKS --subjects="webhooks.>" --storage=memory
nats pub webhooks.send '{"url": "http://failing-endpoint.local"}'
```

### Task

Create a consumer with BackOff configuration:

```go
consumer, _ := js.CreateOrUpdateConsumer(ctx, "WEBHOOKS", jetstream.ConsumerConfig{
    Durable:    "webhook-sender",
    AckPolicy:  jetstream.AckExplicitPolicy,
    BackOff: []time.Duration{
        1 * time.Second,   // 1st retry
        5 * time.Second,   // 2nd retry
        30 * time.Second,  // 3rd retry
    },
    MaxDeliver: 4,  // Original + 3 retries
})
```

Write a handler that:
1. Always fails (simulating a bad webhook endpoint)
2. Logs the delivery count and time between deliveries

### Verification

```bash
# Observe timing between deliveries
# Expected: ~1s, ~5s, ~30s gaps

nats consumer info WEBHOOKS webhook-sender
# Num Redelivered should increase
# Eventually hits MaxDeliver
```

---

## Exercise 3: MaxAckPending Flow Control (Guided)

**Goal**: Observe flow control in action.

### Setup

```bash
nats stream add SLOWSTREAM --subjects="slow.>" --storage=memory

# Publish 50 messages
for i in {1..50}; do nats pub slow.process "Slow job $i"; done
```

### Task

Create a consumer with `MaxAckPending: 5`:

```go
consumer, _ := js.CreateOrUpdateConsumer(ctx, "SLOWSTREAM", jetstream.ConsumerConfig{
    Durable:       "slow-processor",
    AckPolicy:     jetstream.AckExplicitPolicy,
    MaxAckPending: 5,
})
```

Write a processor that:
1. Fetches messages in batches
2. Simulates slow processing (2 seconds per message)
3. Logs when fetches return fewer messages than requested

### Expected Behavior

```
Fetched 5 messages
Processing message 1... (2s)
Processing message 2... (2s)
...
Fetched 5 messages  <- Only after acking previous batch
```

### Verification

```bash
# Monitor during slow processing
nats consumer info SLOWSTREAM slow-processor

# Num Ack Pending should hover at 5
# Num Pending shows backlog
```

---

## Exercise 4: Debug a Stuck Consumer (Guided)

**Goal**: Diagnose and fix a consumer that stops receiving messages.

### Setup

This exercise simulates a common production bug.

```bash
nats stream add BUGGY --subjects="buggy.>" --storage=memory

# Publish many messages
for i in {1..100}; do nats pub buggy.work "Work item $i"; done
```

### Task

Create a consumer with these (intentionally problematic) settings:

```go
consumer, _ := js.CreateOrUpdateConsumer(ctx, "BUGGY", jetstream.ConsumerConfig{
    Durable:       "buggy-worker",
    AckPolicy:     jetstream.AckExplicitPolicy,
    MaxAckPending: 10,
    AckWait:       5 * time.Second,
})
```

Write a processor that:
1. Fetches 10 messages
2. Processes each one in 1 second
3. **Bug**: Only acks every other message

After running, the consumer will appear "stuck". Your task:

1. Use `nats consumer info` to diagnose
2. Identify the problem
3. Fix the code

### Hints

<details>
<summary>Click for hint: What to look for</summary>

```bash
nats consumer info BUGGY buggy-worker

# Red flags:
# - Num Ack Pending = MaxAckPending (10)
# - Num Pending > 0 but no new deliveries
# - Num Redelivered increasing

# The fix: Ack ALL messages, not just some!
```

</details>

---

## Stretch Challenge 1: Implement a Dead Letter Queue

**Goal**: Handle poison messages that can't be processed.

### Requirements

1. Main stream: `ORDERS` 
2. Dead letter stream: `ORDERS_DLQ`
3. Consumer with `MaxDeliver: 3`
4. When a message exceeds MaxDeliver, move it to DLQ

### Implementation Hints

- Use `consumer.Info()` to check `NumRedelivered`
- On the 3rd delivery, publish to DLQ stream before calling `msg.Term()`
- Include metadata: original subject, error reason, timestamps

### Success Criteria

- Poison messages end up in DLQ
- Normal messages processed successfully
- No message loss

---

## Stretch Challenge 2: Build a Consumer Monitor

**Goal**: Create a monitoring dashboard for consumers.

### Requirements

1. List all consumers in a stream
2. For each consumer, display:
   - Name and type (durable/ephemeral)
   - Pending count
   - Ack pending count
   - Redelivery rate (%)
   - Health status (healthy/warning/critical)

### Health Thresholds

```go
type HealthStatus string
const (
    Healthy  HealthStatus = "healthy"   // AckPending < 80% of MaxAckPending
    Warning  HealthStatus = "warning"   // AckPending 80-95% of MaxAckPending
    Critical HealthStatus = "critical"  // AckPending > 95% OR redelivery rate > 5%
)
```

### Output Format

```
STREAM: ORDERS
┌─────────────────┬──────────┬─────────┬─────────────┬──────────┬──────────┐
│ Consumer        │ Type     │ Pending │ Ack Pending │ Redeliv% │ Health   │
├─────────────────┼──────────┼─────────┼─────────────┼──────────┼──────────┤
│ order-processor │ durable  │     150 │      45/100 │     1.2% │ healthy  │
│ audit-logger    │ durable  │       0 │       0/100 │     0.0% │ healthy  │
│ slow-worker     │ durable  │    5000 │     95/100  │    12.5% │ critical │
└─────────────────┴──────────┴─────────┴─────────────┴──────────┴──────────┘
```

---

## Debug Reference Card

Quick commands for debugging consumers:

```bash
# Consumer overview
nats consumer info STREAM CONSUMER

# Key fields to check:
# - Num Pending: messages waiting in stream
# - Num Ack Pending: delivered but not acked
# - Num Redelivered: redelivery count
# - Num Waiting: pull requests waiting (pull consumer)

# Watch in real-time
watch -n 1 'nats consumer info STREAM CONSUMER'

# All consumers report
nats consumer report STREAM

# Stream context
nats stream info STREAM

# Purge and start fresh (careful!)
nats consumer delete STREAM CONSUMER
```

---

## Next Steps

After completing these exercises:

1. Review the [examples/](./examples/) directory for solutions
2. Explore the [NATS Micro module](../../nats-micro/README.md) for request/reply patterns
3. Cross-reference with [RAFT modules](../../nats/modules/) for clustering behavior
