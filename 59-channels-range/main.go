package main

import (
	"runtime"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	sig1 := make(chan struct{})
	sig2 := make(chan struct{})

	go GenereateSq(ch1)
	go GenereateSq(ch2)
	//go Receiver(ch, sig)
	go receiverR(ch1, sig1)
	go receiverR(ch2, sig2)

	<-sig1 // just a signal to be triggred
	<-sig2
}

func GenereateSq(ch chan<- int) { // sender channel
	for i := 1; i <= 10; i++ {
		ch <- i * i
		//time.Sleep(time.Millisecond * 500)
	}
	close(ch) // only a sender can close a channel
}

func Receiver(ch <-chan int, sig chan<- struct{}) {
	for {
		v, ok := <-ch
		if ok {
			println(v, ok)
		} else {
			println("channel is closed", ok, v)
			sig <- struct{}{}
			close(sig)
			runtime.Goexit()
		}
	}
}

func receiverR(ch <-chan int, sig chan<- struct{}) {
	for v := range ch { // range loop iterates until the channel is closed
		println(v)
	}
	sig <- struct{}{}
	close(sig)
}
