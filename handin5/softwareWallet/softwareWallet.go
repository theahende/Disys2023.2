package softwareWallet

import (
	"handin3/aes"
	"handin3/rsa"
	"math/big"
)

type Signature struct {
	Name      string
	Signature big.Int
}

func generateHashOnPassword(pwd *big.Int, n int) *big.Int {
	var hash_pw *big.Int
	for i := 0; i < n; i++ {
		hash_pw = rsa.HashMessage(pwd)
	}
	return hash_pw
}

/*
This function generates a sofware wallet, by encrypting a secret rsa key to a file and returning the public key (n, e)
*/
func Generate(filename string, password string) rsa.PublicKey {
	//generate keypair pkey (n,e) skey (n,d) values
	pk, sk := rsa.KeyGen(1024)

	//hash password
	pwd_bytes := []byte(password)                 //convert string to bytes
	pwd_b_int := new(big.Int).SetBytes(pwd_bytes) //convert to big.Int
	hash_pw := generateHashOnPassword(pwd_b_int, 1024)

	dbytes := sk.D.Bytes() //convert secret part of key to bytes
	//encrypt password with hashed password
	aes.EncryptToFile(dbytes, filename, hash_pw.Bytes())

	return pk
}

/*
This function signs a msg using a software wallet, if the correct password has been given

returns the signature if the correct password was used, else nil
*/
func Sign(filename string, password string, pk rsa.PublicKey, msg []byte) *big.Int /* return type = something we define */ {
	//hash password
	pwd_bytes := []byte(password)
	pwd_b_int := new(big.Int).SetBytes(pwd_bytes) //convert to big.Int
	hash_pw := generateHashOnPassword(pwd_b_int, 1024)

	d_bytes := aes.DecryptFromFile(filename, hash_pw.Bytes())
	sk := rsa.PrivateKey{
		N: pk.N,
		D: new(big.Int).SetBytes(d_bytes),
	}

	m_bytes := new(big.Int).SetBytes(msg)

	signature := rsa.Sign(m_bytes, sk)

	is_correct_signature := rsa.Verify(m_bytes, signature, pk)

	if is_correct_signature {
		return signature
	} else {
		return nil
	}

}
