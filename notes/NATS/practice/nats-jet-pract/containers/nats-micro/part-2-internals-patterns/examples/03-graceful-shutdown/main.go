package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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
		Name:        "GracefulService",
		Version:     "1.0.0",
		Description: "Service with graceful shutdown",
		DoneHandler: func(srv micro.Service) {
			log.Printf("Service %s stopped, cleaning up resources...", srv.Info().Name)
		},
		Endpoint: &micro.EndpointConfig{
			Subject: "greet",
			Handler: micro.HandlerFunc(func(req micro.Request) {
				req.Respond([]byte("Hello, " + string(req.Data()) + "!"))
			}),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Service started: %s (PID: %d)", srv.Info().Name, os.Getpid())
	log.Println("Press Ctrl+C to shutdown gracefully")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Received shutdown signal")
	if err := srv.Stop(); err != nil {
		log.Printf("Error stopping service: %v", err)
	}
	log.Println("Shutdown complete")
}
