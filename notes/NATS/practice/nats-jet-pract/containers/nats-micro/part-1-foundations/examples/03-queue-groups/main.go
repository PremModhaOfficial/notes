package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

func main() {
	instanceID := os.Getenv("INSTANCE_ID")
	if instanceID == "" {
		instanceID = "1"
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	srv, err := micro.AddService(nc, micro.Config{
		Name:        "WorkerService",
		Version:     "1.0.0",
		Description: "Load balanced worker service",
		QueueGroup:  "workers",
		Endpoint: &micro.EndpointConfig{
			Subject: "work.process",
			Handler: micro.HandlerFunc(func(req micro.Request) {
				msg := fmt.Sprintf("Processed by instance %s: %s", instanceID, string(req.Data()))
				req.Respond([]byte(msg))
			}),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Worker instance %s started (queue: workers)", instanceID)
	log.Printf("Service ID: %s", srv.Info().ID)
	runtime.Goexit()
}
