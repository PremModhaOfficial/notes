---
id: nats-core-messaging-patterns
aliases: []
tags: []
backlinks: []
created: 2026-01-06 15:43:19
links: []
modified: 2026-01-06 18:19:05
status: reviewed
---

> [!NOTE] Patterns
> nats give three types of comunication patterns but actualy uses PUB-SUB for all of them under the hood

## PUB-SUB
- fire and forget many to many compatible messages (without the [jetstream](notes/nats-jetstream.md))
- subscriber will recive the messages when they actualy subscribe not the previouse ones

## Req-Rep
- uses the pub-sub backend with special topics called inbox wich are one time use only subjects to listen on the reply of the req
    * if you want one req and many replies you need to do it manually this way

### single-REQ multi-REP (one to many)
```go
package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() error {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	inbox := nc.NewInbox()
	nc.Subscribe(inbox, func(msg *nats.Msg) {
		fmt.Printf(msg.Subject, string(msg.Data))
	})

	nc.PublishRequest("happy.meal", inbox, []byte("hallow"))
	return nil
}

```

## Que-Group
- This builds apon the [pub-sub]() for worker like delegation pattern (fan-out) where the workers are part of the same que-group with in the common subject
- The que-group is formed on the subjects ([wild-cards](nats-core.md#Subject-Matching) are allowed)

