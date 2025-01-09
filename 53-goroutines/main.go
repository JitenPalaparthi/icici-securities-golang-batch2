package main

import "sync"

func main() {
	wg := new(sync.WaitGroup)
	defer wg.Wait()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		fibo1("fib-1", 10)
		wg.Done()
	}(wg)

	wg.Add(1)
	go fibo2(wg, "fib-2", 20)

}

func fibo1(name string, r uint) {
	if r == 0 {
		println(r)
	} else if r == 1 {
		println(0, 1)
	}

	a, b := 0, 1

	for i := 1; i <= int(r); i++ {
		go println(name, "-->", a)
		a, b = b, a+b
	}

}

func fibo2(wg *sync.WaitGroup, name string, r uint) {
	defer wg.Done()
	if r == 0 {
		println(r)
	} else if r == 1 {
		println(0, 1)
	}

	a, b := 0, 1

	for i := 1; i <= int(r); i++ {
		println(name, "-->", a)
		a, b = b, a+b
	}

}
