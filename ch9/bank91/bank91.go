package bank91

import (
	"log"
)

type withdraw struct {
	amount int
	ch     chan<- bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan withdraw)

// Deposit adds amount to balance
func Deposit(amount int) { deposits <- amount }

// Withdraw decreases balance by amount
func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraws <- withdraw{amount, ch}
	return <-ch
}

// Balance returns current balance
func Balance() int { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdraw := <-withdraws:
			if balance >= withdraw.amount {
				balance -= withdraw.amount
				withdraw.ch <- true
			} else {
				withdraw.ch <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	log.Println("init")
	go teller()
}
