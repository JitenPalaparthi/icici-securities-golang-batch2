package main

func main() {
	var ch chan int     // nil
	ch = make(chan int) // unbuffred channel
	// go func() {
	// 	ch <- 100
	// }()
	// v := <-ch
	go func() {
		v := <-ch
		println(v)
	}()

	ch <- 1000
}

// chan is keyword to create a channel
// make func is used to instantiate a chan
// buffered and unbuffered
// to send a value to a channel ch <- 100
// to receive a value from the channel <- ch

// unbuffered channels, when a sender(goroutine) sends a data, the sender(goroutine) is blocked until the receiver(goroutine) receives the value
// unbuffered channel, when a receiver(goroutine) receives data, the receiver(goroutine) is blocked until the sender(goroutine) sends a value
// there is nothing like sender should send first so that go routune has to be started
