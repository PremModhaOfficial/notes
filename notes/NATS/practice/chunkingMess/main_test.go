package main_test

import (
	"testing"

	main "github.com/PremModhaOfficial/nats-message-chunking"
	"github.com/nats-io/nats.go"
)

func BenchMarkFunc(b *testing.B) {
	main.ServerSetup(8, "localhost", 4223)
	// serverSetup(1, "localhost", 4223)

	nc, err := nats.Connect("http://localhost:4223")
	main.MustNotErr(err)
	defer nc.Drain()
	data := main.ReadfileToBuff("./8Mb.html")
	for range b.N {
		main.TimeSendAndRecive(nc, data)
	}
}
