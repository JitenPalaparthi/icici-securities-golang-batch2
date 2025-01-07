package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	slice1 := make([]int, 5, 10)
	for i := range slice1 {
		slice1[i] = rand.IntN(100)
	}
	printSliceHeader("slice1", slice1)
	slice1 = AddNSq(slice1, 11, 12, 13)
	printSliceHeader("slice1", slice1)
	AddNSqP(&slice1, 10, 3, 4, 5, 6, 7, 77)
	printSliceHeader("slice1", slice1)
}

func printSliceHeader(name string, slice []int) {
	fmt.Println(name)
	fmt.Println("Values:", slice)
	if slice != nil {
		fmt.Printf("Address of Slice:%p \nLen: %d\nCap: %d\n", &slice, len(slice), cap(slice))
	}
	if len(slice) > 0 {
		fmt.Println("Pointer:", &slice[0])
	}
	fmt.Println("------------------------")
}

// slice 1 2 3 4
// AddNSq(slice,5,6,7)
func AddNSq(slice []int, nums ...int) []int {
	slice = append(slice, nums...)
	for i, v := range slice {
		slice[i] = v * v
	}
	return slice
}

func AddNSqP(slice *[]int, nums ...int) {
	*slice = append(*slice, nums...)
	for i, v := range *slice {
		(*slice)[i] = v * v
	}
}
