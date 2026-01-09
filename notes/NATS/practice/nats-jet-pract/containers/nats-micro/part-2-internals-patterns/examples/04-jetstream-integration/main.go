package main

import (
	"context"
	"encoding/json"
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/nats-io/nats.go/micro"
)

type Order struct {
	ID       string  `json:"id"`
	Customer string  `json:"customer"`
	Amount   float64 `json:"amount"`
}

var js jetstream.JetStream

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	ctx := context.Background()
	js, err = jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"orders.>"},
	})

	srv, err := micro.AddService(nc, micro.Config{
		Name:        "OrderService",
		Version:     "1.0.0",
		Description: "Order service with JetStream events",
	})
	if err != nil {
		log.Fatal(err)
	}

	srv.AddEndpoint("order.create", micro.HandlerFunc(handleCreateOrder))

	log.Printf("Service started: %s", srv.Info().Name)
	log.Println("Requests go to micro, events published to JetStream")
	runtime.Goexit()
}

func handleCreateOrder(req micro.Request) {
	var order Order
	if err := json.Unmarshal(req.Data(), &order); err != nil {
		req.Error("400", "invalid order data", nil)
		return
	}

	order.ID = time.Now().Format("20060102150405")

	ctx := context.Background()
	eventData, _ := json.Marshal(order)
	ack, err := js.Publish(ctx, "orders.created", eventData)
	if err != nil {
		req.Error("500", "failed to publish event", nil)
		return
	}

	log.Printf("Published order event: seq=%d stream=%s", ack.Sequence, ack.Stream)
	req.RespondJSON(map[string]any{
		"status":   "created",
		"order_id": order.ID,
		"stream":   ack.Stream,
		"sequence": ack.Sequence,
	})
}
