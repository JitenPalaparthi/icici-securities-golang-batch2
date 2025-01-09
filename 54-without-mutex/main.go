package main

import (
	"fmt"
	"sync"
)

// problem application not using mutex

// Do not communicate by sharing memory; instead, share memory by communicating.
var Num int = 0

func main() {
	wg := new(sync.WaitGroup)
	//defer wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= 1000; i++ {
				wg.Add(1)
				go incr(wg)
			}

		}()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= 1000; i++ {
				wg.Add(1)
				go decr(wg)
			}
		}()
	}()
	wg.Wait()
	fmt.Println(Num)
}

func incr(wg *sync.WaitGroup) {
	defer wg.Done()
	Num++

}

func decr(wg *sync.WaitGroup) {
	defer wg.Done()
	Num--
}
