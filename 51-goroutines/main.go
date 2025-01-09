package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() { // parent
	go greet() // child
	go gennums()
	fmt.Println("Hello Main")
	go func() {
		i := 1
		for {
			time.Sleep(time.Second * 1)
			println(i)
			i++
			if i == 11 {
				//runtime.Goexit()
				break
			}
		}
	}()
	runtime.Goexit() // What goexit does? how does it behave?
	// it crashes main , it makes sure that all other goroutines complete their execution
	//os.Exit(0)       // graceful
	//return           // graceful
}

func greet() {
	fmt.Println("Hello ICICI")
}
func gennums() {
	go func() {
		for i := 1; i <= 100; i++ {
			go fmt.Println(i) //
		}
		runtime.Goexit() // normal
	}()
}

// Thread pool --> 10
// connection pool -->
// L7 L4 L2
//

// main is also a go routine
// no goroutine waits for other goroutines to completes it execution
// order of execution is not guaranteed
