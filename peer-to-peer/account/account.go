package account

import (
	"math/big"
	"sync"
)

/*
	P1
	P2

	P1 -> P2

	P2 gi dens encode(pk) til P1

	P1 lav transaction:

	Trams {
		From: encode(P1k)
		To:  encode(P2k)
		amount: 666
		signature: P1 name
	}
*/

type Transaction struct {
	From   string
	To     string
	Id		int
	Amount int
}

type SignedTransaction struct {
	Transaction Transaction
	Signature   *big.Int
}

type Ledger struct {
	Accounts map[string]int
	lock     sync.Mutex
}

func MakeLedger() *Ledger {
	ledger := new(Ledger)
	ledger.Accounts = make(map[string]int)
	return ledger
}

func (l *Ledger) Transact(t *Transaction) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.Accounts[t.From] -= t.Amount
	l.Accounts[t.To] += t.Amount
}
