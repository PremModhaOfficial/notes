// Package componet
package componet

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
)

const randomNess = 25

type Componet struct {
	function func() error
}

func (comp *Componet) Work() error {
	return comp.function()
}

func NewComponent(name string) *Componet {
	return &Componet{
		function: func() error {
			url := nats.DefaultURL
			conn, err := nats.Connect(url)
			if err != nil {
				log.Fatal(err)
			}

			var random int
			for {
				time.Sleep(1 * time.Second)
				random = rand.Int()

				if random%randomNess == 0 {
					err = conn.Publish(fmt.Sprintf("health.%s", name), []byte("BAD!!"))
					if err != nil {
						break
					}
					continue
				}
				conn.Publish(fmt.Sprintf("health.%s", name), []byte("OK"))

			}

			return err
		},
	}
}
