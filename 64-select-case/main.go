package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := Gen(time.Millisecond*2, "gen-1", 2)
	ch2 := Gen(time.Millisecond*2, "gen-2", 3)
	ch3 := Gen(time.Millisecond*2, "gen-3", 4)
	ch4 := Gen(time.Millisecond*2, "gen-4", 5)
	ch5 := Gen(time.Millisecond*1, "gen-5", 6)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go Receive(wg, ch1, ch2, ch3, ch4, ch5)
	wg.Wait()
}

func Gen(d time.Duration, name string, multiply int) chan string {
	ch := make(chan string)
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
			case ch <- fmt.Sprint(name, "-->", i*multiply):
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

func Receive(wg *sync.WaitGroup, chs ...chan string) {
	// fn := func(ch chan string) {
	// 	for v := range ch {
	// 		println(v)
	// 	}
	// }
	// for _, ch := range chs {
	// 	go fn(ch)
	// }

	for _, ch := range chs {
		wg.Add(1)
		go func(ch chan string) {
			for v := range ch {
				println(v)
			}
			wg.Done()
		}(ch)
	}
	wg.Done()
}
