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
}

type testOut []struct {
	Output int `json:"Output"`   // Change me
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
		result := find_fact(testIn[i].N)
		elapsed := time.Since(start)

		// End func
		if (result == testOut[i].Output) && (elapsed.Seconds() <= 1) {
			fmt.Println("Passed test", i+1, "after:", elapsed)
		}
	}
}
