package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

func main() {
	fmt.Println("CPI Count : ", runtime.NumCPU())

	i := 10
	if i > 5 {
		fmt.Println("5 초과")
	} else if i == 5{
		fmt.Println("5")
	} else {
		fmt.Println("5 미만")
	}

	var b []byte
	var err error

	b, err = ioutil.ReadFile("./hello.txt")

	if err == nil {
		fmt.Printf("%s", b)
	}
}