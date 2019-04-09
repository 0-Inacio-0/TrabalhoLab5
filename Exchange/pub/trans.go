package pub

import (
	"TrabalhoLab5/Exchange/config"
	"strconv"
)

func Trans(stock string, qnt int, price float64) {
	pricePer := price / float64(qnt)
	config.Pub("transacao."+stock, []byte("<quant: "+strconv.Itoa(qnt)+", price: "+strconv.FormatFloat(pricePer, 'f', -1, 64)+", total: "+strconv.FormatFloat(price, 'f', -1, 64)+">"))
}
