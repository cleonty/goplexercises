package bank1

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {
	type args struct {
		amount  int
		balance int
	}
	tests := []struct {
		name string
		args args
	}{
		{"add 100 first", args{100, 100}},
		{"add 100 second", args{100, 200}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Deposit(tt.args.amount)
			fmt.Println("deposit done")
			balance := Balance()
			fmt.Println("balance done")
			if balance != tt.args.balance {
				t.Errorf("got balance %d, expected %d\n", balance, tt.args.balance)
			}
		})
	}
}
