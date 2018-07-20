package ooo

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (bitcoin Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", bitcoin)
}

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(depositAmount Bitcoin) {
	wallet.balance += depositAmount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (wallet *Wallet) Withdraw(withdrawAmount Bitcoin) error {
	if withdrawAmount > wallet.balance {
		return ErrInsufficientFunds
	}

	wallet.balance -= withdrawAmount
	return nil
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}
