package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

const sentMessages = 1000
const (
	MultMsgID   string = "MUL-ID"
	MultMsgSq   string = "MUL-SQ"
	MultMsgTotl string = "MUL-MX"
)

func main() {
	// ServerSetup(8, "localhost", 4223)
	ServerSetup(1, "localhost", 4223)

	nc, err := nats.Connect("http://localhost:4223")
	MustNotErr(err)
	defer nc.Drain()

	TimeSendAndReciveChunk(nc, ReadfileToBuff("./8Mb.html"))
}

func TimeSendAndReciveChunk(nc *nats.Conn, data []byte) {
	fmt.Printf("data size: %d\n", len(data))
	MsgChan := make(chan struct{}, sentMessages)

	// nc.Subscribe("msg.prem", func(msg *nats.Msg) {
	// 	fmt.Printf("data got: %d \n", len(msg.Data))
	// 	MsgChan <- struct{}{}
	// })

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		SubscribeChunked(nc, "msg.prem", 1*1024*1024, func(msg *nats.Msg) {
			fmt.Printf("data got: %d \n", len(msg.Data))
			MsgChan <- struct{}{}
		})
	}()

	start := time.Now()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("start publishing: %s\n", start)
		for range sentMessages {
			PublishChunked(nc, "msg.prem", data, 1*1024*1024)
			// nc.Publish("msg.prem", data)
		}
	}()

	recivedMsgs := 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		for recivedMsgs < sentMessages {
			<-MsgChan
			recivedMsgs++
		}

		diff := time.Until(start).Abs().Seconds()
		fmt.Printf("diff in time %f seconds \n(Messages sent:%d)", diff, sentMessages)
	}()

	wg.Wait()
}

func PublishChunked(natsConn *nats.Conn, subject string, data []byte, chunkSize int) {
	id := uuid.New()
	dataLength := len(data)
	totalChuks := dataLength / int(chunkSize)

	for i := 0; i < totalChuks-1; i++ {
		msg := nats.Msg{
			Subject: subject,
		}
		msg.Data = data[i*chunkSize : (i+1)*chunkSize]

		msg.Header = nats.Header{}
		msg.Header.Set(MultMsgID, id.String())
		msg.Header.Set(MultMsgSq, strconv.Itoa(i))
		msg.Header.Set(MultMsgTotl, strconv.Itoa(totalChuks))

		natsConn.PublishMsg(&msg)

	}
}

func SubscribeChunked(natsConn *nats.Conn, subject string, chunkSize int, cb nats.MsgHandler) {
	// TODO: what data type should the sequance be?  int or string??
	// string for now
	type chunkedMsg map[string]*[]byte
	cache := make(map[string]chunkedMsg)
	natsConn.Subscribe(subject, func(msg *nats.Msg) {
		msgID := msg.Header.Get(MultMsgID)
		msgSeq := msg.Header.Get(MultMsgSq)
		total, err := strconv.ParseInt(msg.Header.Get(MultMsgTotl), 10, 32)
		MustNotErr(err)
		if msgPart, ok := cache[msgID]; ok {
			msgPart[msgSeq] = &msg.Data
			data := []byte{}

			if len(*msgPart[msgSeq]) >= int(total) {
				for i := range total {
					data = append(data, ((*msgPart[msgSeq])[i]))
				}
				wholeMsg := new(nats.Msg)
				wholeMsg.Sub = msg.Sub
				wholeMsg.Reply = msg.Reply
				wholeMsg.Data = data

				cb(wholeMsg)
			}
		}
	})
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
