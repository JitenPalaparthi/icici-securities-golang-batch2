package main

import (
	"fmt"
	"time"
)

func main() { // parent
	go greet() // child
	go gennums()
	fmt.Println("Hello Main")
	time.Sleep(time.Millisecond * 1) // not a solution
}

func greet() {
	fmt.Println("Hello ICICI")
}
func gennums() {
	go func() {
		for i := 1; i <= 100; i++ {
			go fmt.Println(i) //
		}
	}()
}

// Thread pool --> 10
// connection pool -->
// L7 L4 L2
//

// main is also a go routine
// no goroutine waits for other goroutines to completes it execution
// order of execution is not guaranteed
