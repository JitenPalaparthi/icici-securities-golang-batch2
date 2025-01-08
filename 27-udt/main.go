package main

import "fmt"

func main() {

	p1 := Person{Name: "Jiten", Email: "Jitenp@outlook.com", Address: &Address{PinCode: "560086", City: "Bangalore"}}
	p1.Display()
	p1.Address.Display()
	c1 := Company{Name: "Jiten", Website: "aka-labs.in", Status: "operational", Address: &Address{PinCode: "560086", City: "Bangalore", Status: "active"}}

	c1.Display()
	c1.Address.Display()
	c1.IsValidPincode()

}

type Person struct {
	Name, Email string
	Address     *Address // it is okay to use the same name as type as field
} //

type Address struct {
	PinCode, City, Status string
}

func (a *Address) Display() {
	fmt.Printf("Pincode:%v \nCity:%v\n", a.PinCode, a.City)
}

func (a *Address) IsValidPincode() bool {
	return true
}

func (a *Person) Display() {
	fmt.Printf("Name:%v \nEmail:%v\n", a.Name, a.Email)
}

type Company struct {
	Name     string
	Website  string
	Status   string
	*Address // promoted field , may be similar to inheritance but in the form of composition
}

func (a *Company) Display() {
	fmt.Printf("Name:%v \nWebsite:%v\n", a.Name, a.Website)
}
