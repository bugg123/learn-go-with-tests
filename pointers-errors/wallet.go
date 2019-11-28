package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.balance += deposit
}

func (w *Wallet) Withdraw(withdraw Bitcoin) error {
	if withdraw > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= withdraw
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
