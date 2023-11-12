package peer

import (
	"encoding/gob"
	"fmt"
	"handin3/rsa"
	"math/big"
	"net"
	"strconv"
)

/* enums for requests and responses */
const (
	GetConnectedPeersRequest = iota
	MakeSignedTransactionRequest
	JoinRequest

	GetConnectedPeersResponse
)

func decodeMsg(conn net.Conn) MessageStruct2 {
	var msg MessageStruct2
	decoder := gob.NewDecoder(conn)
	decodeErr := decoder.Decode(&msg)
	if decodeErr != nil {
		fmt.Println("Couldn't decode msg, got following error: ")
		panic(decodeErr)
	}
	return msg
}

func handleRequest(p *Peer, conn net.Conn) {
	msg := decodeMsg(conn)
	switch msg.Message {
	case GetConnectedPeersRequest:
		msg.Message = GetConnectedPeersResponse
		for _, cp := range p.Peers {
			//msg.Data = append(msg.ConnectedPeers, prepareConnectedPeerWithBalanceToSend(v, &p.Ledger)...)
			msg.ConnectedPeers = append(msg.ConnectedPeers, cp)
			msg.Accounts = append(msg.Accounts, MsgAccount{
				Name:   cp.Name,
				Amount: p.Ledger.Accounts[cp.Name],
			})
		}
		sendResponse(msg, conn)

	case JoinRequest: /* this request doesn't require a response */
		connPeer := msg.ConnectedPeers[0]
		p.Peers[connPeer.Name] = connPeer
	
		acc := msg.Accounts[0]
		p.Ledger.Accounts[acc.Name] = acc.Amount

	case MakeSignedTransactionRequest: /* this request doesn't require a response */
		for _ , stx := range msg.SignedTransactions {
			tx := stx.Transaction
			message := tx.From + tx.To + strconv.Itoa(tx.Id) + strconv.Itoa(tx.Amount)
			message_bigI, _ := new(big.Int).SetString(message, 10)
			signature_bigI, _ := new(big.Int).SetString(stx.Signature.String(), 10)
			fromPk := rsa.DecodePkToStruct(tx.From)
			isValid := rsa.Verify(message_bigI, signature_bigI, fromPk)
			
			if isValid {
				p.Ledger.Transact(&tx)
			}
		}
	}
}

func handleResponse(p *Peer, conn net.Conn) {
	respondedMsg := decodeMsg(conn)

	switch respondedMsg.Message {
	case GetConnectedPeersResponse:
		for _ , connPeer := range respondedMsg.ConnectedPeers{
			//connPeer := makeConnectedPeer(respondedMsg.Data[i], respondedMsg.Data[i+1], respondedMsg.Data[i+2]) /* i = name, i+1 = ip, i+2 = port, i+3 = account */
			p.Peers[connPeer.Name] = connPeer
		}
		for _ , acc := range respondedMsg.Accounts{
			p.Ledger.Accounts[acc.Name] = acc.Amount
		}
		/* Flooding a message to the network that the peer have connected to the network */
		joinMsg := MessageStruct2{
			Message: JoinRequest,
			ConnectedPeers: []ConnectedPeer{p.Peers[p.Name]},
			Accounts: []MsgAccount{
				{
					Name: p.Name,
					Amount: p.Ledger.Accounts[p.Name],
				},
			},
		}
		p.FloodMessage(joinMsg)

	}
}
