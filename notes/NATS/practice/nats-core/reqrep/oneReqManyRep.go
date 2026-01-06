package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

var _ error = oneReqManyRep()

func oneReqManyRep() error {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	inbox := nc.NewInbox()
	nc.Subscribe(inbox, func(msg *nats.Msg) {
		fmt.Printf(msg.Subject, string(msg.Data))
	})

	nc.PublishRequest("happy.meal", inbox, []byte("hallow"))

	defer nc.Drain()

	return nil
}
