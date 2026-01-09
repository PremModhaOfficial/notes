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

	// AckExplicit: Each message must be acknowledged individually
	// AckAll: Acking msg N also acks all messages before N
	// AckNone: No acks required (fire-and-forget)
	cons, err := js.CreateOrUpdateConsumer(context.Background(), "ORDERS", jetstream.ConsumerConfig{
		Durable:   "ack-explicit-consumer",
		AckPolicy: jetstream.AckExplicitPolicy,
		AckWait:   30 * time.Second, // Time before unacked message is redelivered
	})
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := cons.Fetch(3, jetstream.FetchMaxWait(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	for msg := range msgs.Messages() {
		log.Printf("Processing: %s", string(msg.Data()))
		// With AckExplicit, each message MUST be acked
		// Failure to ack = redelivery after AckWait
		if err := msg.Ack(); err != nil {
			log.Printf("Ack failed: %v", err)
		}
		log.Println("Message acknowledged")
	}
}
