package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var num1 uint8 = 100
	//var num2 int = num1 // there is no implicit type casting in go
	var num2 int = int(num1) // type cast it
	fmt.Println("num1:", reflect.TypeOf(num1))
	fmt.Println("num2:", reflect.TypeOf(num2))
	fmt.Printf("num1:%T", num1)
	// all numbers can be type casted agaist each other

	var num3 float32 = 1231.123

	var num4 int = int(num3)
	fmt.Println("num4:", num4)

	//var ok1 bool = true // 1 byte

	//num1 = uint8(ok1) // 1 byte , type casting is not possible here

	char1 := 'A'
	char2 := 22000

	fmt.Println(string(char1))
	fmt.Println(string(char2))

	var (
		num11  uint8   = 100
		num12  uint16  = 200
		num13  uint32  = 300
		num14  uint64  = 12312
		num5   int     = -12
		num6   int8    = -123
		num7   int32   = 123123
		num8   int64   = 12312313123123
		float1 float32 = 12312.12312
		float2 float64 = 1231234343.234
	)

	var sum float64 = float64(num11) + float64(num12) + float64(num13) + float64(num14) + float64(num5) + float64(num6) + float64(num7) + float64(num8) + float64(float1) + float2
	fmt.Printf("sum:%.3f\n", sum)

	add, sub, mul, div := Calc(10, 20)
	fmt.Println(add, sub, mul, div)

	var str1 string = "123123"
	n, err := strconv.Atoi(str1) // err is an interface

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Converted from str to int", n)
	}

	var str2 string = "1a23123"
	n, err = strconv.Atoi(str2) // err is an interface

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Converted from str to int", n)
	}

	var str3 string = "12312.123"

	f1, err := strconv.ParseFloat(str3, 32)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value converted from str3 to flaot64", f1)
	}

	// str4 := "A"
	// b1, err := strconv.ParseBool(str4)
	// fmt.Println(b1)

	c1 := complex(123.123, 1.23) // complex128
	c2 := 123.123 + 2.3i         // it is taken as a complex number
	var (
		r1, im1 float32 = 123.123, 1.23
	)
	c3 := complex(r1, im1)
	c4 := c1 + c2
	c5 := c1 - c2
	c6 := c1 * c2

	c7 := complex64(c1) + c3 // c1 is complex128, and c3 is complex64

	r2 := real(c1)
	im2 := imag(c1)

	fmt.Println(c1, c2, c3, c4, c5, c6, c7, r2, im2)
}

func Calc(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}
