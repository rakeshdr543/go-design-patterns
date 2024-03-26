package facade

import "fmt"

type Ledger struct {
}

func (l *Ledger) createLedger(accountName string, txType string, amount float64) {
	fmt.Printf(" creating ledger for account %s, tx type %s, amount %d ", accountName, txType, amount)

}
