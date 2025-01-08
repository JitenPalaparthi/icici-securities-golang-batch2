package main

import (
	"fmt"

	"github.com/JitenPalaparthi/icici-2-shapes/shape"
	cuboid "github.com/JitenPalaparthi/icici-2-shapes/shape/cube"
	"github.com/JitenPalaparthi/icici-2-shapes/shape/rect"
)

func main() {
	r1 := rect.New(12.2, 6.45)
	a1 := r1.Area()
	p1 := r1.Perimeter()

	fmt.Println("area of r1 a1:", a1)
	fmt.Println("perimeter of r1 p1:", p1)
	//v := rand.IntN(100)
	shape.What()

	c1 := cuboid.New(12.2, 12.4, 1.2)
	a2 := c1.Area()
	p2 := c1.Perimeter()

	fmt.Println("area of c1 a2:", a2)
	fmt.Println("perimeter of c1 p2:", p2)
	shape.What()
	//shape.whatIsthis() cannot call here since it is a camel case which is unexpoeted
	var t shape.T
	t.A3 = 100
	// t.a1 =12
	// t.a2 =123
}
