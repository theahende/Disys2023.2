An overview of what has to be done in this handin.

1. Make Peer that has the following functionality
- Method ``Peer.Connect(addr string, port int)``    (√)
  - When you connect you get the set of peers from the one you connect to
- Method ``Peer.FloodMessage(msg <type we choose>)``
- Method ``Peer.FloodTransaction(tx *Transaction)``

1. Make a program handin.go that demos the system in the following way
- Starts a network with several Peers on the same network
- Starts around n = 10 Peers
- Sends 10 transactions at each Peer relating to the same 5 Accounts at the same time
- Test that all n Peers hold identical Ledgers

I'm not sure if the test we are supposed to make is the same as handin.go.
None the less we should also test we system and are encouraged to do this automatically.

3. A report containing
- A description of how we implemented the system
- A description of how we tested the system
- An example of a scenario that leads to eventual consistency
- An example of a scenario that does not lead to eventual consistency (XXXX)
- Does the system still have eventual consistency if a Transaction is rejected if the account balance goes under zero?

I hope this link becomes useful:
https://github.com/blatchley/dist-sys-Golang
