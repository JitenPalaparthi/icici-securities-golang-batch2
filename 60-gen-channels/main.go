package main

import (
	"time"
)

func main() {
	ch := GenSq()
	<-receiverR(ch)
	<-receiverR(GenSq())
}

func GenSq() chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 100)
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

func receiverR(ch <-chan int) chan struct{} {
	sig := make(chan struct{})
	go func() {
		for v := range ch { // range loop iterates until the channel is closed
			println(v)
		}
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}
