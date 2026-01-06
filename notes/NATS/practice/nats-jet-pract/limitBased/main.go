package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	nconn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nconn.Drain()

	jStream, _ := jetstream.New(nconn)
	cfg := jetstream.StreamConfig{
		Name:     "EVENTS",
		Subjects: []string{"events.>"},
	}

	cfg.Storage = jetstream.FileStorage

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, _ := jStream.CreateStream(ctx, cfg)
	fmt.Println("Stream created")

	jStream.Publish(ctx, "events.page_loaded", nil)
	jStream.Publish(ctx, "events.mouse_clicked", nil)
	jStream.Publish(ctx, "events.mouse_clicked", nil)
	jStream.Publish(ctx, "events.page_loaded", nil)
	jStream.Publish(ctx, "events.mouse_clicked", nil)
	jStream.Publish(ctx, "events.input_focused", nil)
	fmt.Println("published 6 msg")

	jStream.PublishAsync("events.page_loaded", nil)
	jStream.PublishAsync("events.mouse_clicked", nil)
	jStream.PublishAsync("events.mouse_clicked", nil)
	jStream.PublishAsync("events.page_loaded", nil)
	jStream.PublishAsync("events.mouse_clicked", nil)
	jStream.PublishAsync("events.input_focused", nil)

	select {
	case <-jStream.PublishAsyncComplete():
		fmt.Println("published 6 msg")
	case <-time.After(time.Second):
		log.Fatal("publish took too long")
	}

	printStreamState(ctx, stream)
	cfg.MaxMsgs = 10
	jStream.UpdateStream(ctx, cfg)

	fmt.Println("set max messages to 10")
	printStreamState(ctx, stream)

	cfg.MaxBytes = 300
	jStream.UpdateStream(ctx, cfg)
	fmt.Println("set max bytes to 300")

	printStreamState(ctx, stream)

	cfg.MaxAge = time.Second
	jStream.UpdateStream(ctx, cfg)
	fmt.Println("set max age to one second")

	printStreamState(ctx, stream)

	fmt.Println("sleeping one second...")
	time.Sleep(time.Second)

	printStreamState(ctx, stream)
}

func printStreamState(ctx context.Context, stream jetstream.Stream) {
	info, _ := stream.Info(ctx)
	b, _ := json.MarshalIndent(info.State, "", " ")
	fmt.Println("inspecting the state")
	fmt.Println(string(b))
}
