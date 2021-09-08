package workwithmodule

import (
	"golang.org/x/crypto/bcrypt"
)

func CryptPass(Ps string) []byte {
	Pscrypt, err := bcrypt.GenerateFromPassword([]byte(Ps), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return Pscrypt
}
