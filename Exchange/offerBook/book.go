package offerBook

import (
	"TrabalhoLab5/Exchange/transaction"
	"errors"
	"log"
	"time"
)

type Order struct {
	OrderType string
	Stock     string
	Qnt       int
	Brk       string
	Price     float64
}

type stockBook struct {
	stock     string
	orderChan chan Order
}

// OfferBook initializes all stock books and keep them updated with new orders in a fan-out pattern
func OfferBook(orders chan Order, trans chan transaction.Transaction) {
	errChan := make(chan error)
	quit := make(chan bool)
	defer close(quit)
	defer close(errChan)
	stocks := [...]string{"ABEV3", "PETR4", "VALE5", "ITUB4", " BBDC4", "BBAS3", "CIEL3", "PETR3", "HYPE3", "VALE3", "BBSE3", "CTIP3", "GGBR4", "FIBR3", "RADL3"}
	var bks []stockBook

	for _, stck := range stocks {
		bk := stockBook{stock: stck, orderChan: make(chan Order)}
		bks = append(bks, bk)
		go bookKeeper(bk, trans, errChan, quit)
	}

	for {
		select {
		case err := <-errChan:
			quit <- true
			log.Print(err)
			return
		case ord := <-orders:
			findBkChan(bks, ord.Stock) <- ord
		}
	}
}

func bookKeeper(bk stockBook, trans chan transaction.Transaction, errChan chan error, quitChan chan bool) {
	var buy, sell []Order
	for {
		select {
		case <-quitChan:
			return
		case ord := <-bk.orderChan:
			if ord.OrderType == "buy" {
				if len(buy) == 0 {
					buy = append(buy, ord)
				} else {
					for i, buyOrd := range buy {
						if ord.Price > buyOrd.Price {
							temp := buy[i:]
							if i != 0 {
								buy = append(buy[:i-1], ord)
							} else {
								buy = []Order{ord}
							}

							buy = append(buy, temp...)
							break
						}
					}
				}
			} else if ord.OrderType == "sell" {
				if len(sell) == 0 {
					sell = append(sell, ord)
				} else {
					for i, sellOrd := range sell {
						if ord.Price < sellOrd.Price {
							temp := sell[i:]
							if i != 0 {
								sell = append(buy[:i-1], ord)
							} else {
								sell = []Order{ord}
							}
							sell = append(sell, temp...)
							break
						}
					}
				}
			} else {
				errChan <- errors.New("OrderType does not exist.(OrderType:" + ord.OrderType + " stock:" + ord.Stock + ")")
			}
			// check for transaction
			if len(buy) != 0 && len(sell) != 0 {
				for buy[0].Price >= sell[0].Price {
					qnt := 0
					totalPrc := 0.0
					if buy[0].Qnt > sell[0].Qnt {
						qnt = sell[0].Qnt
						totalPrc = sell[0].Price * float64(qnt)
						buy[0].Qnt = buy[0].Qnt - qnt
						sell = sell[1:]
					} else if sell[0].Qnt > buy[0].Qnt {
						qnt = buy[0].Qnt
						totalPrc = sell[0].Price * float64(qnt)
						sell[0].Qnt = sell[0].Qnt - qnt
						buy = buy[1:]
					} else {
						qnt = buy[0].Qnt
						totalPrc = sell[0].Price * float64(qnt)
						buy = buy[1:]
						sell = sell[1:]
					}
					trans <- transaction.Transaction{Qnt: qnt, TotalPrice: totalPrc, Stock: ord.Stock, Time: time.Now()}
					if len(buy) == 0 || len(sell) == 0 {
						break
					}
				}
			}
		}
	}

}
func findBkChan(bks []stockBook, stock string) chan Order {
	for _, bk := range bks {
		if bk.stock == stock {
			return bk.orderChan
		}
	}
	return nil
}
