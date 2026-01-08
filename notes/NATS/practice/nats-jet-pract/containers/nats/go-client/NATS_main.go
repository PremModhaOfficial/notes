package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

const (
	// Stream and Consumer Configuration
	StreamName    = "LEARNING_STREAM"
	ConsumerName  = "LEARNING_CONSUMER"
	SubjectPrefix = "learning"

	// Message Publishing Configuration
	PublishInterval = 2 * time.Second

	// Connection Configuration
	ConnectTimeout = 30 * time.Second
	ContextTimeout = 10 * time.Second
)

// NATSClient represents our NATS JetStream client with all necessary components
type NATSClient struct {
	conn       *nats.Conn          // Core NATS connection
	js         jetstream.JetStream // JetStream context for advanced messaging
	stream     jetstream.Stream    // JetStream stream for message persistence
	consumer   jetstream.Consumer  // JetStream consumer for message processing
	ctx        context.Context     // Context for cancellation
	cancel     context.CancelFunc  // Cancel function for graceful shutdown
	wg         sync.WaitGroup      // WaitGroup to coordinate goroutines
	msgCounter int                 // Counter for published messages
	mu         sync.RWMutex        // Mutex to protect shared state
}

func main() {
	fmt.Println("🚀 Starting NATS RAFT Learning Client")
	fmt.Println("=====================================")

	// Initialize the NATS client
	client, err := NewNATSClient()
	if err != nil {
		log.Fatalf("❌ Failed to create NATS client: %v", err)
	}

	// Setup graceful shutdown handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the client operations
	client.Start()

	// Wait for shutdown signal
	<-sigChan
	fmt.Println("\n🛑 Shutdown signal received, cleaning up...")

	// Perform graceful shutdown
	client.Shutdown()
	fmt.Println("✅ Client shutdown complete")
}

// NewNATSClient creates and initializes a new NATS client with cluster discovery
func NewNATSClient() (*NATSClient, error) {
	// Create context with timeout for initialization
	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()

	// Configure connection options for cluster discovery and resilience
	opts := []nats.Option{
		nats.Name("NATS-RAFT-Learning-Client"),                    // Client identifier
		nats.MaxReconnects(-1),                                    // Unlimited reconnection attempts
		nats.ReconnectWait(2 * time.Second),                       // Wait 2 seconds between reconnects
		nats.ReconnectJitter(500*time.Millisecond, 2*time.Second), // Add jitter to prevent thundering herd
		nats.Timeout(10 * time.Second),                            // Connection timeout

		// Connection event handlers for monitoring
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			if err != nil {
				fmt.Printf("⚠️  Disconnected from NATS: %v\n", err)
			} else {
				fmt.Println("🔌 Disconnected from NATS")
			}
		}),

		nats.ReconnectHandler(func(nc *nats.Conn) {
			fmt.Printf("🔄 Reconnected to NATS server: %s\n", nc.ConnectedUrl())
			fmt.Printf("📊 Connection stats - Reconnects: %d\n", nc.Stats().Reconnects)
		}),

		nats.ClosedHandler(func(nc *nats.Conn) {
			fmt.Println("🔴 NATS connection closed")
			if lastErr := nc.LastError(); lastErr != nil {
				fmt.Printf("❌ Last error: %v\n", lastErr)
			}
		}),
	}

	// Define cluster endpoints (supports both single node and cluster)
	// The client will automatically discover and connect to available nodes
	servers := []string{
		"nats://localhost:4222", // Primary node or single node
		"nats://localhost:4223", // Secondary node (cluster mode)
		"nats://localhost:4224", // Tertiary node (cluster mode)
	}

	fmt.Printf("🔍 Attempting to connect to NATS cluster: %v\n", servers)

	// Establish connection to NATS cluster
	conn, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		// Fallback: try connecting to individual cluster nodes
		for i, server := range servers {
			fmt.Printf("🔄 Trying server %d: %s\n", i+1, server)
			conn, err = nats.Connect(server, opts...)
			if err == nil {
				break
			}
			fmt.Printf("⚠️  Failed to connect to %s: %v\n", server, err)
		}

		if err != nil {
			return nil, fmt.Errorf("failed to connect to any NATS server: %w", err)
		}
	}

	// Display successful connection information
	fmt.Printf("✅ Connected to NATS server: %s\n", conn.ConnectedUrl())
	fmt.Printf("📋 Server ID: %s\n", conn.ConnectedServerId())
	fmt.Printf("🏷️  Server Name: %s\n", conn.ConnectedServerName())

	// Create JetStream context for advanced messaging features
	js, err := jetstream.New(conn)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to create JetStream context: %w", err)
	}

	fmt.Println("🎯 JetStream context created successfully")

	// Create application context for coordinated shutdown
	appCtx, appCancel := context.WithCancel(context.Background())

	client := &NATSClient{
		conn:   conn,
		js:     js,
		ctx:    appCtx,
		cancel: appCancel,
	}

	// Initialize JetStream resources (stream and consumer)
	if err := client.initializeJetStream(ctx); err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to initialize JetStream: %w", err)
	}

	return client, nil
}

