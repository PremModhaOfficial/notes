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

	// MaxAckPending limits unacked messages - provides backpressure
	cons, err := js.CreateOrUpdateConsumer(context.Background(), "ORDERS", jetstream.ConsumerConfig{
		Durable:       "maxack-consumer",
		AckPolicy:     jetstream.AckExplicitPolicy,
		MaxAckPending: 5, // Only 5 unacked messages allowed

	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MaxAckPending = 5: Fetching 10 messages without acking...")

	// First fetch: should get up to 5
	msgs1, _ := cons.Fetch(10, jetstream.FetchMaxWait(2*time.Second))
	count := 0
	for msg := range msgs1.Messages() {
		count++
		log.Printf("Got message %d: %s (NOT acking)", count, string(msg.Data()))
	}
	log.Printf("Received %d messages (limited by MaxAckPending)", count)

	// Second fetch: should block/timeout - MaxAckPending reached
	log.Println("Fetching again without acking previous...")
	msgs2, _ := cons.Fetch(5, jetstream.FetchMaxWait(3*time.Second))
	count2 := 0
	for range msgs2.Messages() {
		count2++
	}
	log.Printf("Second fetch got %d messages (blocked by MaxAckPending)", count2)
}
