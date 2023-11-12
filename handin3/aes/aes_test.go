package aes

import (
	"handin3/rsa"
	"reflect"
	"testing"
)

func TestEncryptAndDecryptFile(t *testing.T) {
	_, sk := rsa.KeyGen(128) //128 bit
	//nBytes := n.Bytes()
	dBytes := sk.D.Bytes()

	//rsa_k := append(nBytes, dBytes...)

	aes_k := make([]byte, 32)

	EncryptToFile(dBytes, "filename", aes_k)
	dec_res := DecryptFromFile("filename", aes_k)

	/* 	println(new(big.Int).SetBytes(dBytes).String())
	   	println(new(big.Int).SetBytes(dec_res).String())
	*/
	decryptedFileIsEqualToDBytes := reflect.DeepEqual(dec_res, dBytes)

	if !decryptedFileIsEqualToDBytes {
		t.Errorf("\n Decrypted file is not equal to original message \n decrypted file: %v \n original message: %v \n", dec_res, dBytes)
	}
}
