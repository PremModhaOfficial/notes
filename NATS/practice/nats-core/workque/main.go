package main

import (
	"fmt"
	"log"
	"os"
	"sync"
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
	workerConut := 20
	subMu := sync.Mutex{}
	var subs []*nats.Subscription

	wg := sync.WaitGroup{}
	for i := workerConut; i > 0; i-- {
		workerID := i
		wg.Go(func() {
			sub, err := conn.QueueSubscribe("ask.*", "answers", func(msg *nats.Msg) {
				subTopic := msg.Subject[4:]
				msg.Respond([]byte(fmt.Sprintf("Worker-%d connected with ", workerID) + subTopic))
			})
			subMu.Lock()
			subs = append(subs, sub)
			subMu.Unlock()
			if err != nil {
				log.Printf("Error subscribing worker %d: %v", workerID, err)
				return
			}
			// defer sub.Unsubscribe()
		})
	}

	size := 1000
	replies := make(chan string, size)

	wg.Go(func() {
		for resp := range replies {
			fmt.Println(resp)
		}
	})

	for i := range size {
		rep, err := conn.Request(fmt.Sprintf("ask.%d", i), []byte("hi"), time.Second)
		if err != nil {
			log.Printf("Error requesting ask.%d: %v", i, err)
			continue
		}
		if rep != nil && rep.Data != nil {
			replies <- string(rep.Data)
		}
	}
	close(replies)
	wg.Wait()
	for _, sub := range subs {
		err := sub.Drain()
		if err != nil {
			log.Fatal(err)
		}
		err = sub.Unsubscribe()
		if err != nil {
			log.Fatal(err)
		}
	}
	conn.Drain()
}
