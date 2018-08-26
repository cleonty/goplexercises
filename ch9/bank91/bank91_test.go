package bank91

import (
	"testing"
)

func TestDepositWithdraw(t *testing.T) {
	type args struct {
		amount  int
		deposit bool
	}
	tests := []struct {
		name    string
		balance int
		args    args
	}{
		{"deposit 100", 100, args{100, true}},
		{"deposit 100", 200, args{100, true}},
		{"withdraw 1000", 200, args{1000, false}},
		{"withdraw 50", 150, args{50, false}},
		{"deposit 500", 650, args{500, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initialBalance := Balance()
			if tt.args.deposit {
				Deposit(tt.args.amount)
			} else {
				Withdraw(tt.args.amount)
			}
			balance := Balance()
			if balance != tt.balance {
				t.Errorf("got balance %d, expected %d\n", balance, tt.balance)
			}
		})
	}
}