// initializeJetStream sets up the JetStream stream and consumer
func (c *NATSClient) initializeJetStream(ctx context.Context) error {
	fmt.Println("🏗️  Initializing JetStream resources...")

	// Configure the stream for message persistence and replication
	streamConfig := jetstream.StreamConfig{
		Name:        StreamName,
		Description: "NATS RAFT Learning Stream - demonstrates JetStream persistence and clustering",
		Subjects:    []string{fmt.Sprintf("%s.*", SubjectPrefix)}, // Match all subjects starting with "learning."

		// Retention and storage policies
		Retention:    jetstream.WorkQueuePolicy, // Messages are removed after acknowledgment
		MaxConsumers: 10,                        // Allow up to 10 consumers
		MaxMsgs:      1000,                      // Store maximum 1000 messages
		MaxBytes:     1024 * 1024,               // Maximum 1MB storage
		MaxAge:       24 * time.Hour,            // Keep messages for 24 hours max

		// Replication settings (important for RAFT consensus)
		Replicas: 1, // Start with 1 replica, can be increased for HA

		// Storage type
		Storage: jetstream.FileStorage, // Persistent file-based storage
	}

	// Create or update the stream
	stream, err := c.js.CreateOrUpdateStream(ctx, streamConfig)
	if err != nil {
		return fmt.Errorf("failed to create stream: %w", err)
	}
	c.stream = stream

	fmt.Printf("✅ Stream '%s' created/updated successfully\n", StreamName)

	// Display stream information
	info, err := stream.Info(ctx)
	if err != nil {
		fmt.Printf("⚠️  Could not get stream info: %v\n", err)
	} else {
		fmt.Printf("📊 Stream info - Messages: %d, Bytes: %d, Consumers: %d\n",
			info.State.Msgs, info.State.Bytes, info.State.Consumers)
	}

	// Configure the consumer for message processing
	consumerConfig := jetstream.ConsumerConfig{
		Name:          ConsumerName,
		Description:   "NATS RAFT Learning Consumer - processes messages with acknowledgment",
		Durable:       ConsumerName,                  // Make it durable to survive restarts
		AckPolicy:     jetstream.AckExplicitPolicy,   // Require explicit acknowledgment
		AckWait:       30 * time.Second,              // Wait 30 seconds for ack before redelivery
		MaxDeliver:    3,                             // Maximum 3 delivery attempts
		DeliverPolicy: jetstream.DeliverAllPolicy,    // Process all messages in stream
		ReplayPolicy:  jetstream.ReplayInstantPolicy, // Process messages as fast as possible

		// Flow control to prevent overwhelming
		MaxAckPending: 100, // Allow up to 100 unacknowledged messages
	}

	// Create or update the consumer
	consumer, err := stream.CreateOrUpdateConsumer(ctx, consumerConfig)
	if err != nil {
		return fmt.Errorf("failed to create consumer: %w", err)
	}
	c.consumer = consumer

	fmt.Printf("✅ Consumer '%s' created/updated successfully\n", ConsumerName)

	return nil
}

// Start begins the client operations (publisher and consumer)
func (c *NATSClient) Start() {
	fmt.Println("🎬 Starting client operations...")

	// Start the message publisher in a separate goroutine
	c.wg.Add(1)
	go c.runPublisher()

	// Start the message consumer in a separate goroutine
	c.wg.Add(1)
	go c.runConsumer()

	fmt.Println("✅ Client started - publisher and consumer are running")
}

// runPublisher publishes messages at regular intervals
func (c *NATSClient) runPublisher() {
	defer c.wg.Done()

	fmt.Println("📤 Starting message publisher...")
	ticker := time.NewTicker(PublishInterval)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			fmt.Println("📤 Publisher shutting down...")
			return
		case <-ticker.C:
			c.publishMessage()
		}
	}
}

// publishMessage publishes a single message with current timestamp and counter
func (c *NATSClient) publishMessage() {
	// Create context with timeout for this publish operation
	ctx, cancel := context.WithTimeout(c.ctx, ContextTimeout)
	defer cancel()

	// Increment message counter safely
	c.mu.Lock()
	c.msgCounter++
	msgID := c.msgCounter
	c.mu.Unlock()

	// Create message payload with useful information
	timestamp := time.Now()
	payload := fmt.Sprintf(`{
		"id": %d,
		"timestamp": "%s",
		"message": "Hello from NATS RAFT Learning Client",
		"node": "%s",
		"counter": %d
	}`, msgID, timestamp.Format(time.RFC3339), c.conn.ConnectedUrl(), msgID)

	// Publish message to JetStream
	subject := fmt.Sprintf("%s.messages", SubjectPrefix)
	ack, err := c.js.Publish(ctx, subject, []byte(payload))
	if err != nil {
		fmt.Printf("❌ Failed to publish message %d: %v\n", msgID, err)
		return
	}

	// Display publish confirmation
	fmt.Printf("📤 Published message %d (seq: %d, stream: %s) to %s\n",
		msgID, ack.Sequence, ack.Stream, subject)
}

