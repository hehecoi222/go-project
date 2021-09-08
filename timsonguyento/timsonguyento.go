package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	n := 3456
	for n > 0 {
		chuso := n % 10
		n /= 10
		dem := 0
		for i := 1; i <= chuso; i++ {
			if chuso%i == 0 {
				dem++
			}
		}
		if dem == 2 {
			fmt.Println(chuso)
		}
	}
	Ps := []byte(`Hello`)
	Pscrypt, err := bcrypt.GenerateFromPassword(Ps, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(Pscrypt))
}
