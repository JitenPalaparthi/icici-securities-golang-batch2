package main

import "fmt"

var G int = 3

func main() {
	num := 2
	r1 := func() int {
		num = num * num * G
		return num
	}()
	println(r1)

	r2 := func(n int) int {
		n = n * n * G
		return n * n * G
	}(num)
	println(r2)
	println(num)

	for i := 1; i <= 10; i++ {
		func(j1 int) {
			println(j1)
		}(i)
	}
	j := 1
loop:
	func(j1 int) {
		println(j1)
	}(j)
	j++
	if j < 11 {
		goto loop
	}

	a, b := 12, 20

	r3 := calc(a, b, func(i1, i2 int) int {
		return i1 + i2
	})
	println(r3)
	r4 := calc(a, b, sub)
	println(r4)

	fn1 := func(a, b int) int {
		return a * b
	}

	r5 := calc(a, b, fn1)
	println(r5)

	var struct1 struct {
		a, b  int
		addfn func(int, int) int
		empty struct{} // empty struct
	} = struct {
		a     int
		b     int
		addfn func(int, int) int
		empty struct{}
	}{
		a: 100,
		b: 200,
		addfn: func(i1 int, i2 int) int {
			return i1 + i2
		},
		empty: struct{}{}, // value of a empty structure
	}

	r6 := struct1.addfn(struct1.a, struct1.b)
	fmt.Println(r6)
	fmt.Println(struct1.empty)

}

func calc(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func sub(a, b int) int {
	return a - b
}

// type FnType func(any,any)any
// calc(a,b any, sn FnType )any