// runConsumer processes incoming messages continuously
func (c *NATSClient) runConsumer() {
	defer c.wg.Done()

	fmt.Println("📥 Starting message consumer...")

	// Create message iterator for continuous processing
	iter, err := c.consumer.Messages(
		jetstream.PullMaxMessages(10),       // Fetch up to 10 messages at a time
		jetstream.PullExpiry(5*time.Second), // Wait up to 5 seconds for messages
	)
	if err != nil {
		fmt.Printf("❌ Failed to create message iterator: %v\n", err)
		return
	}
	defer iter.Stop()

	fmt.Println("📥 Consumer ready - waiting for messages...")

	// Process messages continuously
	for {
		select {
		case <-c.ctx.Done():
			fmt.Println("📥 Consumer shutting down...")
			return
		default:
			// Try to get next message with timeout
			msg, err := iter.Next()
			if err != nil {
				// Check if this is a timeout (expected) or real error
				if err == context.DeadlineExceeded {
					continue // Normal timeout, keep trying
				}
				fmt.Printf("⚠️  Error receiving message: %v\n", err)
				continue
			}

			c.processMessage(msg)
		}
	}
}

// processMessage handles an individual message
func (c *NATSClient) processMessage(msg jetstream.Msg) {
	// Extract message metadata
	meta, err := msg.Metadata()
	if err != nil {
		fmt.Printf("⚠️  Could not get message metadata: %v\n", err)
	}

	// Display message information
	fmt.Printf("📥 Received message (seq: %d, delivered: %d times)\n",
		meta.Sequence.Stream, meta.NumDelivered)
	fmt.Printf("   Subject: %s\n", msg.Subject())
	fmt.Printf("   Data: %s\n", string(msg.Data()))
	fmt.Printf("   Timestamp: %s\n", meta.Timestamp.Format(time.RFC3339))

	// Simulate some processing time
	time.Sleep(100 * time.Millisecond)

	// Acknowledge the message (important for JetStream)
	if err := msg.Ack(); err != nil {
		fmt.Printf("❌ Failed to acknowledge message: %v\n", err)
	} else {
		fmt.Printf("✅ Message acknowledged successfully\n")
	}

	fmt.Println("---")
}

// Shutdown performs graceful shutdown of the client
func (c *NATSClient) Shutdown() {
	fmt.Println("🛑 Initiating graceful shutdown...")

	// Cancel context to signal all goroutines to stop
	c.cancel()

	// Wait for all goroutines to finish (with timeout)
	done := make(chan struct{})
	go func() {
		c.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("✅ All goroutines stopped gracefully")
	case <-time.After(10 * time.Second):
		fmt.Println("⚠️  Timeout waiting for goroutines to stop")
	}

	// Close NATS connection
	c.Close()
}

// Close closes the NATS connection
func (c *NATSClient) Close() {
	if c.conn != nil && !c.conn.IsClosed() {
		fmt.Println("🔐 Closing NATS connection...")
		c.conn.Close()
		fmt.Println("✅ NATS connection closed")
	}
}

// DisplayConnectionInfo shows current connection status and statistics
func (c *NATSClient) DisplayConnectionInfo() {
	if c.conn == nil {
		fmt.Println("❌ No connection available")
		return
	}

	fmt.Println("\n📊 Connection Information:")
	fmt.Println("==========================")
	fmt.Printf("Connected: %t\n", c.conn.IsConnected())
	fmt.Printf("Server URL: %s\n", c.conn.ConnectedUrl())
	fmt.Printf("Server ID: %s\n", c.conn.ConnectedServerId())
	fmt.Printf("Server Name: %s\n", c.conn.ConnectedServerName())

	stats := c.conn.Stats()
	fmt.Printf("Messages In: %d\n", stats.InMsgs)
	fmt.Printf("Messages Out: %d\n", stats.OutMsgs)
	fmt.Printf("Bytes In: %d\n", stats.InBytes)
	fmt.Printf("Bytes Out: %d\n", stats.OutBytes)
	fmt.Printf("Reconnects: %d\n", stats.Reconnects)

	c.mu.RLock()
	fmt.Printf("Published Messages: %d\n", c.msgCounter)
	c.mu.RUnlock()

	fmt.Println("==========================\n")
}
