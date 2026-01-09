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

	cons, err := js.CreateOrUpdateConsumer(context.Background(), "ORDERS", jetstream.ConsumerConfig{
		Durable:   "ack-types-consumer",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := cons.Fetch(5, jetstream.FetchMaxWait(10*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	for msg := range msgs.Messages() {
		log.Printf("Message %d: %s", i, string(msg.Data()))

		switch i {
		case 0:
			msg.Ack() // Success - message processed
			log.Println("  -> Ack: Message processed successfully")
		case 1:
			msg.Nak() // Failure - redeliver immediately
			log.Println("  -> Nak: Request immediate redelivery")
		case 2:
			msg.NakWithDelay(5 * time.Second) // Failure - redeliver after delay
			log.Println("  -> NakWithDelay: Redeliver in 5 seconds")
		case 3:
			msg.InProgress() // Still working - reset ack wait timer
			log.Println("  -> InProgress: Reset ack timer, still processing")
			time.Sleep(2 * time.Second)
			msg.Ack()
		case 4:
			msg.Term() // Terminal failure - never redeliver
			log.Println("  -> Term: Message is poison, never redeliver")
		}
		i++
	}
}
