package main

func main() {
	defer println("end of main-1")
	defer println("end of main-2")
	defer println("end of main-3")

	str := "Hello ICICI Securites"
	defer func() {
		for _, c := range str {
			defer print(string(c))
		}
	}()

	num := 100

	defer func(n int) {
		println("defrfered:", n)
	}(num)

	num += 1
	println("norma:", num)

	println("start of main")
}
