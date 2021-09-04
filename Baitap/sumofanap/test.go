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
	N int `json:"N"` //Change me
	A int `json:"A"`
	D int `json:"D"`
}

type testOut []struct {
	Output int `json:"Output"`  //Change me
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

	json.NewDecoder(fi).Decode(&testIn)

	json.NewDecoder(fo).Decode(&testOut)

	return testIn, testOut
}

// Put func here





// End func

func main() {
	testIn, testOut := ReadIn()

	for i := range testIn {
		//Put func here

		start := time.Now()
		result := sum_of_ap(testIn[i].N,testIn[i].A,testIn[i].D)        // Change me
		elapsed := time.Since(start)

		// End func
		if (result == testOut[i].Output) && (elapsed.Seconds() <= 1) {   // Change me
			fmt.Println("Passed test", i+1, "after:", elapsed)
		}
	}
}
