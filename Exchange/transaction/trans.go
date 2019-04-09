package transaction

import (
	"TrabalhoLab5/Exchange/pub"
	"log"
	"time"
)

type Transaction struct {
	Time       time.Time
	Stock      string
	Qnt        int
	TotalPrice float64
}

type transRecord struct {
	list  []Transaction
	trans chan Transaction
}

func Transactions(trans chan Transaction) {
	errChan := make(chan error)
	defer close(errChan)
	var rec transRecord
	rec.trans = trans
	for {
		select {
		case err := <-errChan:
			log.Print(err)
			return
		case t := <-rec.trans:
			rec.list = append(rec.list, t)
			pub.Trans(t.Stock, t.Qnt, t.TotalPrice)
		}

	}
}
