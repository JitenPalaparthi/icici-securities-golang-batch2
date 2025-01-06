package main

import "fmt"

var G int = 123123 // data segment

func main() {
	//var n = 100 // when the value of n is assigned?
	const (
		PI        float32 = 3.14 // code segment
		SECOND            = 1
		MIN               = 60 * SECOND
		HOUR              = 60 * MIN
		DAY               = 24 * HOUR
		SOMETHING         = 123 + 123/MIN + HOUR*(10/123.123) + MIN // GO compiler SSA evalues constant expressions
		NUM               = 25
		SQNUM             = NUM * NUM
	)
	num1 := 1312                    // stack or heap
	str1 := "Hello ICICI Securites" // stack or heap
	ok1 := true                     // stack or heap
	fmt.Println(G, PI, num1, str1, ok1, SECOND, MIN, HOUR, DAY)
	// what is an expression and statement?

	a, b, c := 10, 12, 15

	d := a/b*(a*c)-a+b-a+b+c > 100
	fmt.Println(d)
	// operator precedence

	//PI = PI + 100
}
