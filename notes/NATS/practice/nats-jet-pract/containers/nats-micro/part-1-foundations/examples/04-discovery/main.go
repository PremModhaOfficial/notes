package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	pingSubject, _ := micro.ControlSubject(micro.PingVerb, "", "")
	infoSubject, _ := micro.ControlSubject(micro.InfoVerb, "", "")
	statsSubject, _ := micro.ControlSubject(micro.StatsVerb, "", "")

	fmt.Println("=== Service Discovery ===")
	fmt.Printf("PING subject: %s\n", pingSubject)
	fmt.Printf("INFO subject: %s\n", infoSubject)
	fmt.Printf("STATS subject: %s\n\n", statsSubject)

	inbox := nc.NewRespInbox()
	sub, _ := nc.SubscribeSync(inbox)
	nc.PublishRequest(pingSubject, inbox, nil)

	fmt.Println("Discovered services:")
	for {
		msg, err := sub.NextMsg(500 * time.Millisecond)
		if err != nil {
			break
		}
		var ping micro.Ping
		json.Unmarshal(msg.Data, &ping)
		fmt.Printf("  - %s v%s (ID: %s)\n", ping.Name, ping.Version, ping.ID)
	}

	nc.PublishRequest(statsSubject, inbox, nil)
	fmt.Println("\nService stats:")
	for {
		msg, err := sub.NextMsg(500 * time.Millisecond)
		if err != nil {
			break
		}
		var stats micro.Stats
		json.Unmarshal(msg.Data, &stats)
		fmt.Printf("  - %s: %d endpoints\n", stats.Name, len(stats.Endpoints))
		for _, ep := range stats.Endpoints {
			fmt.Printf("      %s: %d requests\n", ep.Name, ep.NumRequests)
		}
	}
}
