package main

import (
	"fmt"
	"strconv"
)

func reverse(inp *string) {
	tin := []byte(*inp)
	for i := 0; i < len(tin)/2; i++ {
		tin[i], tin[len(tin)-i-1] = tin[len(tin)-i-1], tin[i]
	}
	*inp = string(tin)
}

func createProduct(tnum1, tnum2 string) []int {
	prod := make([]int, len(tnum1)+len(tnum2))
	for i := range tnum1 {
		for j := range tnum2 {
			prodNum1, _ := strconv.Atoi(string(tnum1[i]))
			prodNum2, _ := strconv.Atoi(string(tnum2[j]))
			prodNum := prodNum1 * prodNum2
			prod[i+j] += prodNum
			if prod[i+j] > 10 {
				prod[i+j+1] += prod[i+j] / 10
				prod[i+j] = prod[i+j] % 10
			}
		}
	}
	return prod
}

func main() {
	var num1 string = "1"
	var num2 string = "9"
	var tnum1, tnum2 string
	if num1[0] == []byte("-")[0] {
		tnum1 = num1[1:]
	} else {
		tnum1 = num1
	}
	if num2[0] == []byte("-")[0] {
		tnum2 = num2[1:]
	} else {
		tnum2 = num2
	}
	reverse(&tnum1)
	reverse(&tnum2)
	prod := createProduct(tnum1, tnum2)
	if (num1[0] == []byte("-")[0] || num2[0] == []byte("-")[0]) && !(num1[0] == []byte("-")[0] && num2[0] == []byte("-")[0]) {
		fmt.Print("-")
	}
	start := false
	for i := len(prod) - 1; i >= 0; i-- {
		if start {
			fmt.Print(prod[i])
		} else if prod[i] != 0 {
			fmt.Print(prod[i])
			start = true
		}
	}
}
