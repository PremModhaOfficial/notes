package main

import (
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	srv, err := micro.AddService(nc, micro.Config{
		Name:        "EchoService",
		Version:     "1.0.0",
		Description: "Simple echo service",
		Endpoint: &micro.EndpointConfig{
			Subject: "echo",
			Handler: micro.HandlerFunc(func(req micro.Request) {
				req.Respond(req.Data())
			}),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Service started: %s (%s)", srv.Info().Name, srv.Info().ID)
	runtime.Goexit()
}
