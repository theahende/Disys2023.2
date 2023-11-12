package rsa

import (
	"crypto/rand"
	"math/big"
	"reflect"
	"testing"
)

func TestRsaEncryptionAndDecryptionShouldGiveSameMessage(t *testing.T) {
	//generate a public key (n,e) and private key (n,d)
	pk, sk := KeyGen(1024)
	//make a message
	message := big.NewInt(1234659)

	//encrypt the message (generate ciphertext for the message)
	c := Encrypt(message, pk)

	//Decrypt the ciphertext to get the message
	decrypted_message := Decrypt(c, sk)

	//since *big.Int values, represent arbitrary-precision integers.
	// we use the Cmp function from the math/big library to compare them
	MessageIsEqualToDecryptedMessage := message.Cmp(decrypted_message) == 0
	if !MessageIsEqualToDecryptedMessage {
		//message and decrypted message are converted to string inorder to for us to better understand their values
		t.Errorf("\n Message and Decrypted message didn't match.\n message = %s \n decrypted_message = %s \n", message.String(), sk.D.String())
	}
}

func TestRsaSignAndVerify(t *testing.T) {
	//generate a public key (n,e) and private key (n,d)
	pk, sk := KeyGen(1024)
	//make a message
	message := big.NewInt(1234659)

	s := Sign(message, sk)

	verificationWorks := Verify(message, s, pk)

	if !verificationWorks {
		t.Errorf("\n Message and unsigned message don't match.\n")
	}
}

func BenchmarkHashMessage(b *testing.B) {
	// Calculate the number of bytes needed for 10KB (10 * 1024 bytes)
	var dataSize = 10 * 1024

	// Create a byte slice with the desired size
	var data = make([]byte, dataSize)
	// fill slice with random values
	rand.Read(data)

	// Create a big.Int from the bytes in the data slice
	var messageToHash = new(big.Int).SetBytes(data)

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		HashMessage(messageToHash)
	}
	b.StopTimer()
}

/*
Measure the time you code spends to produce an RSA signature on a hash value when
using a 2000 bit RSA key
*/
func BenchmarkRsaSignatureOnAHashValueUsing2000BitRsaKey(b *testing.B) {
	// Calculate the number of bytes needed for 10KB (10 * 1024 bytes)
	var dataSize = 10 * 1024
	// Create a byte slice with the desired size
	var data = make([]byte, dataSize)
	// fill slice with random values
	rand.Read(data)
	// Create a big.Int from the bytes in the data slice
	var messageToHash = new(big.Int).SetBytes(data)

	//making a rsa-key of 2000 bits
	_, sk := KeyGen(2000)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		hash_res := HashMessage(messageToHash)
		Sign(hash_res, sk)
	}
	b.StopTimer()
}

func BenchmarkEntireMessageUsingRsa(b *testing.B) {
	// Calculate the number of bytes needed for 10KB (10 * 1024 bytes)
	var dataSize = 10 * 1024
	// Create a byte slice with the desired size
	var data = make([]byte, dataSize)
	// fill slice with random values
	rand.Read(data)

	var message = new(big.Int).SetBytes(data)

	//making a rsa-key of 2000 bits
	_, sk := KeyGen(2000)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Sign(message, sk)
	}
	b.StopTimer()
}

func TestEncodeAndDecodePk(t *testing.T) {
	pk, _ := KeyGen(126)
	pk_str := EncodePkToString(pk)
	pk1 := DecodePkToStruct(pk_str)

	if !reflect.DeepEqual(pk, pk1) {
		t.Errorf("Keys weren't of the same length: %v, %v", pk, pk1)
	}
}
