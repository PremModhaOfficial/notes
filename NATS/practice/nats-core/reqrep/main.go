package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	conn, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Drain()

	sub, _ := conn.Subscribe("ask.*", func(msg *nats.Msg) {
		subTopic := msg.Subject[4:]
		msg.Respond([]byte("hello " + subTopic))
	})

	rep, _ := conn.Request("ask.prem", nil, time.Second)
	fmt.Println(string(rep.Data))

	rep, _ = conn.Request("ask.modha", nil, time.Second)
	fmt.Println(string(rep.Data))

	rep, _ = conn.Request("ask.anyone", nil, time.Second)
	fmt.Println(string(rep.Data))

	sub.Unsubscribe()

	_, err = conn.Request("greet.joe", nil, time.Second)
	fmt.Println(err)
}
