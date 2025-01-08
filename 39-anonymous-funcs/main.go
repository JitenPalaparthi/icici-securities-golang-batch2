package main

import "fmt"

type fnType func(int, int) int

func (f fnType) Sq(n int) int {
	return n * n
}

func main() {
	var fn1 func(int, int) int
	func() {
		println("hello ICICI Securites")
	}() // executed

	r1 := func(a, b int) int {
		return a + b
	}(10, 20)

	f1 := func(a, b int) int {
		return a - b
	}
	r2 := f1(10, 20)

	fmt.Println(r1)
	fmt.Println(r2)

	fmt.Println("--------------")

	map1 := make(map[string]any)
	map1["add"] = func(a, b int) int {
		return a + b
	}
	fn1 = func(i1, i2 int) int {
		return i1 - i2
	}
	map1["sub"] = fn1
	map1["mul"] = mul

	var fn2 fnType = func(i1, i2 int) int {
		return i1 / i2
	}
	map1["div"] = fn2

	map1["pi"] = 3.14
	map1["greet"] = func() {
		fmt.Println("Hello ICICI")
	}

	a, b := 100, 200
	for k, v := range map1 {
		switch v.(type) {
		case float64:
			println("Some float value for the key", k, "is", v.(float64))
		case func():
			fmt.Println("Executing a function without any input params and output params")
			v.(func())()
		case func(int, int) int:
			fmt.Println("Executing a function without any input params and output params")
			r := v.(func(int, int) int)(a, b)
			fmt.Println(r)
		case fnType:
			r := v.(fnType)(a, b)
			fmt.Println("Functype:", r)
		default:
			println("not a defined type")
		}

	}

}

func mul(a, b int) int {
	return a * b
}
