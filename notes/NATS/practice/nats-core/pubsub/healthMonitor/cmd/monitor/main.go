package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/PremModhaOfficial/reactive/healthMonitor/cmd/componet"
	"github.com/nats-io/nats.go"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	url := os.Getenv("NATS_URL")

	if url == "" {
		url = nats.DefaultURL
	}

	conn, err := nats.Connect(url)
	if err != nil {
		return err
	}

	_, err = conn.Subscribe("health.>", func(msg *nats.Msg) {
		msgData := string(msg.Data)
		if msgData != "OK" {
			fmt.Println(msg.Subject)

			// id, err := strconv.ParseInt(strings.Split(msg.Subject, ".")[2], 10, 10)

			// if id == 0 || err != nil {
			// 	return
			// }
			fmt.Printf("alert.%s.%s", strings.Split(msg.Subject, ".")[3], strings.Split(msg.Subject, ".")[2])
			err = conn.Publish(fmt.Sprintf("alert.%s.%s", strings.Split(msg.Subject, ".")[3], strings.Split(msg.Subject, ".")[2]), msg.Data)
			if err != nil {
				return
			}
		}
	})
	if err != nil {
		return err
	}

	names := []string{"sche", "poll", "disc"}
	wg := sync.WaitGroup{}
	for i := range 10 {
		wg.Go(func() {
			err := componet.NewComponent("componet." + names[i%len(names)] + "." + strconv.FormatInt(int64(i), 10)).Work()
			if err != nil {
				log.Fatal(err)
			}
		})
	}

	for i := range 3 {
		wg.Go(func() {
			al := &componet.Alerting{}
			al.Run(fmt.Sprintf("%d", i))
		})
	}
	wg.Wait()

	return nil
}
