package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	_ "time"
)

func main() {

	var (
		num1   uint8   = 123
		num2   uint16  = 1231
		num3   int64   = -12322323
		any1   any     = 12312321.123 // float64
		any2   any     = 1232112312
		float1 float64 = 12312.123
		str1   string  = "12321.123"
		sum    float64
	)

	//sum1, _, _ := Calc(123, 23) // _ blank identifier

	f1, _ := strconv.ParseFloat(str1, 64) // _ blank identifier

	sum = float64(num1) + float64(num2) + float64(num3) + any1.(float64) + float64(any2.(int)) + float1 + f1

	fmt.Printf("%.3f", sum)
	area1(10, 12.3)
	area2(54.5)
	v, err := Add(true, false)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(v)
	}

	v, err = Add(123.123, 884.435)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(v)
	}

	v, err = Add(uint16(786), uint16(884))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(v)
	}

}

func Calc1(a, b int) (int, int, int) { // it should not put ;
	return a + b, a - b, a * b
}

func Add(a any, b any) (any, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return nil, errors.New("both a and b must of of same type")
	}

	_, ok := a.(int) // asserting
	if ok {
		return a.(int) + b.(int), nil
	}
	_, ok = a.(float64)
	if ok {
		return a.(float64) + b.(float64), nil
	}
	// implement it for all numbers
	// implement it for strings (only if you can ParseFloat or ParseInt or Atoi)
	// implement it for bool that means( if true add 1 for false  add 0)

	return nil, fmt.Errorf("\ncannot perform arthimetic operations on a and b because of type:%T", a) // arthe ops can be performed only on numbers
}

func area1(l, b float32) (a float32) {
	a = l * b
	return a
}

func area2(s float32) float32 {
	return s * s
}

// inheritance   -->
// polymorphism  --> static dispatch and dynamic dispatch , compilime time or runtime
// encapsulation --> no special access modifiers or specifiers for encap
// data abstraction -->

// fast compilation

// goc // asm //linker

// .o
// .so
// .a
