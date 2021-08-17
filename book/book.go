package main

import (
	"fmt"
	"math"
)

var n = 3

func check(hor int, ver int, input [][]int) bool {
	if input[hor][ver] == 0 {
		return false
	}
	var horup,hordo,verle,verri bool = true,true,true,true
	for i := 0; i < hor; i++ {
		if input[i][ver] > input[hor][ver] {
			horup = false
			break
		}
	}
	for i := 0; i < n-hor-1; i++{
		if input[n-i-1][ver] > input[hor][ver] {
			hordo = false
			break
		}
	}
	for i := 0; i < ver; i++ {
		if input[hor][i] > input[hor][ver] {
			verle = false
			break
		}
	}
	for i := 0; i < n-ver-1; i++ {
		if input[hor][n-1-i] > input[hor][ver] {
			verri = false
			break
		}
	}
	if !(horup) && !(hordo) && !(verle) && !(verri){
		return false
	} else {
		return true
	}
}

func main() {
	var input = [][]int{
		[]int{3, 3, 2},
		[]int{4, 1, 3},
		[]int{3, 2, 0},
	}
	possible := math.Pow(float64(n), 2)
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if !(check(i, j, input)) {
				possible--
			}
		}
	}
	fmt.Println(possible)
}
