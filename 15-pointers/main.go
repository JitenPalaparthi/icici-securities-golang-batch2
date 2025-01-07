package main

import "fmt"

func main() {
	// var num int
	// var arr [5]int
	var ptr *int // what is the zero value of this pointer?
	//*ptr = 100   // what is the error here?

	var num int = 100
	ptr = &num
	println("num:", *ptr)

	var ptr2 *int = new(int) // 8 bytes is allocated and 0 is put there
	fmt.Println(*ptr2)
	*ptr2 = 100
	fmt.Println(*ptr2)

	ptr3 := new(float64)
	*ptr3 = 12312.12312
	// var f2 float64
	// var ptr5 *float64 = &f2

	var ptr4 **int = &ptr
	//var ptr5 ***int = &ptr4
	fmt.Println(**ptr4)

	var arrPtr [5]*int
	fmt.Println(arrPtr)
	//arrPtr[0] = &num
	num = 200
	incr(num)
	println(num) // 200

	incrP(&num)
	println(num) // 200
	// {            // 0x11
	// 	a, b := 10, 20
	// }

	// println(a, b)
}

// nil pointer?
// what deference?
// nil pointer dereference?
// tory hoare --> inventing null pointer

func incr(n int) { // call or pass by value
	n++
}

func incrP(p *int) {
	if p != nil {
		*p++
	}
}

func sq(n int) *int { // dangling pointer
	p := new(int) // creating a new variable , pointer inside a function
	*p = n * n
	return p // giving accesss of the variable outside of the funcion
}
