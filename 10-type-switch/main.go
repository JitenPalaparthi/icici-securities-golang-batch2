package main

import (
	"errors"
	"fmt"
	"reflect"
)

// complile time -->
func main() {
	if s1, err := add(100, 200); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("sum:", s1)
	}
	//fmt.Println(s1)

	s2 := addG(10, 20) // func addGInt(a int, b int)(int)
	fmt.Println("Sum", s2)

	s3 := addG(120.23, 14.45) //func addGFloat(a flaot64, b flaot64)(flaot64)
	fmt.Println("Sum", s3)

	if s4, err := add(uint8(10), uint8(20)); err != nil { // vtables , interface definition table (idt)
		fmt.Println(err.Error())
	} else {
		fmt.Println("sum:", s4)
	}

	s5 := addG(uint8(10), uint8(20)) // func addGuint8(a uint8, b uint8)(uint8)
	fmt.Println("sum:", s5)
	//sb1 := subG(10, 20)     // subGint(a int, b int)int{return a-b}
	//sb2 := subG(10.3, 20.0) // subGfloat64(a float64, b flaot64)float64{return a-b}
	//sb3 := subG(int8(100), float32(100))
}

func add(a, b any) (any, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return nil, errors.New("both a and b are different types")
	}
	if !isNumber(a) {
		return nil, errors.New("a and b are not number types")
	}
	var sum any

	switch a.(type) { // type switch
	case nil:
		return nil, errors.New("a and b is nil type")
	case int:
		sum = a.(int) + b.(int)
	case uint:
		sum = a.(uint) + b.(uint)
	case uint8:
		sum = a.(uint8) + b.(uint8)
	case uint16:
		sum = a.(uint16) + b.(uint16)
	case uint32:
		sum = a.(uint32) + b.(uint32)
	case uint64:
		sum = a.(uint64) + b.(uint64)
	case int8:
		sum = a.(int8) + b.(int8)
	case int16:
		sum = a.(int16) + b.(int16)
	case int32:
		sum = a.(int32) + b.(int32)
	case int64:
		sum = a.(int64) + b.(int64)
	case float32:
		sum = a.(float32) + b.(float32)
	case float64:
		sum = a.(float64) + b.(float64)
	}
	return sum, nil
}
func addG[T IType](a T, b T) T {
	return a + b
}

func subG[T IType](a T, b T) T {
	return a - b
}

func isNumber(a any) bool {
	switch a.(type) {
	case int, uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		return true
	default:
		return false
	}
}

type IType interface {
	int | uint | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

type IType2 interface {
	int | uint | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}

// vtables
//
