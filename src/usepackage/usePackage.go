package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

func main() {
	fmt.Println("CPU Count : ", runtime.NumCPU())

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

	for i := 0; i < 5 ; i++ {
		fmt.Println(i)

		switch i {
		case 0 :
			fmt.Println(0)
		case 1: 
			fmt.Println(1)
		case 2: 
			fmt.Println(2)
			fallthrough // 아래 case 까지 처리
		case 3: 
			fmt.Println(3)
		case 4: 
			fmt.Println(4)
		default:
			fmt.Println(5)
		}

		switch i {
		case 0,2,4,6:
			fmt.Println("짝수")
		case 1,3,5:
			fmt.Println("홀수")
		}
	}
}