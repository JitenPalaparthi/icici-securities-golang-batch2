package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	slice1 := make([]int, 5)
	for i := range slice1 {
		slice1[i] = rand.IntN(100)
	}
	printSliceHeader("slice1", slice1)
	slice1 = append(slice1, 99999, 8888)
	printSliceHeader("slice1", slice1)
	slice1 = append(slice1, 1111, 7777)
	printSliceHeader("slice1", slice1)
	slice2 := slice1 // headers are copied, shallow copy
	printSliceHeader("slice2", slice2)
	slice2[0] = 5555555
	printSliceHeader("slice1", slice1)
	slice2 = append(slice2, 11, 12, 13)
	slice2[1] = 666666
	printSliceHeader("slice1", slice1)
	printSliceHeader("slice2", slice2)
	//slice1 = append(slice1, 100, 200, 300, 400, 500)

}

// array has a fixed length but slice can be extended
// slice contains len and cap
// if slice is only declared but not instantiated then it is nil
// make is used to instantitate a slice, which is going to allcoate memory for the slice
// make is used for slices, maps and channels

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
