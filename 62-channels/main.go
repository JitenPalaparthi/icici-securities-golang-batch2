package main

import (
	"runtime"
	"time"
)

func main() {

	ch := GenSq(time.Millisecond * 10)
	go func() {
		for v := range ch {
			println(v)
		}
	}()
	runtime.Goexit()
}

func GenSq(d time.Duration) chan int {
	ch := make(chan int)
	//timeout := TimeOut(d)
	timeout := time.After(d)
	i := 1
	go func() {
	out:
		for {
			select { // only with channels
			case <-timeout:
				close(ch)
				break out
			case ch <- i * i:
			}
			i++
		}
	}()
	return ch
}

func TimeOut(d time.Duration) chan struct{} {
	sig := make(chan struct{})
	go func() {
		time.Sleep(d) // after the duration
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}
