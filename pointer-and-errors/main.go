package ooo

import "fmt"

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

func (wallet *Wallet) Withdraw(depositAmount Bitcoin) {
	wallet.balance -= depositAmount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}
