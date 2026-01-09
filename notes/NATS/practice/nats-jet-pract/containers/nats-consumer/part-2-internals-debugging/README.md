# Part 2: Consumer Internals & Debugging

**Time**: ~1 hour  
**Goal**: Understand ACK protocol, state machine, and debug like a senior engineer  
**Prerequisites**: [Part 1: Foundations](../part-1-foundations/README.md)

---

## 1. ACK Protocol Deep Dive

When you acknowledge a message, your client sends a message to a special **reply subject** embedded in the delivered message.

### Wire Format

```
$JS.ACK.<stream>.<consumer>.<deliveryCount>.<streamSeq>.<consumerSeq>.<timestamp>.<pending>
```

Example:
```
$JS.ACK.ORDERS.order-processor.1.42.15.1704067200000000000.5
         │       │              │  │  │  │                   │
         │       │              │  │  │  │                   └── 5 pending messages
         │       │              │  │  │  └── timestamp (ns)
         │       │              │  │  └── consumer sequence 15
         │       │              │  └── stream sequence 42
         │       │              └── 1st delivery attempt
         │       └── consumer name
         └── stream name
```

> **🎓 Feynman Box**
> 
> *"The secret handshake between your app and the server. The subject itself contains all the tracking info."*

This is why **you never construct ACK subjects manually** - the client library handles it.

---

## 2. The 5 ACK Types

### +ACK (Positive Acknowledgment)

```go
msg.Ack()
```

"I processed this message successfully. Don't send it again."

### -NAK (Negative Acknowledgment)

```go
msg.Nak()  // Immediate redelivery
```

"I can't process this right now. Send it again **immediately**."

⚠️ **Warning**: Can cause redelivery storms! Use `NakWithDelay` instead.

### -NAK with Delay

```go
msg.NakWithDelay(5 * time.Second)
```

"I can't process this. Wait 5 seconds, then try again."

### +WPI (Work in Progress)

```go
msg.InProgress()
```

"I'm still working on it. Reset the AckWait timer."

Use for long-running tasks to prevent premature redelivery.

### +TERM (Terminate)

```go
msg.Term()
```

"Stop trying to deliver this message. It's a poison pill."

Message stays in stream but won't be redelivered to this consumer.

```
                              ┌─────────────┐
                              │   Message   │
                              │  Delivered  │
                              └──────┬──────┘
                                     │
           ┌─────────────────────────┼─────────────────────────┐
           │                         │                         │
           ▼                         ▼                         ▼
    ┌──────────┐              ┌──────────┐              ┌──────────┐
    │   Ack    │              │   Nak    │              │   Term   │
    │  (+ACK)  │              │  (-NAK)  │              │ (+TERM)  │
    └────┬─────┘              └────┬─────┘              └────┬─────┘
         │                         │                         │
         ▼                         ▼                         ▼
    ┌──────────┐              ┌──────────┐              ┌──────────┐
    │ Complete │              │ Redeliver│              │  Dead    │
    │ (removed │              │  Queue   │              │ (stays   │
    │ from     │              │          │              │ in stream│
    │ pending) │              │          │              │ forever) │
    └──────────┘              └──────────┘              └──────────┘
```

> **🎓 Feynman Box**
> 
> *"Five ways to respond to the waiter:*
> - *Got it! (Ack)*
> - *Not now, come back immediately! (Nak)*
> - *Not now, come back in 5 minutes! (NakWithDelay)*
> - *Still eating, don't rush me! (InProgress)*
> - *Take it away, I'll never eat this! (Term)"*

**See**: [examples/01-ack-types](./examples/01-ack-types/main.go)

---

## 3. Consumer State Machine

The server maintains these key fields for each consumer:

