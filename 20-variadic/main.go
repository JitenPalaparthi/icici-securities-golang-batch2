package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Println("Hello", "World")
	fmt.Println("Hello", "World", "How are you")
	fmt.Println()

	fmt.Println(SumOf())
	fmt.Println(SumOf(10, 20))
	fmt.Println(SumOf(10, 20, 30, 30, 40, 50, 60, 70, 80, 10, 30))
	slice := []int{10, 20, 10, 30, 40, 50, 60, 10, 30}
	fmt.Println(SumOf(slice...))
	arr := [9]int{10, 20, 10, 30, 40, 50, 60, 10, 30}
	//slice2 := arr[:]
	fmt.Println(SumOf(arr[:]...))
}

// variadic params
// the variadic parameter must be the last parameter
// func SumOf(nums ...int, n int)
// variadic can only be used in functions and methods

func SumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func SumOfMul(m int, nums ...int) int {
	//func SumOfMul(nums ...int, m int) int {
	sum := 0
	for _, v := range nums {
		sum += v * m
	}
	return sum
}
