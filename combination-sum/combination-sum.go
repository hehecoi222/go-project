package main

import (
	"fmt"
	"sort"
)

var arr []int = []int{7,2,5,6}

const B = 16

func removeDuplicate(arr []int) []int {
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
	return arr
}

func copyTo(m interface{}, arr []int) interface{} {

	switch m.(type) {

	case map[int]int:
		for i := range arr {
			m.(map[int]int)[arr[i]] = arr[i]
		}
		return m.(map[int]int)
	case []int:
		for i := 0; i < len(arr); i++ {
			m.([]int)[i] = arr[i]
		}
		return m.([]int)
	}
	return 0
}

func formResult(arr []int, i int) []int {
	sum := 0
	check := false
	var re []int
	for sum != B {
		sum += arr[i]
		re = append(re, arr[i])
		if sum > B {
			if !(check) && (i+1 < len(arr)) && (len(re)-2 > 0) {
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
	return re
}

func main() {
	// const N = 4
	sort.Ints(arr)
	arr = removeDuplicate(arr)

	m := make(map[int]int, 100)
	m[0] = 1
	m = copyTo(m, arr).(map[int]int)

	for i := range arr {
		re := formResult(arr, i)
		sum := 0
		su := make([]int, len(re))
		su = copyTo(su,re).([]int)

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
						su = copyTo(su,re).([]int)
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
						su = copyTo(su,re).([]int)
						continue
					}
				}
			}
		}
	}
}
