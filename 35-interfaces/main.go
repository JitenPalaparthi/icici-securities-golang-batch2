package main

import (
	"demo/shape"
	cuboid "demo/shape/cube"
	"demo/shape/rect"
	"demo/shape/square"
	"fmt"
)

func main() {
	//var any1 interface{}
	r1 := rect.New(12., 12.)
	s1 := square.New(12.)
	c1 := cuboid.New(12., 12., 12.)

	CallShape(r1)
	CallShape(s1)
	CallShape(c1)

}

func CallShape(s shape.IShape) {
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
