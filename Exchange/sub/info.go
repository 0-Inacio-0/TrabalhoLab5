package sub

import (
	"TrabalhoLab5/Exchange/config"
	"github.com/nats-io/go-nats"
	"log"
)

func printMsgInfo(m *nats.Msg) {
	log.Printf("[%s] - INFO : '%s'", m.Subject, string(m.Data))
}
func Info() {
	msgs := make(chan *nats.Msg)
	go config.Sub("info", msgs)
	log.Printf("Escutando os pedidos de Info...")
	for {
		printMsgInfo(<-msgs)
	}
}
