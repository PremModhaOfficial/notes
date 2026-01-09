# Part 1 Exercises: Consumer Foundations

Complete these exercises after reading the [README](./README.md).

---

## Exercise 1: Pull Consumer Basics (Guided)

**Goal**: Create a pull consumer that processes messages from a stream.

### Setup

1. Start the NATS server:
   ```bash
   docker-compose up -d
   ```

2. Create a test stream:
   ```bash
   nats stream add ORDERS --subjects="orders.>" --storage=file --retention=limits --max-msgs=-1 --max-bytes=-1 --max-age=1h --max-msg-size=-1 --discard=old --dupe-window=2m --replicas=1
   ```

3. Publish test messages:
   ```bash
   for i in {1..10}; do nats pub orders.new "Order $i"; done
   ```

### Task

Write a Go program that:

1. Connects to NATS
2. Creates a durable pull consumer named `order-processor`
3. Fetches messages in batches of 5
4. Prints each message and acknowledges it
5. Stops when no more messages are available (use 2-second timeout)

### Hints

<details>
<summary>Click for hint 1: Consumer creation</summary>

```go
consumer, err := js.CreateOrUpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
    Durable:   "order-processor",
    AckPolicy: jetstream.AckExplicitPolicy,
})
```

</details>

<details>
<summary>Click for hint 2: Fetching messages</summary>

```go
msgs, err := consumer.Fetch(5, jetstream.FetchMaxWait(2*time.Second))
for msg := range msgs.Messages() {
    // process msg
}
```

</details>

### Verification

```bash
# Check consumer state after running
nats consumer info ORDERS order-processor

# Expected: 
# - Num Pending: 0 (all processed)
# - Delivered: 10
# - Ack Floor: sequence 10
```

---

## Exercise 2: Compare ACK Policies (Guided)

**Goal**: Understand the difference between AckExplicit, AckAll, and AckNone.

### Setup

Create 3 consumers with different ACK policies:

```bash
nats stream add EVENTS --subjects="events.>" --storage=memory --retention=limits --max-msgs=1000

# Publish 5 messages
for i in {1..5}; do nats pub events.test "Event $i"; done
```

### Task

Create 3 separate Go programs (or one program with 3 functions):

1. **ExplicitConsumer**: Uses `AckExplicitPolicy`, acks each message individually
2. **AllConsumer**: Uses `AckAllPolicy`, only acks the last message in batch
3. **NoneConsumer**: Uses `AckNonePolicy`, no acks needed

For each, fetch 5 messages and observe:
- How many network messages are sent for acks?
- What happens if processing fails mid-batch?

### Hints

<details>
<summary>Click for hint: AckAll usage</summary>

```go
// With AckAll, acking message 5 also acks 1-4
msgs := fetchMessages(5)
for i, msg := range msgs {
    process(msg)
    if i == len(msgs)-1 {
        msg.Ack()  // Only ack the last one
    }
}
```

</details>

### Verification

```bash
nats consumer info EVENTS explicit-consumer
nats consumer info EVENTS all-consumer
nats consumer info EVENTS none-consumer
```

Compare the `Ack Floor` and `Delivered` values.

---

## Exercise 3: Filter Subjects (Guided)

**Goal**: Create consumers that only receive a subset of stream messages.

### Setup

```bash
nats stream add LOGS --subjects="logs.>" --storage=memory

# Publish mixed logs
nats pub logs.app.info "App started"
nats pub logs.app.error "Connection failed"
nats pub logs.db.info "Query executed"
nats pub logs.db.error "Deadlock detected"
nats pub logs.cache.info "Cache hit"
nats pub logs.cache.error "Cache miss"
```

### Task

Create 2 consumers:

1. **ErrorsOnly**: Uses `FilterSubject: "logs.*.error"` - should receive 3 messages
2. **AppLogs**: Uses `FilterSubject: "logs.app.>"` - should receive 2 messages

### Verification

```bash
# ErrorsOnly should have:
nats consumer info LOGS errors-only
# Num Pending: 3

# AppLogs should have:
nats consumer info LOGS app-logs
# Num Pending: 2
```

---

## Exercise 4: Delivery Policies (Guided)

**Goal**: Start consuming from different positions in a stream.

### Setup

```bash
nats stream add METRICS --subjects="metrics.>" --storage=memory

# Publish 10 messages with 1-second gaps
for i in {1..10}; do 
    nats pub metrics.cpu "CPU: $i%"
    sleep 1
done
```

### Task

Create consumers with different delivery policies:

1. **FromStart**: `DeliverAll` - receives all 10 messages
2. **FromEnd**: `DeliverLast` - receives only message 10
3. **FromSeq5**: `DeliverByStartSequence: 5` - receives messages 5-10
4. **NewOnly**: `DeliverNew` - receives nothing (messages already published)

After creating NewOnly, publish a new message and verify it receives it.

### Hints

<details>
<summary>Click for hint: DeliverByStartSequence</summary>

```go
consumer, _ := js.CreateOrUpdateConsumer(ctx, "METRICS", jetstream.ConsumerConfig{
    Durable:              "from-seq5",
    DeliverPolicy:        jetstream.DeliverByStartSequencePolicy,
    OptStartSeq:          5,
})
```

</details>

---

## Stretch Challenge 1: Consumer Failover

**Goal**: Simulate consumer failure and observe redelivery.

### Task

1. Create a durable consumer with `AckWait: 10 * time.Second`
2. Fetch a message but **don't acknowledge it**
3. Exit the program
4. Run the program again - you should receive the same message

### Questions to Answer

- How does the server know to redeliver?
- What's in the "pending map"?
- What happens to the consumer sequence?

---

## Stretch Challenge 2: Build a Work Queue

**Goal**: Create a robust work queue with multiple workers.

### Requirements

1. Stream: `JOBS` with subjects `jobs.>`
2. Consumer: `job-worker` (durable, pull, explicit ack)
3. Run 3 instances of your worker simultaneously
4. Publish 100 jobs
5. Observe load balancing

### Success Criteria

- Each worker gets roughly 33 jobs
- No job is processed twice
- All 100 jobs complete

### Hints

- Use the same consumer name in all workers
- Print worker ID + job data for each processed job
- Track total jobs per worker

---

## Solutions

Solutions are available in the [examples/](./examples/) directory:

- `01-pull-consumer/main.go` - Exercise 1
- `03-ack-policies/main.go` - Exercise 2
- Additional solutions in `solutions/` (create your own first!)

---

## Next Steps

After completing these exercises:

1. Review your solutions against the examples
2. Proceed to [Part 2: Internals & Debugging](../part-2-internals-debugging/README.md)
3. Complete the Part 2 exercises for deeper understanding
