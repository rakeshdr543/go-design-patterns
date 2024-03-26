package facade

import (
	"fmt"
)

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func NewWalletFacade(accountId string, code int) *WalletFacade {
	fmt.Println("creating new account")

	walletFacade := &WalletFacade{
		account:      newAccount(accountId),
		wallet:       newWallet(),
		securityCode: newSecurityCode(code),
		notification: &Notification{},
		ledger:       &Ledger{},
	}

	fmt.Println("created new account")
	return walletFacade
}

func (wallet *WalletFacade) AddMoneyToAccount(code int, amount float64) error {
	fmt.Println(" Adding money to account ")

	err := wallet.account.checkAccount(wallet.account.name)

	if err != nil {
		return err
	}

	err = wallet.securityCode.checkSecurityCode(code)

	if err != nil {
		return err
	}

	wallet.wallet.creditBalance(amount)

	wallet.notification.sendCreditNotification()
	wallet.ledger.createLedger(wallet.account.name, "credit", amount)
	return nil
}

func (wallet *WalletFacade) DeductMoneyFromAccount(code int, amount float64) error {
	fmt.Println("Deducting money from account")

	err := wallet.account.checkAccount(wallet.account.name)

	if err != nil {
		return err
	}

	err = wallet.securityCode.checkSecurityCode(code)
	if err != nil {
		return err
	}

	err = wallet.wallet.debitBalance(amount)

	if err != nil {
		return err
	}

	wallet.notification.sendDebitNotification()

	wallet.ledger.createLedger(wallet.account.name, "debit", amount)

	return nil
}
