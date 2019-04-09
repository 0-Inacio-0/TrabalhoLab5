// Broker sends requests to buy, sell and request info and
// logs all the transactions that happens on the exchange.
package main

import (
	"TrabalhoLab5/Broker/pub"
	"TrabalhoLab5/Broker/sub"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"
)

func main() {
	var k int
	var brk string
	fmt.Print("Qual o nome da sua corretora?\n")
	_, err := fmt.Scanf("%s\n", &brk)
	if err != nil {
		log.Fatal(err)
	}
	for k != -1 {
		fmt.Print("Digite a opção desejada:\n")
		fmt.Print("   1:Comprar\n   2:Vender\n   3:Info\n   4:Escutar Transações\n   -1:Sair\n")
		_, err = fmt.Scanf("%d\n", &k)
		if err != nil {
			log.Fatal(err)
		}
		switch k {
		case -1:
			fmt.Print("Saindo... Volte Sempre!\n")
			break
		case 1:
			stockPublisher(0, brk)
		case 2:
			stockPublisher(1, brk)
		case 3:
			go pub.Info(time.Now())
		case 4:
			go sub.Trans()
		default:
			fmt.Print("Opção invalida!\n")
			continue
		}
	}
	runtime.Goexit()
}
func stockPublisher(op int, brk string) {
	stocks := [...]string{"ABEV3", "PETR4", "VALE5", "ITUB4", " BBDC4", "BBAS3", "CIEL3", "PETR3", "HYPE3", "VALE3", "BBSE3", "CTIP3", "GGBR4", "FIBR3", "RADL3"}
	opts := [...]string{"comprar", "vender"}
	var stock string
	var qnt int
	var price float64

	fmt.Print("Qual ação deseja " + opts[op] + "?\n")
	for i, ele := range stocks {
		fmt.Print("   " + strconv.Itoa(i+1) + ": " + ele + "\n")
	}
	k := 0
	_, err := fmt.Scanf("%d\n", &k)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Você selecionou a ação " + stocks[k-1] + " para " + opts[op] + ".\n")
	stock = stocks[k-1]
	fmt.Print("Quantas ações você deseja " + opts[op] + "?\n")
	_, err = fmt.Scanf("%d\n", &qnt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Por qual valor você deseja " + opts[op] + "?\n")
	_, err = fmt.Scanf("%f\n", &price)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("")
	fmt.Print("Realmente deseja " + opts[op] + " " + strconv.Itoa(qnt) + "x " + stock + " por R$" + strconv.FormatFloat(price, 'f', -1, 64) + "(Y/n)?\n")
	var conf string
	_, err = fmt.Scanf("%s\n", &conf)
	if err != nil {
		log.Fatal(err)
	}
	if conf == "Y" {
		if op == 0 {
			go pub.Buy(stock, qnt, price, brk)
		}
		if op == 1 {
			pub.Sell(stock, qnt, price, brk)
		}
	} else {
		fmt.Print("Operação abortada!\n")
	}
}
