package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type testIn []struct {
	N int `json:"N"`   // Change me
	Arr []int `json:"Arr"`
}

type testOut []struct {
	Output1 int `json:"Output1"`   // Change me
	Output2 int `json:"Output2"`
}

func ReadIn() (testIn, testOut) {
	fi, e := os.Open("testIn.json")
	check(e)
	fo, e := os.Open("testOut.json")
	check(e)
	defer fi.Close()
	defer fo.Close()

	var testIn testIn
	var testOut testOut

	Decod := json.NewDecoder(fi)
	Decod.Decode(&testIn)

	Decod = json.NewDecoder(fo)
	Decod.Decode(&testOut)

	return testIn, testOut
}

// Put func here




// End func

func main() {
	testIn, testOut := ReadIn()

	for i := range testIn {
		//Put func here

		start := time.Now()
		result1,result2 := EvenOddSum(testIn[i].N,testIn[i].Arr)
		elapsed := time.Since(start)

		// End func
		if (result1== testOut[i].Output1) && (result2 ==testOut[i].Output2) && (elapsed.Seconds() <= 1) {
			fmt.Println("Passed test", i+1, "after:", elapsed)
		}
	}
}
