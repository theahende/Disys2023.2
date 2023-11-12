# Handin 6 (exercise 6.13)

## Ledger

### Overview:

- a client can make a 'SignedTransactions' which are broadcast of the type 'SignedTransaction'
  - The ledger only handles valid transactions

## Transactions

### Overview:

- the name of the sender and the receiver transactions are RSA public keys **encoded as strings**
- the client can only make a transaction if it knows the secret key of the sending account (you can only take money from your own account)

## Accounts

- an account is generated as a random RSA key-pair (pk,sk)
- encode(pk) is the account which can only be accessed with the corresponding secret key
- to transfer money TOO an account you must know encode(pk)