package facade

import "errors"

type Account struct {
	name string
}

func newAccount(name string) *Account {
	return &Account{
		name: name,
	}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return errors.New("account name mismatch ")
	}
	return nil
}
