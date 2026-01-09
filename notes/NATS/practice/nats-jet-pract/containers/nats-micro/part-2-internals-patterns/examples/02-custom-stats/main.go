package main

import (
	"encoding/json"
	"log"
	"runtime"
	"strings"
	"sync/atomic"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

var (
	echoCount  atomic.Int64
	upperCount atomic.Int64
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	srv, err := micro.AddService(nc, micro.Config{
		Name:        "TextService",
		Version:     "1.0.0",
		Description: "Text processing with custom stats",
		StatsHandler: func(e *micro.Endpoint) any {
			return map[string]any{
				"echo_requests":      echoCount.Load(),
				"uppercase_requests": upperCount.Load(),
			}
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	srv.AddEndpoint("echo", micro.HandlerFunc(func(req micro.Request) {
		echoCount.Add(1)
		req.Respond(req.Data())
	}))

	srv.AddEndpoint("upper", micro.HandlerFunc(func(req micro.Request) {
		upperCount.Add(1)
		data := string(req.Data())
		req.Respond([]byte(strings.ToUpper(data)))
	}))

	log.Printf("Service started: %s (query $SRV.STATS for custom metrics)", srv.Info().Name)

	stats := srv.Stats()
	statsJSON, _ := json.MarshalIndent(stats, "", "  ")
	log.Printf("Stats: %s", statsJSON)

	runtime.Goexit()
}
