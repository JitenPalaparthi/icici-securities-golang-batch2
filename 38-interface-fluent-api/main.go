package main

import "fmt"

func main() {
	c1 := New(100)
	v := c1.Add(100).Add(20).Sub(100).Mul(10).Div(10).Get()
	//fmt.Println(v)
	println(v)
	//Query("Select * from something").Where("id=1").Filter().Order()

	var any1 any = 100
	num := any1.(int)
	fmt.Println(num)

	var e1 IEmpty = 100
	num1 := e1.(int)
	fmt.Println(num1)

	var e2 IEmpty = struct {
		Num   int
		Email string
	}{Num: 101, Email: "JitenP@outlook.com"}

	fmt.Println(e2)

}

type IEmpty interface{} // data pointer and type pointer

func New(d int) *Calc {
	return &Calc{data: d}
}

type Calc struct {
	data int
}

func (c *Calc) Add(n int) ICalc {
	c.data += n
	return c
}

func (c *Calc) Sub(n int) ICalc {
	c.data -= n
	return c
}
func (c *Calc) Mul(n int) ICalc {
	c.data *= n
	return c
}
func (c *Calc) Div(n int) ICalc {
	c.data /= n
	return c
}
func (c *Calc) Get() int {
	return c.data
}

type ICalc interface {
	Add(int) ICalc
	Sub(int) ICalc
	Mul(int) ICalc
	Div(int) ICalc
	Get() int
}
