package main

import (
	"fmt"
	"math"
	"time"
)

func reverseNum(numIn int) int {
	var returnNum int
	for numIn/10 != 0 {
		returnNum = returnNum*10 + (numIn % 10)
		numIn /= 10
	}
	returnNum = returnNum*10 + (numIn % 10)
	return returnNum
}

func check(numIn, numReverse int) bool {
	if numIn > numReverse {
		gen := numIn
		numIn = numReverse
		numReverse = gen
	}
	if numReverse%numIn == 0 {
		return false
	}
	if (numIn%2 == 0) && (numReverse%2 == 0) {
		return false
	}
	if (numIn%3 == 0) && (numReverse%3 == 0) {
		return false
	}
	if (numIn%5 == 0) && (numReverse%5 == 0) {
		return false
	}
	for i := 6; float64(i) < math.Sqrt(float64(numIn)); i++ {
		if (numIn%i == 0) && (numReverse%i == 0) {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	const a = 10
	const b = 30000
	var inc int
	for i := a; i <= b; i++ {
		numReverse := reverseNum(i)
		if check(i, numReverse) {
			inc++
		}
	}
	fmt.Println(inc)
	elapsed := time.Since(start)
	fmt.Print(elapsed)
}
