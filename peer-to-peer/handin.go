package main

//This file should contain a demo of the system

import (
	"encoding/json"
	"fmt"
	"handin3/rsa"
	"log"
	"peer-to-peer/account"
	"peer-to-peer/peer"
	"time"
)

/* func PopulateLedger(l *account.Ledger) {
	l.Accounts["account1"] = 0
	l.Accounts["account2"] = 0
	l.Accounts["account3"] = 0
	l.Accounts["account4"] = 0
	l.Accounts["account5"] = 0
} */

func makePeer() *peer.Peer {
	//TODO:
	//our code only works when the argument in Keygen is a multiple of 2 that's greater than or equal to 256
	//I do not understand why
	pk, sk := rsa.KeyGen(256)
	name := rsa.EncodePkToString(pk)
	peer := peer.Peer{
		Name:   name,
		Sk:     sk,
		Ip:     "127.0.0.1",
		Ledger: *account.MakeLedger(),
		Peers:  make(map[string]peer.ConnectedPeer),
	}
	peer.Ledger.Accounts[peer.Name] = 200
	return &peer
}

/* this is a pretty good pretty printer for structs xD */
func PrintJsonString(s any) {
	JSON, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(JSON))
}

/* func sendTransactionTest(peer *peer.Peer) {

	for i := 1; i <= 4; i++ {
		var transaction account.Transaction = account.Transaction{}
		transaction.From = "account" + strconv.Itoa(i)
		transaction.To = "account" + strconv.Itoa(i+1)
		transaction.Amount = 1

		peer.FloodTransaction(&transaction)
	}

	for i := 1; i <= 4; i++ {
		var transaction account.Transaction = account.Transaction{}
		transaction.From = "account" + strconv.Itoa(i+1)
		transaction.To = "account" + strconv.Itoa(i)
		transaction.Amount = 2

		peer.FloodTransaction(&transaction)
	}

	for i := 2; i <= 3; i++ {
		var transaction account.Transaction = account.Transaction{}
		transaction.From = "account" + strconv.Itoa(i)
		transaction.To = "account" + strconv.Itoa(i+1)
		transaction.Amount = 3

		peer.FloodTransaction(&transaction)
	}

} */

func main() {

	/*
		make 10 peers
		10 transactions send per peer
		relating to all 5 accounts

	*/

	/* CREATE PEERS */
	p1 := makePeer()
	p2 := makePeer()
	p3 := makePeer()

	p1.Connect(p1.Ip, p1.Port)
	time.Sleep(time.Second)
	p2.Connect(p1.Ip, p1.Port)
	time.Sleep(time.Second)
	p3.Connect(p2.Ip, p2.Port)
	time.Sleep(time.Second)

	tx1 := account.Transaction{
		From:   p1.Name,
		To:     p2.Name,
		Amount: 100,
	}
	
	tx2 := account.Transaction{
		From:   p2.Name,
		To:     p3.Name,
		Amount: 500,
	}
	
	time.Sleep(time.Second)
	p1.FloodSignedTransaction(p1.MakeSignedTransaction(tx1))
	time.Sleep(time.Second)
	p1.FloodSignedTransaction(p1.MakeSignedTransaction(tx2))
	time.Sleep(time.Second * 5)

	

	fmt.Println(p1.Name)
	PrintJsonString(p1)
	
	p1.FloodSignedTransaction(p1.MakeSignedTransaction(tx1))
	time.Sleep(time.Second)
	p1.FloodSignedTransaction(p1.MakeSignedTransaction(tx2))
	time.Sleep(time.Second * 5)

	fmt.Println(p1.Name)
	PrintJsonString(p1.Ledger.Accounts)
	fmt.Println(p2.Name)
	PrintJsonString(p2.Ledger.Accounts)
	fmt.Println(p3.Name)
	PrintJsonString(p3.Ledger.Accounts)


	for {
		continue
	}
}