package main

import (
	"fmt"
	v2 "math/rand/v2"
	"sync"
	"time"
)

var G int = func() int { // this gets executed even before main..
	return v2.IntN(100)
}()

var wg *sync.WaitGroup

func init() {
	wg = new(sync.WaitGroup)
	println("Hello calling init1", G) // main is not the starting point
}

// func init() {
// 	println("Hello calling init2")
// }

// func init() {
// 	println("Hello calling init3")
// }

// func init() {
// 	println("Hello calling init4")
// }

func main() { // just assume parent
	wg.Add(1)
	go greet() // child
	wg.Add(1)
	go func() {
		gennums()
		wg.Done()
	}()
	fmt.Println("Hello Main")
	wg.Add(1)
	go func() {
		i := 1
		for {
			time.Sleep(time.Microsecond * 10)
			println(i)
			i++
			if i == 11 {
				//runtime.Goexit()
				break
			}
		}
		wg.Done()
	}()
	wg.Wait() //115  the state/counter becomes zero
	// it crashes main , it makes sure that all other goroutines complete their execution
	//os.Exit(0)       // graceful
	//return           // graceful
}

func greet() {
	fmt.Println("Hello ICICI")
	wg.Done()
}
func gennums() {
	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go func() {
				fmt.Println(i) //
				wg.Done()
			}()
		}
		wg.Done()
		//runtime.Goexit() // normal
	}()
}

// Thread pool --> 10
// connection pool -->
// L7 L4 L2
//

// main is also a go routine
// no goroutine waits for other goroutines to completes it execution
// order of execution is not guaranteed
