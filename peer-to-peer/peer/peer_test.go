package peer

import (
	"peer-to-peer/account"
	"reflect"
	"testing"
	"handin3/rsa"
	"time"
)

/*Method to create a new peer*/
func makePeer() *Peer {
	pk, sk := rsa.KeyGen(2048)
	name := rsa.EncodePkToString(pk)
	peer := Peer{
		Name:   name,
		Sk:     sk,
		Ip:     "127.0.0.1",
		Ledger: *account.MakeLedger(),
		Peers:  make(map[string] ConnectedPeer),
	}
	peer.Ledger.Accounts[peer.Name] = 200
	return &peer
}

func SignAndFloodTransaction(p *Peer, tx *account.Transaction) {
	signed_tx := p.MakeSignedTransaction(*tx)
	p.FloodSignedTransaction(signed_tx)
}

/* 
func setup() {
	p1 := makePeer()

	p2 := makePeer()

} */

/* This test checks if two peers can connect to eachother by checking if their set of peers are equal */
func TestIfTwoConnectedPeersHaveTheSamePeersList(t *testing.T) {
	p1 := makePeer()
	p2 := makePeer()

	p1.Connect(p1.Ip, p1.Port)
	p2.Connect(p1.Ip, p1.Port)

	result := reflect.DeepEqual(p1.Peers, p2.Peers)
	expected := true
	if result != expected {
		t.Errorf("\n result was: %v \n expected was: %v.\n p1.Peers = %v \n p2.Peers = %v \n", result, expected, p1.Peers, p2.Peers)
	}
}

/* Test if ledgers update correctly when sending a flood transaction*/
func TestIfTwoPeersHaveTheSameLedgerAfterFloodValidTransaction(t *testing.T) {
	p1 := makePeer()
	p2 := makePeer()

	p1.Connect(p1.Ip, p1.Port)
	time.Sleep(time.Second)
	p2.Connect(p1.Ip, p1.Port)
	time.Sleep(time.Second*3)

	tx1 := account.Transaction{
		From:   p1.Name,
		To:     p2.Name,
		Id:		1,
		Amount: 100,
	}


	SignAndFloodTransaction(p1, &tx1)
	time.Sleep(time.Second*3)

	expected_amount_p1 := 100
	expected_amount_p2 := 300
	result_ledger := reflect.DeepEqual(&p1.Ledger, &p2.Ledger)
	result_amount_p1 := reflect.DeepEqual(p1.Ledger.Accounts[p1.Name], expected_amount_p1)
	result_amount_p2 := reflect.DeepEqual(p1.Ledger.Accounts[p2.Name], expected_amount_p2)
	result := result_ledger && result_amount_p1 && result_amount_p2
	expected := true

	if result != expected {
		t.Errorf("\n result was: %v \n expected was: %v.\n p1.Ledger = %v \n p2.Ledger = %v \n", result, expected, &p1.Ledger, &p2.Ledger)
	}

}

/* Test if ledgers update correctly when sending an invalid flood transaction*/
func TestIfTwoPeersHaveTheSameLedgerAfterFloodInvalidTransaction(t *testing.T) {
	p1 := makePeer()
	p2 := makePeer()

	p1.Connect(p1.Ip, p1.Port)
	time.Sleep(time.Second)
	p2.Connect(p1.Ip, p1.Port)
	time.Sleep(time.Second*3)

	tx1 := account.Transaction{
		From:   p1.Name,
		To:     p2.Name,
		Id:		1,
		Amount: 100,
	}


	SignAndFloodTransaction(p2, &tx1)
	time.Sleep(time.Second*3)

	expected_amount_p1 := 200
	expected_amount_p2 := 200
	result_ledger := reflect.DeepEqual(&p1.Ledger, &p2.Ledger)
	result_amount_p1 := reflect.DeepEqual(p1.Ledger.Accounts[p1.Name], expected_amount_p1)
	result_amount_p2 := reflect.DeepEqual(p1.Ledger.Accounts[p2.Name], expected_amount_p2)
	result := result_ledger && result_amount_p1 && result_amount_p2
	expected := true

	if result != expected {
		t.Errorf("\n result was: %v \n expected was: %v.\n p1.Ledger = %v \n p2.Ledger = %v \n", result, expected, &p1.Ledger, &p2.Ledger)
	}

}