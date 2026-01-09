package main

import (
	"context"
	"log"

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

	cons, err := js.Consumer(context.Background(), "ORDERS", "pull-consumer")
	if err != nil {
		log.Fatal(err)
	}

	info, err := cons.Info(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("=== Consumer Debug Info ===")
	log.Printf("Name:           %s", info.Name)
	log.Printf("Stream:         %s", info.Stream)
	log.Printf("NumPending:     %d (messages waiting to be delivered)", info.NumPending)
	log.Printf("NumAckPending:  %d (delivered but not yet acked)", info.NumAckPending)
	log.Printf("NumRedelivered: %d (messages redelivered due to nak/timeout)", info.NumRedelivered)
	log.Printf("NumWaiting:     %d (pull requests waiting for messages)", info.NumWaiting)
	log.Printf("Delivered:      Stream=%d, Consumer=%d", info.Delivered.Stream, info.Delivered.Consumer)
	log.Printf("AckFloor:       Stream=%d, Consumer=%d", info.AckFloor.Stream, info.AckFloor.Consumer)

	// Debugging tips
	if info.NumAckPending > 0 {
		log.Println("\nWARNING: Messages delivered but not acked - check processing!")
	}
	if info.NumRedelivered > 0 {
		log.Println("\nWARNING: Redeliveries detected - check for failures or slow consumers!")
	}
}
