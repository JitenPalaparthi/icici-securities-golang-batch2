package main

import (
	"fmt"
	"sync"
)

// problem application not using mutex

// Do not communicate by sharing memory; instead, share memory by communicating.
var (
	Num int             = 0
	mu  *sync.Mutex     = new(sync.Mutex)
	wg  *sync.WaitGroup = new(sync.WaitGroup)
)

func main() {
	//defer wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= 10000000; i++ {
				wg.Add(1)
				go incr(wg, mu)
			}
		}()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= 10000000; i++ {
				wg.Add(1)
				go decr(wg, mu)
			}
		}()
	}()
	wg.Wait()
	fmt.Println(Num)

	var i int
	go func() {
		i = add(10, 20)
		println(i)
	}()

}

func add(a, b int) int {
	return a + b
}

func incr(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	Num++
	mu.Unlock()

}

func decr(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	Num--
	mu.Unlock()
}
