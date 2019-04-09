package pub

import (
	"TrabalhoLab5/Broker/config"
	"strconv"
)

func Sell(stock string, qnt int, price float64, broker string) {
	config.Pub(stock+".sell", []byte("<quant: "+strconv.Itoa(qnt)+", price: "+strconv.FormatFloat(price, 'f', -1, 64)+", corretora : "+broker+">"))
}
