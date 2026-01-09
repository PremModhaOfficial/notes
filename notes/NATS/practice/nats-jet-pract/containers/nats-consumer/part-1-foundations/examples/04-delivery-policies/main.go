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

	ctx := context.Background()

	// DeliverLast: Start from last message in stream
	_, err = js.CreateOrUpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
		Durable:       "deliver-last-consumer",
		DeliverPolicy: jetstream.DeliverLastPolicy,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created consumer with DeliverLast policy")

	// DeliverByStartSequence: Start from specific sequence number
	_, err = js.CreateOrUpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
		Durable:       "deliver-seq-consumer",
		DeliverPolicy: jetstream.DeliverByStartSequencePolicy,
		OptStartSeq:   100, // Start from sequence 100
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created consumer starting at sequence 100")

	// DeliverByStartTime: Start from specific time
	_, err = js.CreateOrUpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
		Durable:       "deliver-time-consumer",
		DeliverPolicy: jetstream.DeliverByStartTimePolicy,
		OptStartTime:  ptr(time.Now().Add(-1 * time.Hour)), // Last hour
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created consumer starting from 1 hour ago")
}

func ptr(t time.Time) *time.Time { return &t }
