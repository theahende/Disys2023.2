package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"os"
)

/*
takes in a RSA key and encrpyts it to a file.
*/
func EncryptToFile(rsa_key []byte, fileName string, aes_key []byte) {
	//make aes to block
	aes_block, _ := aes.NewCipher(aes_key) //make new cipherblock of aes key

	//make ciphertext and nonce
	ciphertext := make([]byte, aes.BlockSize+len(rsa_key))
	iv := ciphertext[:aes.BlockSize] //nonce is a part of the ciphertext
	rand.Read(iv)                    //generates random values for nonce

	//make new cipher stream from aes block and iv
	ctr := cipher.NewCTR(aes_block, iv)

	//take AES part from ciphertext and XOR with RSA key
	aes_part := ciphertext[aes.BlockSize:]
	ctr.XORKeyStream(aes_part, rsa_key)

	//writing the file
	os.WriteFile(fileName, ciphertext, 0644) //de giver bare en konstant med i den sidste parameter, magic stuff
}

/*
takes a filename and an aes key
*/
func DecryptFromFile(filename string, aes_key []byte) []byte {
	//make ciphertext by reading from file
	ciphertext, _ := os.ReadFile(filename)
	//make aes to block from key
	aes_block, _ := aes.NewCipher(aes_key)

	//make iv by taking the last part of the ciphertext
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	//new CTR whatever that is
	ctr := cipher.NewCTR(aes_block, iv)
	rsa_key := make([]byte, len(ciphertext))

	//XOR RSA key with the ciphertext (without the IV)
	ctr.XORKeyStream(rsa_key, ciphertext)

	return rsa_key
}
