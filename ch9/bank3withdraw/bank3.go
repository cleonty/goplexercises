package bank3

import (
	"sync"
)

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	deposit(amount)
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func deposit(amount int) {
	balance += amount
}
