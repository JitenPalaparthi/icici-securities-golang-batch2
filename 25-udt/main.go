package main

import "fmt"

func main() {
	var p1 Person // it is not a pointer
	p1.Name = "Jiten"
	p1.Email = "JitenP@outlook.com"

	p2 := New("Priya", "P@gmail.com")
	fmt.Println(p2)

	p3 := new(Person) // builtin function new
	p3.Name = "JP"
	p3.Email = "jp@gmail.com"

	p4 := Person{Name: "Karthik", Email: "karthik@gmail.com"}
	fmt.Println(p4)

	p5 := &Person{Name: "Karthik", Email: "karthik@gmail.com"}
	fmt.Println(p5)

	cc1 := NewColorCode("red", 123, 123)
	fmt.Println(cc1)

}

// user defined types
type integer = int // this is not a new type it is just an alias

type Person struct { // this is creating a new type
	Name  string
	Email string
}

func New(n, e string) *Person {
	return &Person{Name: n, Email: e}
}

type ColorCode struct { // anonymous fields of a struct
	string
	int
	integer
}

func NewColorCode(colour string, code int, acode int) *ColorCode {
	return &ColorCode{string: colour, int: code, integer: acode}
}
