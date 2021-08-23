package main

import "fmt"

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
}
