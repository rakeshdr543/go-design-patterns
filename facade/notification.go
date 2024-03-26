package facade

import "fmt"

type Notification struct {
}

func (n *Notification) sendCreditNotification() {
	fmt.Println("Credit Notification")
}

func (n *Notification) sendDebitNotification() {
	fmt.Println("Debit Notification")
}
