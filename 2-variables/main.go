package main

import (
	"fmt" // no need to give a path while using standard package
	"time"
)

func main() {

	fmt.Println("Hello ICICI Securities", time.Now()) // It is afunction from a standard package which fmt

	var num1 int // zero value of all numbers is 0
	fmt.Println("num:", num1)
	num1 = 1000

	var ok1 bool // zero value is false

	var str1 string // zero value of string os ""

	var any1 any // zero value of any is nil

	fmt.Println(ok1, str1, "", "", "", any1)

	var (
		a, b, c  uint8  = 10, 20, 30
		ok2, ok3 bool   = true, false
		str2     string = "Hello"
	)
	fmt.Println(a, b, c, ok2, ok3, str2)

	var num2 = 112312    // based on the value type is inferred, int is given
	var num3 = 12312.123 // based on the value float64 is given
	var ok4 = true       // bool is given

	var age = 39 // uint8
	// statically typed language

	// short hand declaration

	a1 := 100 // shorthand declaration

	str3 := "hello icici" // shorthand declartion

	a1, b1, c1 := 10, 20, 30
	// t1 := b1
	// b1 = a1
	// a1 = t1

	a1, b1 = b1, a1
	a1, b1, c1 = b1, c1, a1

	fmt.Println(ok3, str2, str3, a, b, c, ok4, age, num2, num3, a1, b1, c1)

	char1 := 'A' // rune is alias of int32
	//var char6 rune = 'E'
	var char2 int32 = 'B'
	var char3 uint8 = 'C'
	var char4 uint64 = 'D'
	//var char8 uint8 = 123 // 0- 255
	var char5 int32 = 22000

	fmt.Println("A:", char1)
	fmt.Println("B:", char2)
	fmt.Println("C:", char3)
	fmt.Println("D:", char4)
	fmt.Println("Some", char5)
	var num5 integer = 100
	fmt.Println("num5", num5)

}

type integer = int // integer is an alias for int, in all aspects int and integer are same

// three kinds of packages

// 1. standard package --> GOROOT
// 2. user defined packages --> GOPATH
// 3. third party packages  --> GOPATH

// GOBIN --> if GOBIN is empty, it automaticallys takes the bin directory of GOPATH

/* numbers
uint, uint8,uint16,uint32,uint64
int, int8 , int16, int32, int64
float32, float64
rune,	  byte // are not a separate types, they are just alias for int32 and uint8 respectively
uintptr
*/

/* boolean
bool
*/

/* strings
string
*/

/* interface

any or interface{}

*/

// zero value
// type inference
