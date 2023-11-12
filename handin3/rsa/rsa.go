package rsa

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
)

type PublicKey struct {
	N *big.Int
	E *big.Int
}
type PrivateKey struct {
	N *big.Int
	D *big.Int
}

func handleError(err error) {
	if err != nil {
		fmt.Println("some error occured", err)
		os.Exit(1)
	}
}

/*
This function generates a keypair (public and private keys)

where the public key is formed of (n,e)
and the private key is formed of (n,d)

returns PublicKey (n, e) and PrivateKey (n,d)

Check chapter 5.2.1 of the book to find the meaning of the variables used in the function
*/
func KeyGen(k int) (PublicKey, PrivateKey) {

	p, err := rand.Prime(rand.Reader, k/2) // k/2 since p*q should equal k
	handleError(err)

	q, err := rand.Prime(rand.Reader, k/2) // k/2 since p*q should equal k
	handleError(err)

	e := big.NewInt(3) // e = 3, from task description

	n := new(big.Int).Mul(p, q) //since n = p * q

	pSub1 := new(big.Int).Sub(p, big.NewInt(1)) // (p - 1)
	qSub1 := new(big.Int).Sub(q, big.NewInt(1)) // (q - 1)

	rightSideModolu := new(big.Int).Mul(pSub1, qSub1) // (p-1) * (q - 1)

	d := new(big.Int).ModInverse(e, rightSideModolu) // e^-1 mod ((p-1) * (q-1))

	if d == nil {
		//since p and q are not relative primes, the keys wouldn't be valid so we generate another pair.
		return KeyGen(k)
	} else {
		//we have generated a valid key pair, so we return the values
		return PublicKey{N: n, E: e}, PrivateKey{N: n, D: d}
	}

}

/*
This function encrypts a given message with the given public key (n,e)
and returns the ciphertext

Check chapter 5.2.1 of the book to find the math for the encryption
*/
func Encrypt(m *big.Int, pk PublicKey) *big.Int {
	c := new(big.Int).Exp(m, pk.E, pk.N) //encrypts m with the public key (n,e) to get the ciphertext c
	return c
}

/*
This function decrypts the given ciphertext with the given private key (n,d)

Check chapter 5.2.1 of the book to find the math for the decryption
*/
func Decrypt(c *big.Int, sk PrivateKey) *big.Int {
	m := new(big.Int).Exp(c, sk.D, sk.N) //decrypts ciphertext c with secret key (n,d) to get the original message m
	return m
}

/*
This function hashes a given message m and hashes it using the "crypto/sha256" library

returns the hashed message as a big.Int
*/
func HashMessage(m *big.Int) *big.Int {
	mByte := m.Bytes()    //take the message m and convert to byte array
	h := sha256.New()     //make new hash
	h.Write(mByte)        //make new hash of message
	hashOfM := h.Sum(nil) //save the hashed value

	hash_m := new(big.Int)   //make a new big int
	hash_m.SetBytes(hashOfM) //make hash back into a big Int
	return hash_m
}

/*
This function signs the given message m with the given secret key (n,d)
and returns a signature s.

Before signing the message m is hashed using sha256.

See Authenticity slides, slide 28
*/
func Sign(m *big.Int, sk PrivateKey) *big.Int {
	hash_m := HashMessage(m)
	s := new(big.Int).Exp(hash_m, sk.D, sk.N) //sign the hashed message with the secret key
	return s
}

/*
This function verifies the given signature s in terms of the message m with the given public key (n,e)

# Since the signed message was hashed we have to hash message m to verify

See Authenticity slides, slide 28
*/
func Verify(m *big.Int, s *big.Int, pk PublicKey) bool {
	hash_m := HashMessage(m)
	unpacked := new(big.Int).Exp(s, pk.E, pk.N)
	verifification := unpacked.Cmp(hash_m) == 0
	return verifification
}

func EncodePkToString(pk PublicKey) string {
	n_string := pk.N.String()
	e_string := pk.E.String()
	return n_string + e_string
}

func DecodePkToStruct(pk string) PublicKey {
	length := len(pk)
	n_str := pk[:length-1]
	e_str := pk[length-1:]

	n_bigI, success := new(big.Int).SetString(n_str, 10)

	if !success {
		panic("failed in decoding n of pk to big.Int")
	}

	e_bigI, success := new(big.Int).SetString(e_str, 10)
	if !success {
		panic("failed in decoding e of pk to big.Int")
	}

	return PublicKey{
		N: n_bigI,
		E: e_bigI,
	}
}