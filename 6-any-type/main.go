package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	//var any1 interface{} // empty interface
	var any1 any
	if any1 == nil {
		fmt.Println("any1 is nil")
		fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))
	}

	fmt.Println("any1 size", unsafe.Sizeof(any1))
	// what is the zero value of any nil
	any1 = 123123 // int
	//num1 := any1.(float64)      // type assertion
	num1, done := any1.(float64) // trying to assert it
	if done {
		fmt.Println("asserted successfully", num1)
	} else {
		fmt.Println("Could not be asserted", num1)
	}

	//fmt.Println(num1)
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

	any1 = 123123.123123 // float64
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

	any1 = true // bool
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

	any1 = "Hello World" /// string
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

	any1 = uint64(12312) // uint64
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

	any1 = 'A' // int32 or rune
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

}

// cannot do type casting on any /rather use type assertion
// type assertion is used for interfaces since any is also an interface, can use it
// any.(type)
// V.(T) where V is of type interface and T is the type that any has
