package main

import (
	"encoding/json"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

type MathRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type MathResponse struct {
	Result int `json:"result"`
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	srv, err := micro.AddService(nc, micro.Config{
		Name:        "MathService",
		Version:     "1.0.0",
		Description: "Math operations service",
	})
	if err != nil {
		log.Fatal(err)
	}

	mathGroup := srv.AddGroup("math")
	mathGroup.AddEndpoint("add", micro.HandlerFunc(handleAdd))
	mathGroup.AddEndpoint("multiply", micro.HandlerFunc(handleMultiply))

	log.Printf("Service started: %s - endpoints: math.add, math.multiply", srv.Info().Name)
	runtime.Goexit()
}

func handleAdd(req micro.Request) {
	var r MathRequest
	json.Unmarshal(req.Data(), &r)
	req.RespondJSON(MathResponse{Result: r.A + r.B})
}

func handleMultiply(req micro.Request) {
	var r MathRequest
	json.Unmarshal(req.Data(), &r)
	req.RespondJSON(MathResponse{Result: r.A * r.B})
}
