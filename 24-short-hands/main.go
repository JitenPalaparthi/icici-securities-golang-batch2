package main

import "fmt"

func main() {
	const PI = 3.14
	var num1 = 123.23 //
	num2 := 100       // shorthand of int
	c1 := 12.4 + 1.1i // complex number shorthand
	ptr1 := new(int)
	ptr2 := &num2
	arr := [3]int{1, 2, 3}
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 5)
	map1 := map[string]any{"1": 1}
	map2 := make(map[string]any)
	fmt.Println(num1, num2, c1, ptr1, ptr2, arr, slice1, slice2, map1, map2)
}
