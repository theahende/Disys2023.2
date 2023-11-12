package peer

import (
	"fmt"
	"handin3/rsa"
	"math/big"
	"net"
	"peer-to-peer/account"
	"strconv"
)

type Peer struct {
	Name   string /* the name of the peer is its encoded publicKey */
	Sk     rsa.PrivateKey
	Ip     string
	Port   int
	Ledger account.Ledger
	Peers  map[string]ConnectedPeer
}

type ConnectedPeer struct {
	Name string
	Ip   string
	Port int
}

func (p *Peer) Connect(addr string, port int) {
	go clientHandling(p, port)
	go serverHandling(p)
}

func (p *Peer) FloodMessage(msg MessageStruct2) {
	connectedPeers := p.Peers
	for _, v := range connectedPeers {
		conn, err := net.Dial("tcp", ":"+strconv.Itoa(v.Port))
		if err != nil {
			fmt.Println("Couldn't connect to another client in order to flood message: ", err)
			return
		}

		sendRequest(msg, conn)

		defer conn.Close()
	}

}

func (p *Peer) FloodSignedTransaction(stx account.SignedTransaction) {
	msg := MessageStruct2 {
		Message: MakeSignedTransactionRequest,
		SignedTransactions: []account.SignedTransaction{stx},
	}

	p.FloodMessage(msg)
}

func(p *Peer) MakeSignedTransaction(tx account.Transaction) account.SignedTransaction {
	if tx.Amount <= 0 {
		fmt.Println("Can't transfer a negative amount")
		panic("Negative amount")
	}
	
	message := tx.From + tx.To + strconv.Itoa(tx.Id) + strconv.Itoa(tx.Amount)
	msg_bigI, _ := new(big.Int).SetString(message, 10)
	signature := rsa.Sign(msg_bigI, p.Sk)
	return account.SignedTransaction{
		Transaction: tx,
		Signature: signature,
	}
}
