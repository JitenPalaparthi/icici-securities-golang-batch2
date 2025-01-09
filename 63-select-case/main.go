package main

import (
	"fmt"
	"time"
)

func main() {
	sig := make(chan struct{})
	ch1 := Gen(time.Millisecond*10, "gen-1", 2)
	ch2 := Gen(time.Millisecond*10, "gen-2", 3)

	go func(sig chan struct{}) {
		done1, done2 := false, false
		for {
			if done1 && done2 {
				sig <- struct{}{}
				break
			}
			select {
			case v, ok := <-ch1:
				if ok {
					println(v)
				} else {
					done1 = true
				}
			case v, ok := <-ch2:
				if ok {
					println(v)
				} else {
					done2 = true
				}
			default:
			}
		}
		//return sig
	}(sig)
	<-sig
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
