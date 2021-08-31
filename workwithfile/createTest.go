package main

import (
	"encoding/json"
	"os"
)

type testIn []struct {
	N int `json:"N"`   //Change me
	Arr []int `json:"Arr"`
}

type testOut []struct {
	Output1 int `json:"Output1"`  //Change me
	Output2 int `json:"Output2"`
}

// Put func here
func EvenOddSum(N int, Arr []int) (int,int)  {
	EvenOddSum := []int{0,0}
	for i:= range Arr {
		if (i+1) % 2 ==0 {
			EvenOddSum[1] += Arr[i]
		} else {
			EvenOddSum[0] += Arr[i]
		}
	}
	return EvenOddSum[0],EvenOddSum[1]
}
// End func

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func workFile() testIn {
	fi, e := os.Open("testIn.json")
	check(e)
	defer fi.Close()
	
	var testIn testIn

	Decod := json.NewDecoder(fi)
	Decod.Decode(&testIn)
	return testIn
}

func writeFile(testOut testOut) {
	fo, e := os.Create("testOut.json")
	check(e)
	defer fo.Close()
	Encod := json.NewEncoder(fo)
	Encod.Encode(testOut)
}

func main()  {
	testIn := workFile()
	testOut:= make(testOut,len(testIn)) 
	for i := range testIn {
		testOut[i].Output1,testOut[i].Output2 = EvenOddSum(testIn[i].N,testIn[i].Arr)
	}
	writeFile(testOut)
}