package main

func main() {
	//p := sq(100)
	var arr1 [99231311]int
	var str1 string = "hello icici securities"
	//fmt.Println(str1) // it is very expensive stores in heap memory
	for _, v := range arr1 {
		print(v)
	}

	println()
	for _, v := range str1 {
		print(string(v))
	}

	num := 100
	incr(num)
	println(num)
	incrP(&num)
	println(num)

	p := new(int)
	sqP(num, p)
	println(*p)

	p1 := sq(num)
	println(*p1)
}

func incr(n int) { // call or pass by value
	n++
}

func incrP(p *int) {
	if p != nil {
		*p++
	}
}

func sqP(num int, s *int) {
	if s != nil {
		*s = num * num
	}
}

func sq(n int) *int { // dangling pointer
	p := new(int) // creating a new variable , pointer inside a function
	*p = n * n
	return p // giving accesss of the variable outside of the funcion
}
