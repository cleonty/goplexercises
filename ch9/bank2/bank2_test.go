package bank2

import (
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i <= 1000; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			Deposit(amount)
		}(i)
	}
	wg.Wait()

	if got, want := Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
