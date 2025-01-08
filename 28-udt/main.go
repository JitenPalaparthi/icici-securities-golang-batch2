package main

import "fmt"

func main() {

	p1 := NewPerson("jiten", "jitenP@outlook.com", "522001", "Guntur")
	//fmt.Println(p1)
	p1.Display()
	//var p2 string

	var a1 struct {
		PinCode, City string
	}
	a1.City = "Guntur"
	a1.PinCode = "522001"

	fmt.Println(a1)
	var a2 struct {
		PinCode, City string
	}

	a2.City = "Guntur"
	a2.PinCode = "522001"
	fmt.Println(a2)

	var s1 struct{} // there is no size to the structure
	fmt.Println(s1)

	e1 := Empty{}
	e1.SayHi()

}

func NewPerson(n, e, p, c string) *Person {
	return &Person{Name: n, Email: e, Address: struct {
		PinCode string
		City    string
	}{PinCode: p, City: c}}
}

type Person struct {
	Name, Email string
	Address     struct { // embedded struct
		PinCode, City string
	}
}

func (p *Person) Display() {
	fmt.Printf("Name:%s\nEmail:%s\nPincode:%s\nCity:%s\n", p.Name, p.Email, p.Address.PinCode, p.Address.City)
}

type Empty struct{} // this can be used for signalling purpose

func (e *Empty) SayHi() {
	fmt.Println("hey you")
}