```
┌────────────────────────────────────────────────────────────────┐
│                    CONSUMER STATE                              │
├────────────────────────────────────────────────────────────────┤
│                                                                │
│  Sequence Tracking:                                            │
│  ┌─────────────────────────────────────────────────────────┐  │
│  │ sseq  = 42    (next stream sequence to deliver)         │  │
│  │ dseq  = 15    (consumer delivery sequence)              │  │
│  │ adflr = 12    (ack delivery floor - all acked up to)    │  │
│  │ asflr = 40    (ack store floor - stream seq acked to)   │  │
│  └─────────────────────────────────────────────────────────┘  │
│                                                                │
│  Pending Map: (unacked messages)                               │
│  ┌─────────────────────────────────────────────────────────┐  │
│  │ streamSeq:41 → {consumerSeq:13, timestamp:1704067200}   │  │
│  │ streamSeq:42 → {consumerSeq:14, timestamp:1704067201}   │  │
│  └─────────────────────────────────────────────────────────┘  │
│                                                                │
│  Redelivery Queue: [41, 42]  (needs redelivery)               │
│                                                                │
│  Redelivery Counts:                                            │
│  ┌─────────────────────────────────────────────────────────┐  │
│  │ streamSeq:41 → 2  (delivered twice, not yet acked)      │  │
│  └─────────────────────────────────────────────────────────┘  │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```

### Message Lifecycle

```
                    ┌───────────┐
                    │  Stream   │
                    │  Message  │
                    └─────┬─────┘
                          │
                          ▼
                    ┌───────────┐
                    │ Delivered │───────────────┐
                    │ (pending  │               │
                    │   map)    │               │
                    └─────┬─────┘               │
                          │                     │
            ┌─────────────┼─────────────┐       │ AckWait
            │             │             │       │ timeout
            ▼             ▼             ▼       │
       ┌─────────┐  ┌───────────┐  ┌─────────┐ │
       │   Ack   │  │    Nak    │  │  Term   │ │
       └────┬────┘  └─────┬─────┘  └────┬────┘ │
            │             │             │       │
            ▼             ▼             ▼       │
       ┌─────────┐  ┌───────────┐  ┌─────────┐ │
       │ Removed │  │ Redeliver │  │ Removed │ │
       │  from   │  │   Queue   │◄─┘ (dead)  │ │
       │ pending │  │   (rdq)   │─────────────┘ │
       └─────────┘  └───────────┘               │
                          │                     │
                          └─────────────────────┘
```

> **🎓 Feynman Box**
> 
> *"The server keeps a scoreboard:*
> - *Pending map: who has the ball?*
> - *Redelivery queue: who dropped it?*
> - *Redelivery counts: how many chances left?"*

---

## 4. AckWait & BackOff

### AckWait

Default: **30 seconds**

If a message isn't acked within AckWait, it goes to the redelivery queue.

```go
AckWait: 60 * time.Second,
```

### BackOff

**Overrides AckWait entirely** when set. Provides exponential backoff.

```go
BackOff: []time.Duration{
    1 * time.Second,   // 1st retry after 1s
    5 * time.Second,   // 2nd retry after 5s
    30 * time.Second,  // 3rd retry after 30s
    5 * time.Minute,   // 4th+ retry after 5m
},
MaxDeliver: 5,
```

```
Delivery Timeline with BackOff:

    D1          D2      D3          D4              D5 (MaxDeliver)
    │           │       │           │               │
    ▼           ▼       ▼           ▼               ▼
────┼───────────┼───────┼───────────┼───────────────┼────────►
    │    1s     │  5s   │    30s    │      5m       │
    │◄─────────►│◄─────►│◄─────────►│◄─────────────►│
```

> **🎓 Feynman Box**
> 
> *"How long before the waiter checks on you again? With BackOff, they wait longer each time - maybe you need more time to decide."*

**See**: [examples/02-backoff-retry](./examples/02-backoff-retry/main.go)

---

## 5. MaxAckPending Flow Control

Default: **1000 messages**

When this limit is reached, the server **stops delivering** new messages until some are acked.

```
MaxAckPending = 5

Client state:        Server state:
                     
Pending: [1,2,3,4,5] ──► Limit reached!
                         No more deliveries
        │
        ▼
Ack message 1
        │
        ▼
Pending: [2,3,4,5]   ──► Room for 1 more
                         Message 6 delivered
        │
        ▼
Pending: [2,3,4,5,6] ──► Limit reached again
```

### Tuning Strategies

