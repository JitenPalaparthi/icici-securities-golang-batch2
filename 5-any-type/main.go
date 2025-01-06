package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	str1 := "Hello World"
	str2 := "Hello ICICI Securities"
	//str4 := str1 + str2
	//strings.Builder
	var str3 string
	/*
				str1
				---
				string header
				ptr -- 8 bytes
				len -- 8 bytes

				any1
		         ----
				 interface header
				 dataPtr --> pointer to the actual data 8 bytes
				 typePtr --> pointer of the type        8 bytes
	*/

	fmt.Println("Size of str1:", unsafe.Sizeof(str1))
	fmt.Println("Size of str2:", unsafe.Sizeof(str2))
	fmt.Println("Size of str3:", unsafe.Sizeof(str3))

	//var any1 interface{} // empty interface
	var any1 any
	if any1 == nil {
		fmt.Println("any1 is nil")
		fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))
	}

	fmt.Println("any1 size", unsafe.Sizeof(any1))
	// what is the zero value of any nil
	any1 = 123123 // int
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

	any1 = 123123.123123 // float64
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

	any1 = true // bool
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

	any1 = str1 /// string
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

	any1 = uint64(12312) // uint64
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

	any1 = 'A' // int32 or rune
	fmt.Println("Value of any1:", any1, "Type of any1:", reflect.TypeOf(any1))

	str1 = "hello Again world" // mutable

}
