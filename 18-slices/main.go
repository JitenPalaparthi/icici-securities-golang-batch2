package main

import (
	"fmt"
	"math/rand/v2"
	"unsafe"
)

func main() {
	var slice1 []int // define a slice
	fmt.Println("Size of slice1:", unsafe.Sizeof(slice1))
	if slice1 == nil {
		fmt.Println("yes nil slice")
	}
	slice1 = make([]int, 5)
	fmt.Println(slice1)

	for i := range slice1 {
		slice1[i] = rand.IntN(100)
	}
	//slice1[0] = 999999

	fmt.Println(slice1)

	slice2 := make([]int, 3)
	slice3 := []int{10, 12, 14}

	fmt.Println(sumOf(slice1))
	fmt.Println(sumOf(slice2))
	fmt.Println(sumOf(slice3))
	printSliceHeader("slice1", slice1)
	printSliceHeader("slice2", slice2)
	printSliceHeader("slice3", slice3)

}

// array has a fixed length but slice can be extended
// slice contains len and cap
// if slice is only declared but not instantiated then it is nil
// make is used to instantitate a slice, which is going to allcoate memory for the slice
// make is used for slices, maps and channels

func sumOf(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func printSliceHeader(name string, slice []int) {
	fmt.Println(name)
	if slice != nil {
		fmt.Printf("Address of Slice:%p \nLen: %d\nCap: %d\n", &slice, len(slice), cap(slice))
	}
	if len(slice) > 0 {
		fmt.Println("Pointer:", &slice[0])
	}
	fmt.Println("------------------------")
}
