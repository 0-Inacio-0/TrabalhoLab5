package sub

import (
	"TrabalhoLab5/Exchange/config"
	"TrabalhoLab5/Exchange/offerBook"
	"github.com/nats-io/go-nats"
	"log"
	"regexp"
	"strconv"
)

func buyOrder(m *nats.Msg, orders chan offerBook.Order) {
	log.Printf("[%s] - COMPRA : '%s'", m.Subject, string(m.Data))
	rgx := regexp.MustCompile(`<quant: (.*?), price: (.*?), corretora : (.*?)>`)
	rs := rgx.FindStringSubmatch(string(m.Data))
	rgx = regexp.MustCompile(`Broker\.(.*?)\.buy`)
	stock := rgx.FindStringSubmatch(string(m.Subject))
	price, _ := strconv.ParseFloat(rs[2], 64)
	qnt, _ := strconv.Atoi(rs[1])
	order := offerBook.Order{"buy", stock[1], qnt, rs[3], price}
	orders <- order
}
func Buy(orders chan offerBook.Order) {
	msgs := make(chan *nats.Msg)
	go config.Sub("*.buy", msgs)
	log.Printf("Escutando os pedidos de Compra...")
	for {
		buyOrder(<-msgs, orders)
	}

}
