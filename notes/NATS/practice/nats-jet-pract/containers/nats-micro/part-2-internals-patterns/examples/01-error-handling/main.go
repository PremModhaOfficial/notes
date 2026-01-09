package main

import (
	"encoding/json"
	"log"
	"runtime"
	"strconv"

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
		Name:        "DivisionService",
		Version:     "1.0.0",
		Description: "Division with error handling",
		Endpoint: &micro.EndpointConfig{
			Subject: "math.divide",
			Handler: micro.HandlerFunc(handleDivide),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Service started: %s", srv.Info().Name)
	runtime.Goexit()
}

func handleDivide(req micro.Request) {
	var input struct {
		A int `json:"a"`
		B int `json:"b"`
	}
	if err := json.Unmarshal(req.Data(), &input); err != nil {
		req.Error("400", "invalid JSON input", []byte(err.Error()))
		return
	}
	if input.B == 0 {
		req.Error("422", "division by zero", nil)
		return
	}
	result := input.A / input.B
	req.Respond([]byte(strconv.Itoa(result)))
}
