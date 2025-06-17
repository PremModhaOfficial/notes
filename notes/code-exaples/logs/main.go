package main

import (
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/andrewhowdencom/courses.pito/delivery-service/carriers"
	"github.com/andrewhowdencom/courses.pito/delivery-service/server"
)

var addr = flag.String("a", "lcalhost:9093", "server address")
var log *slog.Logger

func main() {
	log.Info("Application Started")
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGINT)

	srv := server.New(&carriers.Carriers{})

	go func() {
		if err := srv.Listen(*addr); err != nil {

			log.Error("server cannot listen", "error", err, "address", addr)
			os.Exit(1)
		}
	}()

	log.Info("awiting shutdown signal")
	<-ch
	log.Info("shutdown signal recieved")

	if err := srv.Shutdown(); err != nil {
		log.Error("server shutdown failed", "error", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func init() {
	flag.Parse()

	log = slog.New(slog.NewJSONHandler(os.Stderr, nil))
}
