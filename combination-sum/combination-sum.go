package main

import (
	"fmt"
	"sort"
)

func main() {
	// const N = 4
	arr := []int{6, 5, 7, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 8, 2, 9, 9, 7, 7, 9}
	const B = 6
	sort.Ints(arr)
	for i := 0; i < len(arr); {
		if i == 0 {
			i++
			continue
		} else if arr[i] == arr[i-1] {
			arr = append(arr[:i], arr[i+1:]...)
		} else {
			i++
		}
	}
	m := make(map[int]int, 100)
	m[0] = 1
	for i := range arr {
		m[arr[i]] = arr[i]
	}
	for i := range arr {
		sum := 0
		var re []int
		check := false
		for sum != B {
			sum += arr[i]
			re = append(re, arr[i])
			if sum > B {
				if (check == false) && (i+1 < len(arr)) && (len(re)-2 > 0) {
					check = true
					re = re[:len(re)-2]
					sum -= arr[i] * 2
					re = append(re, arr[i+1])
					sum += arr[i+1]
				} else {
					re = []int{}
					break
				}
			}
		}
		if len(re) != 0 {
			fmt.Println(re)
		}
		sum = 0
		su := make([]int, len(re))
		for i := 0; i < len(re); i++ {
			su[i] = re[i]
		}
		for j := len(re) - 1; j > 0; j-- {
			i := i + 1
			sum += re[j]
			for ; i < len(arr); i++ {
				if sum >= arr[i] {
					if (m[sum-arr[i]] != 0) && (re[j] != arr[i]) {
						su = su[:j]
						if sum-arr[i] == 0 {
							su = append(su, arr[i])
						} else if sum-arr[i] >= arr[i] {
							su = append(su, arr[i], sum-arr[i])
						} else {
							continue
						}

						fmt.Println(su)
						su = make([]int, len(re))
						for i := 0; i < len(re); i++ {
							su[i] = re[i]
						}
						continue
					}
					if (sum/arr[i] > 2) && (sum%arr[i] == 0) {
						sa := make([]int, sum/arr[i])
						su = su[:j]
						for k := 0; k < sum/arr[i]; k++ {
							sa[k] = arr[i]
						}
						su = append(su, sa...)
						fmt.Println(su)
						su = make([]int, len(re))
						for i := 0; i < len(re); i++ {
							su[i] = re[i]
						}
						continue
					}
				}
			}
		}
	}
}
