package main

import (
	"fmt"
	"time"
)

func count(num, Llimit, Hlimit int) int {
	if num == 0 {
		return 0
	}
	return (Hlimit / num) - (Llimit / num)
}

func main() {
	start := time.Now()
	const x = 2
	const a = 1
	const b = 10e10
	fmt.Println(count(x, a, b))
	elapsed := time.Since(start)
	fmt.Print(elapsed)
}
