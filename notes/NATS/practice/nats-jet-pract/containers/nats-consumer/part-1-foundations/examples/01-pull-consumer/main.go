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

	// Get or create consumer on ORDERS stream
	cons, err := js.CreateOrUpdateConsumer(ctx(), "ORDERS", jetstream.ConsumerConfig{
		Durable:   "pull-consumer",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Pull 5 messages
	msgs, err := cons.Fetch(5, jetstream.FetchMaxWait(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	for msg := range msgs.Messages() {
		log.Printf("Received: %s", string(msg.Data()))
		msg.Ack()
	}

	if msgs.Error() != nil {
		log.Printf("Fetch error: %v", msgs.Error())
	}
}

func ctx() context.Context { return context.Background() }
