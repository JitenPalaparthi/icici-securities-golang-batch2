package main

import "fmt"

func main() {
	//num := 0
	//slice := []int{1, 2, 3}
	var ptr *int
	func() { // func1
		defer recoverme()
		println("start of func1")
		func() { // func2
			defer recoverme()
			println("start of func2")

			func() { //func3
				defer recoverme()
				println("start of func3")
				// for i := 0; i <= len(slice); i++ {
				// 	println(slice[i])
				// }
				//println(100 / num) // this error is caught by what?
				*ptr = 100
				println("end of func3")
			}()
			println("end of func2")
		}()
		println("end of func1")
	}()
	println("end of main")
}

// error --> can handle errors
// panic --> can recover from a panic
// fatal --> cannot handle or recover becase it is critical issues and the process gets terminated

// built in function called recover

func recoverme() {
	if r := recover(); r != nil { // if r is not nil, there is panic
		fmt.Println("recovering from panic", r)
	}
}
