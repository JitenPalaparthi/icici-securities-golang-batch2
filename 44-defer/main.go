package main

func main() {
	defer println("end of main") // defers the execution to the end of the caller
	func() {                     //func1
		defer println("end of func-1") // defers the execution to the end of the caller
		func() {                       //func2
			defer println("end of func-2") // defers the execution to the end of the caller
			panic("Something went wrong")  // before going to panic all the defers in that callstack are called
		}()
	}()

	defer func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				println(i)
			}
		}
	}()
	println("start of main")

	// func() { //func2
	// 	defer println("end of funcw") // defers the execution to the end of the caller
	// 	panic("Something went wrong") // before going to panic all the defers in that callstack are called
	// }()
}
