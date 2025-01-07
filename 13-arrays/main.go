package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
)

func main() {
	var arr1 [5]int                                                   // what is the type of an array? [5]int
	fmt.Println(reflect.TypeOf(arr1), arr1, "len of arr:", len(arr1)) // what is the zero value of the array?
	//for i, _ := range arr1 {
	for i := range arr1 {
		arr1[i] = rand.IntN(100)
	}
	fmt.Println("Sum of arr1", sumOf(arr1))
	arr1[2] = 123
	arr1[0] = 221
	//var arr2 [5]bool // what is the zero value of arr2

	arr2 := [6]int{1, 2, 3, 4, 5, 6} // shorthand declaration
	//arr3 := [...]int{1, 2, 3, 4, 5}  // the compiler evaluates the length
	var arr4 [5]int

	//m := min(len(arr2), len(arr4))

	for i := 0; i < min(len(arr2), len(arr4)); i++ {
		arr4[i] = arr2[i]
	}

	fmt.Println(sumOf(arr4))
	//sumOf()

	var arr5 [5]int

	arr5 = arr1 // since both are same type can be assigned

	fmt.Println(sumOf(arr5))
	// var arr6 [4]int
	// arr6 = arr1

	arr2d := [2][2]int{{1, 2}, {3, 4}}
	arr3d := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	for _, a1 := range arr2d {
		for _, v2 := range a1 {
			fmt.Print(v2, " ")
		}
		println()
	}

	fmt.Println("-- 3d array")
	for _, a1 := range arr3d {
		for _, a2 := range a1 {
			for _, v := range a2 {
				fmt.Print(v, " ")
			}
			println()
		}
	}
}

// array is a fixed size data structure

func sumOf(arr [5]int) int {
	sum := 0
	for _, v := range arr { // for array and slice , the range loop returns index and value
		sum += v
	}
	return sum
}

// range loop is used for collections like array, slice , map and channel
