package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

const sentMessages = 1000

// func main() {
// 	ServerSetup(8, "localhost", 4223)
// 	// serverSetup(1, "localhost", 4223)
//
// 	nc, err := nats.Connect("http://localhost:4223")
// 	MustNotErr(err)
// 	defer nc.Drain()
//
// 	TimeSendAndRecive(nc, ReadfileToBuff("./8Mb.html"))
// }

func TimeSendAndRecive(nc *nats.Conn, data []byte) {
	fmt.Printf("data size: %d\n", len(data))
	MsgChan := make(chan struct{}, sentMessages)
	nc.Subscribe("msg.prem", func(msg *nats.Msg) {
		fmt.Printf("data got: %d \n", len(msg.Data))
		MsgChan <- struct{}{}
		fmt.Println("msg got")
	})

	start := time.Now()
	fmt.Printf("start publishing: %s\n", start)
	for range sentMessages {
		nc.Publish("msg.prem", data)
	}

	recivedMsgs := 0

	for recivedMsgs < sentMessages {
		<-MsgChan
		fmt.Println("msg chunk recived")
		recivedMsgs++
	}

	diff := time.Until(start).Abs().Seconds()
	fmt.Printf("diff in time %f seconds \n(Messages sent:%d)", diff, sentMessages)
}

func ReadfileToBuff(path string) []byte {
	var data []byte

	data, err := os.ReadFile(path)
	MustNotErr(err)

	return data
}

func ServerSetup(msgMaxPayload int, host string, port int) {
	natsServer, err := server.NewServer(&server.Options{
		ServerName:   strconv.Itoa(msgMaxPayload) + "m-server",
		Host:         host,
		Port:         port,
		Trace:        true,
		Debug:        true,
		TraceVerbose: false,
		TraceHeaders: false,
		MaxPayload:   int32(msgMaxPayload) * 1024 * 1024,
	})
	MustNotErr(err)

	natsServer.Start()
}

func MustNotErr(err error) {
	if err != nil {
		panic(err)
	}
}
