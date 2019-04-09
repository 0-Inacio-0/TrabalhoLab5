package sub

import (
	"TrabalhoLab5/Exchange/config"
	"TrabalhoLab5/Exchange/offerBook"
	"github.com/nats-io/go-nats"
	"log"
	"regexp"
	"strconv"
)

func printMsgSell(m *nats.Msg, orders chan offerBook.Order) {
	log.Printf("[%s] - VENDA : '%s'", m.Subject, string(m.Data))
	rgx := regexp.MustCompile(`<quant: (.*?), price: (.*?), corretora : (.*?)>`)
	rs := rgx.FindStringSubmatch(string(m.Data))
	rgx = regexp.MustCompile(`Broker\.(.*?)\.sell`)
	stock := rgx.FindStringSubmatch(string(m.Subject))
	price, _ := strconv.ParseFloat(rs[2], 64)
	qnt, _ := strconv.Atoi(rs[1])
	order := offerBook.Order{"sell", stock[1], qnt, rs[3], price}
	orders <- order
}
func Sell(orders chan offerBook.Order) {
	msgs := make(chan *nats.Msg)
	go config.Sub("*.sell", msgs)
	log.Printf("Escutando os pedidos de Venda...")
	for {
		printMsgSell(<-msgs, orders)
	}
}
