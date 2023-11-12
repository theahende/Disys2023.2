package peer

import (
	"encoding/gob"
	"fmt"
	"net"
	"peer-to-peer/account"
	"strconv"
)

type MessageStruct struct {
	Message int
	Data    []string
}
type MsgAccount struct {
	Name   string
	Amount int
}
type MessageStruct2 struct {
	Message            int
	ConnectedPeers     []ConnectedPeer
	SignedTransactions []account.SignedTransaction
	Accounts           []MsgAccount
}

/* Returns a string array containing the information of a connected peer. */
func prepareConnectedPeerWithBalanceToSend(p ConnectedPeer, l *account.Ledger) []string {
	accBalance := l.Accounts[p.Name]
	return []string{p.Name, p.Ip, strconv.Itoa(p.Port), strconv.Itoa(accBalance)}
}

func sendResponse(response MessageStruct2, conn net.Conn) {
	encoder := gob.NewEncoder(conn)
	err := encoder.Encode(response)
	if err != nil {
		fmt.Println("Couldn't send response with error: ", err)
		return
	}
}

func sendRequest(msg MessageStruct2, conn net.Conn) {
	encoder := gob.NewEncoder(conn)
	err := encoder.Encode(msg)
	if err != nil {
		fmt.Println("Couldn't send request with error: ", err)
		return
	}
}

/* Function to handle the client side */
func clientHandling(p *Peer, port int) {
	conn, err := net.Dial("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Couldn't connect to another client", err)
		return
	}
	defer conn.Close()

	sendMsg := MessageStruct2{
		Message: GetConnectedPeersRequest,
	}
	sendRequest(sendMsg, conn)
	
	handleResponse(p, conn)
}

/* Function to handle the server side */
func serverHandling(p *Peer) {

	ln, err := net.Listen("tcp", "") // Choose a random port to listen on
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	defer ln.Close()

	//fmt.Printf("I %s  am listening on the port: %d \n", p.Name, ln.Addr().(*net.TCPAddr).Port)

	//Get the port from the ln-object
	p.Port = ln.Addr().(*net.TCPAddr).Port
	//Add yourself to your own set of peers
	p.Peers[p.Name] = ConnectedPeer{
		Name: p.Name,
		Ip:   p.Ip,
		Port: p.Port,
	}

	/* Handle each connection we receive */
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept connection failed", err)
			continue
		}
		/* Handle the message we have received and send an appropriate answer to the client */
		handleRequest(p, conn)
	}
}
