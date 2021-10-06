package main

import (
	"fmt"
	"strconv"
)

func reverse(tnumIn string) string {
	tnum := []byte(tnumIn)
	for i := 0; i < len(tnum)/2; i++ {
		tnum[i], tnum[len(tnum)-i-1] = tnum[len(tnum)-i-1], tnum[i]
	}
	return string(tnum)
}

func main() {
	var num1 string = "999999"
	var num2 string = "99"
	var tnum1, tnum2 string
	if num1[0] == []byte("-")[0] {
		tnum1 = num1[1:]
	} else if num2[0] == []byte("-")[0] {
		tnum2 = num2[1:]
	} else {
		tnum1 = num1
		tnum2 = num2
	}
	tnum1 = reverse(tnum1)
	tnum2 = reverse(tnum2)
	tpro := make([]int, len(tnum1)+len(tnum2))
	for i, v := range tnum2 {
		for j, vnum := range tnum1 {
			n1, _ := strconv.Atoi(string(v))
			n2, _ := strconv.Atoi(string(vnum))
			tpro[i+j] += n1*n2
		}
	}
	for i, v := range tpro {
		if v >= 10 {
			tpro[i+1] += tpro[i] / 10
			tpro[i] = tpro[i] % 10
		}
	}
	for i := len(tpro) - 1; i >= 0; i-- {
		fmt.Print(tpro[i])
	}
}
