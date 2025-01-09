package main

import "fmt"

func main() {

	var ch chan int        // nil
	ch = make(chan int, 2) // bufferd channel
	ch <- 100
	ch <- 200
	fmt.Println(<-ch)
	ch <- 300
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// chan is keyword to create a channel
// make func is used to instantiate a chan
// buffered and unbuffered
// to send a value to a channel ch <- 100
// to receive a value from the channel <- ch

// the sender does not block until the buffer is full
