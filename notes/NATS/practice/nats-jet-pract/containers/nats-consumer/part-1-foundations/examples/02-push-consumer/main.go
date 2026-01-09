package main

import (
	"context"
	"log"
	"os"
	"os/signal"
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

	// Create push consumer with deliver subject and heartbeat
	cons, err := js.CreateOrUpdateConsumer(context.Background(), "ORDERS", jetstream.ConsumerConfig{
		Durable:        "push-consumer",
		DeliverSubject: "orders.push.deliver",
		AckPolicy:      jetstream.AckExplicitPolicy,
		IdleHeartbeat:  5 * time.Second, // Server sends heartbeat if no messages
	})
	if err != nil {
		log.Fatal(err)
	}

	// Consume with callback
	consCtx, err := cons.Consume(func(msg jetstream.Msg) {
		log.Printf("Received: %s", string(msg.Data()))
		msg.Ack()
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consCtx.Stop()

	log.Println("Listening for messages (Ctrl+C to exit)...")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
}
