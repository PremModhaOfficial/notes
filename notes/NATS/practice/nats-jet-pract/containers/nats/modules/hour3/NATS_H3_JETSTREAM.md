# Hour 3: JetStream and RAFT Replication Deep Dive

Welcome back! In Hour 2, you successfully built a 3-node NATS cluster using Docker. Now, it's time to leverage that cluster to explore the power of JetStream, NATS' built-in persistence and streaming engine. This hour, we'll focus on how JetStream uses the underlying RAFT consensus protocol to provide highly available, fault-tolerant data streams.

## Progress Checklist

- [ ] Understand the core concepts of JetStream streams and consumers.
- [ ] Create a replicated stream with 3 replicas.
- [ ] Learn how each stream forms its own RAFT group.
- [ ] Configure durable consumers for fault-tolerant message processing.
- [ ] Use placement tags to control where streams are located.
- [ ] Monitor the health and status of your streams and their replicas.
- [ ] Perform basic backup and restore operations.

---

## 1. JetStream Concepts Recap

In Hour 1, we introduced the theoretical side of RAFT. Now, let's see it in action with JetStream.

*   **Stream:** A log of messages. Think of it as a persistent, append-only log that can be replayed. In our cluster, a stream's data will be replicated across multiple nodes for fault tolerance.
*   **Consumer:** A stateful view into a stream. Consumers let you read messages from a stream and, crucially, keep track of your progress. If your application disconnects and reconnects, a durable consumer knows exactly where you left off.
*   **Replication:** JetStream achieves fault tolerance by replicating stream data across multiple servers in the cluster. Each stream has its own independent RAFT consensus group.

---

## 2. Creating Your First Replicated Stream

With our 3-node cluster running, creating a replicated stream is simple. The key is the `--replicas` flag.

### Step 1: Create a Stream with 3 Replicas

Let's create a stream named `ORDERS` to store incoming orders. We'll tell JetStream to keep 3 copies of this stream's data.

```bash
nats stream add ORDERS --subjects "ORDERS.>" --replicas 3 --storage file
```

**Expected Output:**
```
? Stream Name ORDERS
? Subjects to consume ORDERS.>
? Storage backend file
? Retention Policy Limits
? Discard Policy Old
? Stream Messages Limit -1
? Message size limit -1
? Maximum message age limit -1
? Maximum individual message size -1
? Duplicate tracking time window 2m0s
? Allow message republishing No
? Allow direct access No
? Mirror other stream No
? Add sources to this stream No
? Re-publish message on different subject No
? Message TTL 0s
? Placement
  Cluster
  Tags
? Select a Cluster nats
✔ Stream ORDERS was created

Information for Stream ORDERS created 2026-01-07T12:05:58Z

Configuration:

             Subjects: ORDERS.>
          Replicas: 3
           Storage: File

State:

            Messages: 0
               Bytes: 0 B
...
```

### Step 2: Understanding the RAFT Group

When you created the `ORDERS` stream with `replicas=3`, JetStream automatically formed a new, independent RAFT group just for this stream. The three nodes in your cluster (`nats-1`, `nats-2`, `nats-3`) became the three peers in this group.

One of these nodes was elected the **leader** for the `ORDERS` stream. All writes (publishing messages) to the `ORDERS` stream must go through the leader. The leader then replicates the message to the other two **followers**. A message is only acknowledged back to the publisher once a quorum (2 out of 3 nodes) has confirmed it's been stored.

---

## 3. Stream Information and Leadership

You can inspect the stream's status to see which node is the current leader and which nodes are followers.

### Step 1: Get Stream Info

Use the `nats stream info` command:

```bash
nats stream info ORDERS
```

**Expected Output:**
```
Information for Stream ORDERS created 2026-01-07T12:05:58Z

...
Cluster:

                Name: nats
              Leader: nats-2  <-- The current leader for this stream
             Replicas:
                        nats-1 (current)
                        nats-2 (current)
                        nats-3 (current)
```
*Output might vary slightly. The key is to note the `Leader` and the list of `Replicas`.*

### Step 2: Publish a Message

Now, let's publish a message to the stream.

```bash
nats pub ORDERS.new "First order"
```

The NATS client library is smart enough to route this `pub` command to the correct node—the RAFT leader for the `ORDERS` stream.

