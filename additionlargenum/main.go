package main

import (
	"fmt"
	"strconv"
)

func createProduct(num1, num2 string, lenmax, lenmin int) []int {
	prod := make([]int, lenmax+1)
	for i := 0; i < lenmin; i++ {
		prodNum1, _ := strconv.Atoi(string(num1[i]))
		prodNum2, _ := strconv.Atoi(string(num2[i]))
		prodNum := prodNum1 + prodNum2
		prod[i] += prodNum
		if prod[i] > 10 {
			prod[i+1]++
			prod[i] %= 10
		}
	}
	for i := lenmin; i < lenmax; i++ {
		prodNum := 0
		if lenmin == len(num1) {
			prodNum, _ = strconv.Atoi(string(num2[i]))
		} else {
			prodNum, _ = strconv.Atoi(string(num1[i]))
		}
		prod[i] += prodNum
		if prod[i] > 9 {
			prod[i+1]++
			prod[i] %= 10
		}
	}
	return prod
}

func reverse(num *string) {
	numr := []byte(*num)
	for i := 0; i < len(numr)/2; i++ {
		numr[i], numr[len(numr)-i-1] = numr[len(numr)-i-1], numr[i]
	}
	*num = string(numr)
}

func main() {
	var num1 string = "1999"
	var num2 string = "9"
	lenmax := len(num1)
	lenmin := len(num2)
	if len(num1) < len(num2) {
		lenmax = len(num2)
		lenmin = len(num1)
	}
	reverse(&num1)
	reverse(&num2)
	prod := createProduct(num1, num2, lenmax, lenmin)
	start := false
	for i := len(prod) - 1; i >= 0; i-- {
		if start {
			fmt.Print(prod[i])
		} else if prod[i] != 0 {
			start = true
			fmt.Print(prod[i])
		}
	}
}
