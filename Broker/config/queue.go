package config

import (
	"github.com/nats-io/go-nats"
	"log"
	"time"
)

type Queue struct {
	Server string
	Subj   string
}

var broker = Queue{
	Server: "nats://localhost:4222",
	Subj:   "Broker",
}

var exchange = Queue{
	Server: "nats://localhost:4222",
	Subj:   "exchange",
}

func Pub(subj string, msg []byte) {
	// Connect Options.
	opts := []nats.Option{nats.Name("Transaction Publisher")}

	// Connect to NATS
	nc, err := nats.Connect(broker.Server, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.Publish((broker.Subj + "." + subj), msg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Enviado [%s] : '%s'\n", broker.Subj+"."+subj, string(msg))
	}
}

func Sub(subj string, msgChan chan *nats.Msg) {
	// Connect Options.
	opts := []nats.Option{nats.Name("Exchange Sub")}
	opts = setupConnOptions(opts)

	// Connect to NATS
	nc, err := nats.Connect(exchange.Server, opts...)
	if err != nil {
		log.Fatal(err)
	}

	_, _ = nc.Subscribe(exchange.Subj+"."+subj, func(msg *nats.Msg) {
		msgChan <- msg
	})
	_ = nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}
}
func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectHandler(func(nc *nats.Conn) {
		log.Printf("Disconnected: will attempt reconnects for %.0fm", totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatal("Exiting, no servers available")
	}))
	return opts
}
