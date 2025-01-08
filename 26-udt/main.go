package main

import "fmt"

func main() {
	r1 := NewRect(10.3, 12.2)
	a1 := Area(r1)
	p1 := Perimeter(r1)
	fmt.Println("Area of r1:", a1)
	fmt.Println("Perimeter of r1:", p1)

	fmt.Println(r1.A, r1.P) // ?
	r2 := NewRect(10.3, 12.2)
	a2 := r2.Area() // no need to call like this
	p2 := r2.Perimeter()
	fmt.Println("Area of r2:", a2)
	fmt.Println("Perimeter of r2:", p2)
	fmt.Println(r2.A, r2.P) // ?
	fmt.Println("shapes slice\n")
	shapes := []any{NewRect(10, 20), NewRect(10, 40), Rect{L: 10.34, B: 34.34}, NewSquare(12.12), 12312.123, "hello wrold", Square{S: 123.123}}
	for _, s := range shapes {

		switch s := s.(type) {
		case *Rect:
			fmt.Println(s.Area())
			fmt.Println(s.Perimeter())
		case Rect:
			fmt.Println(s.Area())
			fmt.Println(s.Perimeter())

		case *Square:
			fmt.Println(s.Area())
			fmt.Println(s.Perimeter())
		case Square:
			fmt.Println(s.Area())
			fmt.Println(s.Perimeter())
		default:
			fmt.Println("cannot perform any operations as undefined type")
		}

	}
}

type Rect struct {
	// L float32
	// B float32
	// A float32
	// P float32
	L, B, A, P float32
}

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func Area(r *Rect) float32 {
	r.A = r.L * r.B
	return r.A
}

func Perimeter(r *Rect) float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

// func (r Rect) Area() float32 {
// 	r.A = r.L * r.B
// 	return r.A
// }

// func (r Rect) Perimeter() float32 {
// 	r.P = 2 * (r.L + r.B)
// 	return r.P
// }

func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

type Square struct {
	S float32
}

func NewSquare(s float32) *Square {
	return &Square{S: s}
}
func (r *Square) Area() float32 {
	return r.S * r.S

}

func (r *Square) Perimeter() float32 {
	return 4 * r.S
}
