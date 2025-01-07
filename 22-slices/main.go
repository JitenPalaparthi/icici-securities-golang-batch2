package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSliceHeader("slice1", slice1)
	slice2 := slice1
	printSliceHeader("slice2", slice2)
	slice3 := slice1[:]
	printSliceHeader("slice3", slice3)

	slice4 := slice1[:5] // 0-4 but not 5
	printSliceHeader("slice4", slice4)
	slice5 := slice1[2:8]
	//slice5[0] = 11111
	printSliceHeader("slice5", slice5)
	slice6 := slice1[5:]
	printSliceHeader("slice6", slice6)

	slice7 := append(slice1[:4], slice1[5:]...)
	fmt.Println(slice7)
	//clear(slice7)
	fmt.Println(slice7)
	slice8 := make([]int, 4)
	copy(slice8, slice1)
	fmt.Println(slice8)
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
