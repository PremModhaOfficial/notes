package main

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	// Exponential backoff: 1s, 5s, 30s between retries
	// After MaxDeliver attempts, message goes to dead letter (if configured)
	cons, err := js.CreateOrUpdateConsumer(context.Background(), "ORDERS", jetstream.ConsumerConfig{
		Durable:    "backoff-consumer",
		AckPolicy:  jetstream.AckExplicitPolicy,
		MaxDeliver: 4, // Max 4 delivery attempts
		BackOff: []time.Duration{
			1 * time.Second,  // 1st retry after 1s
			5 * time.Second,  // 2nd retry after 5s
			30 * time.Second, // 3rd retry after 30s (and subsequent)
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Consumer created with backoff: [1s, 5s, 30s], MaxDeliver: 4")

	// Simulate processing with failures
	cons.Consume(func(msg jetstream.Msg) {
		meta, _ := msg.Metadata()
		log.Printf("Attempt %d/%d: %s", meta.NumDelivered, 4, string(msg.Data()))

		// Simulate failure - Nak triggers backoff
		msg.Nak()
	})

	time.Sleep(45 * time.Second) // Watch retry behavior
}