| Scenario | MaxAckPending | Reasoning |
|----------|---------------|-----------|
| Fast processing | 1000+ (default) | High throughput |
| Slow/external calls | 10-100 | Prevent memory buildup |
| Ordered processing | 1 | One at a time |
| Batch processing | Match batch size | Natural batching |

> **🎓 Feynman Box**
> 
> *"Only hold 10 plates at once. If you're still carrying 10, the kitchen stops giving you more until you put some down."*

**See**: [examples/03-maxackpending](./examples/03-maxackpending/main.go)

---

## 6. Common Bugs & Debugging

### Consumer Stall

**Symptom**: Consumer stops receiving messages but stream has data.

**Check**:
```bash
nats consumer info ORDERS order-processor
```

Look for:
```
Num Pending: 10000      # Messages waiting
Num Ack Pending: 1000   # = MaxAckPending! This is your problem
```

**Fix**: Either ack faster, increase MaxAckPending, or add more consumers.

### Redelivery Storm

**Symptom**: Same messages delivered repeatedly at high rate.

**Cause**: Using `msg.Nak()` without delay in a loop.

```go
// BAD - causes storm
for msg := range messages {
    if err := process(msg); err != nil {
        msg.Nak()  // Immediately redelivered, fails again, Nak again...
    }
}

// GOOD - controlled retry
for msg := range messages {
    if err := process(msg); err != nil {
        msg.NakWithDelay(5 * time.Second)
    }
}
```

### Double Ack Error

**Symptom**: `ErrMsgAlreadyAckd` error.

**Cause**: Acking the same message twice.

```go
// BAD
msg.Ack()
// ... later ...
msg.Ack()  // Error!

// GOOD - track ack state or use AckSync
if err := msg.DoubleAck(ctx); err != nil {
    // Handle - might already be acked
}
```

### Debugging Commands

```bash
# Consumer info (most useful)
nats consumer info ORDERS order-processor

# Key metrics to check:
# - Num Pending: messages waiting to be delivered
# - Num Ack Pending: delivered but not acked
# - Num Redelivered: how many redeliveries happened
# - Num Waiting: pull requests waiting (pull consumer)

# Watch real-time
nats consumer report ORDERS

# Stream context
nats stream info ORDERS
```

**See**: [examples/04-consumer-info](./examples/04-consumer-info/main.go)

---

## 7. Monitoring Metrics

### Key Metrics to Track

| Metric | Healthy | Warning | Critical |
|--------|---------|---------|----------|
| NumPending | < 1000 | 1000-10000 | > 10000 |
| NumAckPending | < MaxAckPending * 0.8 | > 0.8 | = MaxAckPending |
| Redelivery Rate | < 1% | 1-5% | > 5% |

### Alerting Thresholds

```go
info, _ := consumer.Info(ctx)

if info.NumAckPending >= maxAckPending * 0.9 {
    alert("Consumer approaching MaxAckPending limit")
}

if info.NumPending > 10000 {
    alert("Consumer backlog growing")
}

redeliveryRate := float64(info.NumRedelivered) / float64(info.Delivered.Consumer)
if redeliveryRate > 0.05 {
    alert("High redelivery rate: check processing failures")
}
```

---

## Summary: Debug Checklist

When a consumer isn't working:

1. **Check NumAckPending** - Are you at MaxAckPending limit?
2. **Check NumPending** - Is the backlog growing?
3. **Check redelivery count** - Are messages being retried excessively?
4. **Check AckWait** - Is it shorter than your processing time?
5. **Check MaxDeliver** - Have messages exceeded retry limit?

---

## Cross-References

- **RAFT Replication**: Consumer state is replicated via RAFT. See [[../../nats/modules/hour1/NATS_H1_THEORY|Hour 1 RAFT Theory]]
- **Cluster Testing**: Test consumer failover. See [[../../nats/modules/hour4/NATS_H4_PROGRESS|Hour 4 Chaos Testing]]

---

## Next Steps

1. Complete the [exercises](./exercises.md)
2. Explore the [NATS Micro module](../../nats-micro/README.md) for request/reply patterns
