package main

import (
	"fmt"
	"runtime"
	"github.com/hehecoi222/go-project/workwithmodule"
)

func main(){
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println("Hello,  世界")
	fmt.Println(string(workwithmodule.CryptPass("Hi")))
	a := 0
	fmt.Scanln(&a)
	fmt.Println(a)
}