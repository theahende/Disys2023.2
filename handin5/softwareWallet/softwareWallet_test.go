package softwareWallet

import (
	"testing"
)

func TestGenerateAndSignValid(t *testing.T) {
	fil_navn := "sejfil"
	password := "hejmeddig"
	msg := "disys er lidt svært"
	msg_bytes := []byte(msg)

	//encrypt password and secret key to file and return public key
	pk := Generate(fil_navn, password)

	//try to sign message
	res := Sign(fil_navn, password, pk, msg_bytes)

	if res == nil {
		t.Errorf("shit occured")
	}

}

func TestGenerateAndSignInvalid(t *testing.T) {
	fil_navn := "sejfil"

	password_correct := "hejmeddig"
	password_incorrect := "hejsameddigsa"

	msg := "disys er lidt svært"
	msg_bytes := []byte(msg)

	//encrypt password and secret key to file and return public key
	pk := Generate(fil_navn, password_correct)

	//try to sign message
	res := Sign(fil_navn, password_incorrect, pk, msg_bytes)

	//the things shouldn't be equal
	if res != nil {
		t.Errorf("shit occured")
	}

}
