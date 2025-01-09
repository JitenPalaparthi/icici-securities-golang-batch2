package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	sig := make(chan struct{})
	go FibSender(10, ch)
	go FibReceivrer(10, ch, sig)
	<-sig // until it recveices the value
}

func FibSender(r int, ch chan int) {
	if r == 0 {
		println(r)
	} else if r == 1 {
		println(0, 1)
	}

	a, b := 0, 1

	for i := 1; i <= int(r); i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- a
		a, b = b, a+b
	}
}

func FibReceivrer(r int, ch chan int, sig chan struct{}) {
	defer func() {
		// time.Sleep(time.Second*5)
		sig <- struct{}{}
	}()
	for i := 1; i <= r; i++ {
		fmt.Println(<-ch)
	}

}
