package main

import (
	"demo/shape"
	cuboid "demo/shape/cube"
	"demo/shape/rect"
	"demo/shape/square"
	"fmt"
	"math/rand/v2"
)

func main() {
	//var any1 interface{}
	r1 := rect.New(12., 12.)
	s1 := square.New(12.)
	c1 := cuboid.New(12., 12., 12.)
	r2 := rect.New(124., 12.)
	s2 := square.New(13.)
	c2 := cuboid.New(12., 14., 12.)
	r3 := rect.New(124., 11.)
	s3 := square.New(16.)
	c3 := cuboid.New(15., 14., 12.)
	shapes := []shape.IShape{r1, s1, c1, r2, s2, c2, r3, s3, c3}

	// for _, c := range shapes {
	// 	CallShape(c)
	// }

	r := rand.IntN(9)
	CallShape(shapes[r]) // dynamic dispatch --> runtime poly
}

// any interface
// data ptr
// type ptr

// vtables --> virtual tables in c++ and in rust etc./.
// itables --> interface definition table (idt)
// itable is created for  shape.IShape --> Type info and Data info.. all the methods pointers are maintaned
func CallShape(s shape.IShape) { // dynamic dispatch
	fmt.Println("Area of ", s.What(), " is ", s.Area(), " and Perimeter is ", s.Perimeter())
}

func CallRect(s *rect.Rect) { // static dispatch
	fmt.Println("Area of ", s.What(), " is ", s.Area(), " and Perimeter is ", s.Perimeter())
}

// What is an interface?
// An interface is a contract that defines the behavior of an object.
// It is a shared behavior that can be implemented by any type.
// It is an agreement between the object and the interface.
// An interface is a collection of method signatures.
// An interface is a type that defines a set of methods.
// An interface is a type that defines a set of methods that a type must implement.
// It is also called as protocol or trait in other languages.
