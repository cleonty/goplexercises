package bank1

import (
	"log"
)

var deposits = make(chan int)
var balances = make(chan int)

// Deposit adds amount to balance
func Deposit(amount int) { deposits <- amount }

// Balance returns current balance
func Balance() int { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	log.Println("init")
	go teller()
}
