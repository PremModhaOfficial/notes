package componet

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type Alerting struct{}

func (alerter *Alerting) Run(name string) error {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return err
	}
	conn.QueueSubscribe("alert.*.*", "alerting-que", func(msg *nats.Msg) {
		fmt.Printf("\n\nALERT(%s):: componet is down::\nName: %s\nStatus: %s\n", name, msg.Subject[len("health.BAD."):], string(msg.Data))
		msg.Ack()
	})

	return nil
}
