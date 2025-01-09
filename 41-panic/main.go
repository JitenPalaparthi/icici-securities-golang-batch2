package main

func main() {
	//num := 0
	//slice := []int{1, 2, 3}
	var ptr *int
	func() { // func1
		println("start of func1")
		func() { // func2
			println("start of func2")

			func() { //func3
				println("start of func3")

				// for i := 0; i <= len(slice); i++ {
				// 	println(slice[i])
				// }
				//println(100 / num) // this error is caught by what?
				*ptr = 100
				println("end of func2")
			}()
			println("end of func2")
		}()
		println("end of func1")
	}()
	println("end of main")
}

// error
// panic
// Fatal
