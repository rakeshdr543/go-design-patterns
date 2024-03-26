package facade

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance float64
}

func newWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) creditBalance(amount float64) {
	w.balance += amount

	fmt.Println(
		"Balance: ", w.balance)

	return
}

func (w *Wallet) debitBalance(amount float64) error {
	if w.balance < amount {
		return errors.New("insufficient balance")
	}

	w.balance -= amount
	return nil
}
