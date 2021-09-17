// A test with module in Go
package workwithmodule

import (
	"golang.org/x/crypto/bcrypt"
)

// CryptPass is used to create an bcrypt hashed password
func CryptPass(Ps string) []byte {
	Pscrypt, err := bcrypt.GenerateFromPassword([]byte(Ps), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return Pscrypt
}
