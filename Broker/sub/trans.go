package sub

import (
	"TrabalhoLab5/Broker/config"
	"github.com/nats-io/go-nats"
	"log"
)

func printMsg(m *nats.Msg) {
	log.Printf("[%s] - Transação : '%s'", m.Subject, string(m.Data))
}
func Trans() {
	msgs := make(chan *nats.Msg)
	go config.Sub("transacao.*", msgs)
	log.Printf("Escutando as transações ...")
	printMsg(<-msgs)
}