### Step 3: Check Stream State

Check the stream info again. You'll see the message count has increased.

```bash
nats stream info ORDERS
```
**Expected Output:**
```
...
State:

            Messages: 1
               Bytes: 11 B
...
```
This confirms the message was received by the leader, replicated to at least one follower, and is now safely stored in the stream.

---

## 4. Consumer Replication and State

Just like streams, consumers can also be durable and replicated. When you create a consumer on a replicated stream, its state (which messages it has acknowledged) is also replicated via the same RAFT group.

### Step 1: Create a Durable, Replicated Consumer

Let's create a durable consumer called `PROCESSOR` to process the orders.

```bash
nats consumer add ORDERS PROCESSOR --durable PROCESSOR
```

This creates a consumer that can be used by multiple application instances. If one instance goes down, another can pick up where it left off.

### Step 2: Consume a Message

Now, let's consume the message we published earlier.

```bash
nats sub ORDERS.new --consumer PROCESSOR --ack
```

**Expected Output:**
```
[#1] Received on "ORDERS.new": "First order"
```

You'll see the message, and because of `--ack`, the command will acknowledge it. This acknowledgment is sent to the stream leader, written to the RAFT log, and replicated to the followers. This ensures that even if the leader fails, the new leader will know that this message was processed and won't be delivered again to this durable consumer.

---

## 5. Stream Placement and Tags

What if you have a large cluster and want to control which nodes a stream's replicas live on? This is where placement tags are useful. For example, you might have some nodes with faster SSDs and want to pin high-throughput streams to them.

### Step 1: Add Tags to Servers (Conceptual)

In a real-world scenario, you would start your NATS servers with specific tags. For our Docker setup, we would modify the `nats-server.conf` for each node.

**Example `nats-1.conf`:**
```
server_name: nats-1
jetstream: { store_dir: /data/jetstream }
cluster {
  name: nats
  # ... other cluster config
}
server_tags: ["region:west", "ssd:true"]
```
*We won't modify the Docker files now, but it's important to understand the concept.*

### Step 2: Create a Stream with Tag Placement

If we had servers with these tags, we could create a stream that will *only* be placed on nodes matching the tags.

```bash
# This command would work if our servers had the 'ssd:true' tag
nats stream add FAST_STREAM --subjects "fast.>" --replicas 3 --storage file --placement-tags "ssd:true"
```
JetStream would then find 3 nodes with the `ssd:true` tag and place the stream's RAFT group there.

---

## 6. Monitoring and Observability

The `nats` CLI is a powerful tool for monitoring.

*   **Stream Health:** `nats stream info <stream>` shows leader, replicas, and message counts.
*   **Replica Status:** Look for `(current)` next to replica names. If a replica is lagging or offline, it will be indicated here.
*   **Consumer Lag:** `nats consumer info <stream> <consumer>` shows you how many messages are unacknowledged, a key indicator of processing lag.

### Example: Checking Consumer Lag

```bash
nats consumer info ORDERS PROCESSOR
```

Look for the `Unprocessed Messages` count. A high number might indicate your consumer application is too slow or has crashed.

---

## 7. Common Operations

### Backup and Restore

JetStream streams can be easily backed up and restored.

**To Backup a Stream:**
The `nats stream backup` command captures a snapshot of the stream's configuration and all its messages.

```bash
nats stream backup ORDERS ./backup_dir
```
This will save the stream's data into the `backup_dir` directory.

**To Restore a Stream:**
The `nats stream restore` command can recreate a stream from a backup.

```bash
nats stream restore ORDERS ./backup_dir
```

This is useful for disaster recovery or migrating streams between clusters.

---

## Troubleshooting Common Issues

*   **"no leader" error:** This means the RAFT group for your stream cannot elect a leader. This usually happens if you've lost too many nodes (e.g., 2 out of 3 nodes are down). Check your cluster status (`nats server list`) and make sure enough nodes are online.
*   **Stream not placing:** If you use placement tags and the stream isn't created, it means JetStream couldn't find enough servers with the required tags to meet the replica count.
*   **Messages not being published:** Ensure you have a leader for the stream. If there's no leader, writes will be rejected.

You've now had a hands-on tour of JetStream's replication capabilities. You can see how the RAFT consensus mechanism you learned about is the foundation for providing data safety and high availability in NATS.
