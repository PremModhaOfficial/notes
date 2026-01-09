package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

func main() {
	var (
		host           string
		port           int
		certFile       string
		keyFile        string
		clientCertFile string
		clientKeyFile  string
		caFile         string
	)

	flag.StringVar(&host, "host", "localhost", "Client connection host/IP.")
	flag.IntVar(&port, "port", 4222, "Client connection port.")
	flag.StringVar(&certFile, "tls.cert.server", "cert.pem", "TLS cert file.")
	flag.StringVar(&keyFile, "tls.key.server", "key.pem", "TLS key file.")
	flag.StringVar(&caFile, "tls.ca", "ca.pem", "TLS CA file.")
	flag.StringVar(&clientCertFile, "tls.cert.client", "client-cert.pem", "TLS cert file.")
	flag.StringVar(&clientKeyFile, "tls.key.client", "client-key.pem", "TLS key file.")

	flag.Parse()

	serverConf, err := server.GenTLSConfig(&server.TLSConfigOpts{
		CertFile: certFile,
		KeyFile:  keyFile,
		CaFile:   caFile,
		Verify:   true,
		Timeout:  2,
	})
	if err != nil {
		log.Fatal(err)
	}

	opts := server.Options{
		Host:      host,
		Port:      port,
		TLSConfig: serverConf,
		TLS:       true,
	}

	ns, err := server.NewServer(&opts)
	if err != nil {
		log.Fatal(err)
	}
	ns.ConfigureLogger()

	go ns.Start()
	defer ns.Shutdown()

	time.Sleep(3 * time.Second)

	nc, err := nats.Connect(
		fmt.Sprintf("tls://%s:%d", host, port),
		nats.RootCAs(caFile),
		nats.ClientCert(clientCertFile, clientKeyFile),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	sub, _ := nc.SubscribeSync("PREM")
	defer sub.Drain()

	nc.Publish("PREM", []byte("DATA"))

	msg, _ := sub.NextMsg(1 * time.Second)
	fmt.Println("msg: ", string(msg.Data))
}
