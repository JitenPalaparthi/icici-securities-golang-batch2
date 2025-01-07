package main

import "fmt"

func main() {

	var (
		age    uint = 18
		gender rune = 'F'
	)
	//fmt.Println('m')

	cond := age >= 18 && (gender == 'F' || gender == 'f') && true && (false || true)
	//if age >= 18 && (gender == 'F' || gender == 'f') && true && (false || true){
	if cond {
		// true && (true || flase)
		// true && true
		// true
		println("she is eligible for marriage")
	} else if age >= 21 && (gender == 'M' || gender == 109) {
		println("he is eligible for marriage")
	} else {
		fmt.Println("not eligible for marraige")
	}
}
