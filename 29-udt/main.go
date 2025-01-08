package main

import "fmt"

func main() {

	var m1 myint1 = 100
	str1 := m1.ToString()
	s1 := int(myint2(m1).Sq())
	fmt.Println(str1, s1)

	// type casting -- > the same as the underlining type

	var num int = 123

	fmt.Println(myint1(num).ToString())
	fmt.Println(myint2(num).Sq())

	var float1 float32 = 123.23

	fmt.Println(myint1(float1).ToString())
	fmt.Println(myint2(float1).Sq())

	var m3 myint3 = 123

	str3 := myint1(m3).ToString()

	i3 := int(myint2(m3).Sq())

	var i4 int = m3.Cube()

	fmt.Println(str3, i3, i4)

	var f4 float64 = 12.2

	str4 := myint1(f4).ToString()
	i5 := float64(myint2(f4).Sq())
	i6 := float64(myint3(f4).Cube())
	fmt.Println(str4, i5, i6)
}

type myint1 int

func (i myint1) ToString() string {
	return fmt.Sprint(i) // formatter it formats all arguments into a string format
}

type myint2 int

func (i myint2) Sq() myint2 {
	return i * i
}

type myint3 myint1

func (i myint3) Cube() int {
	return int(i * i * i)
}
