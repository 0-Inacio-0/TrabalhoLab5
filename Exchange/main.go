// Broker sends requests to buy, sell and request info and
// logs all the transactions that happens on the exchange.
package main

import (
	"TrabalhoLab5/Exchange/offerBook"
	"TrabalhoLab5/Exchange/sub"
	"TrabalhoLab5/Exchange/transaction"
	"runtime"
)

func main() {
	order := make(chan offerBook.Order)
	trans := make(chan transaction.Transaction)
	go offerBook.OfferBook(order, trans)
	go transaction.Transactions(trans)
	go sub.Buy(order)
	go sub.Sell(order)
	go sub.Info()
	runtime.Goexit()
}
