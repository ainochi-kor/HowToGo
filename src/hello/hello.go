package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Println("Hello, world");
	
	// if 문
	i := 10
	if i >= 5 {
		fmt.Println("5 이상")
	}

	// for 문
	for i := 0; i < 5 ; i++ {
		fmt.Println(i)
	}

	// 변수 사용
	//var age int = 10
	age, name := 10, "Keria"
	// var name string = "Keria"
	fmt.Println(age)
	fmt.Println(name)
	
	num := 10.0
	var num2 float64 = 10.0

	for i :=0; i< 10; i++ {
		num = num - 0.1
		num2 = num2 - 0.1
	}

	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(num == 9.0)
	fmt.Println(num2 == 9.0)

	// 실수 비교 방법
	const epslion = 1e-14
	fmt.Println((math.Abs(num-9.0)) <= epslion)

	// 복소수 : 실수부 + 허수부
	complex_num1 := 1 + 2i
	complex_num2 := 4.2342 + 2.2527i
	complex_num3 := 1.e+3i
	complex_num4 := 7.27151e-10i

	fmt.Println(complex_num1)
	fmt.Println(complex_num2)
	fmt.Println(complex_num3)
	fmt.Println(complex_num4)

	r1 := real(complex_num1)
	i1 := imag(complex_num1)

	fmt.Println(r1)
	fmt.Println(i1)

	// 바이트
	byte_1 := 10
	byte_2 := 0x32
	byte_3 := 'a'
	
	fmt.Println(byte_1)
	fmt.Println(byte_2)
	fmt.Println(byte_3)

	// 룬 : 유니코드(UTF-8) 문자 코드를 저장할 때 사용.
	rune_1 := '한'
	rune_2 := '\ud55c'
	rune_3 := '\ud55c'

	fmt.Println(rune_1)
	fmt.Println(rune_2)
	fmt.Println(rune_3)

	// 연산
	calc_num1 := 3
	calc_num2 := 2

	fmt.Println(calc_num1 + calc_num2)
	fmt.Println(calc_num1 - calc_num2)
	fmt.Println(calc_num1 * calc_num2)
	fmt.Println(calc_num1 / calc_num2)
	fmt.Println(calc_num1 % calc_num2)
	fmt.Println(calc_num1 << calc_num2)
	fmt.Println(calc_num1 >> calc_num2)
	fmt.Println(^calc_num1)

	// Sizeof : 변수 크기

	var sizeof_num1 int8 = 1;
	var sizeof_num2 int16 = 1;
	var sizeof_num3 int32 = 1;
	var sizeof_num4 int64 = 1;

	fmt.Println(unsafe.Sizeof(sizeof_num1))
	fmt.Println(unsafe.Sizeof(sizeof_num2))
	fmt.Println(unsafe.Sizeof(sizeof_num3))
	fmt.Println(unsafe.Sizeof(sizeof_num4))

	// 문자열
	string_1 := "Hello, world! \n"
	string_2 := "안녕하세요	\n"
	string_3 := "\ud55c\uae00"
	string_4 := "\U0000d55c\U0000ae00"
	string_5 := "\xed\x95\x9c\xea\xb8\x80"
	string_6 := 	`안녕하세요 
								Hello world.` // 다중 문자열

	fmt.Println(string_1)
	fmt.Println(string_2)
	fmt.Println(string_3)
	fmt.Println(string_4)
	fmt.Println(string_5)
	fmt.Println(string_6)

	// len: 문자 길이
	fmt.Println(len(string_1))
	fmt.Println(len(string_2))
	fmt.Println(len(string_3))

	// 문자열은 배열이다
	fmt.Printf("%c\n", string_1[1])

	// 열거형
	const (
		Sunday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturady
		numerOfDays
	)

	fmt.Println(Sunday)
}