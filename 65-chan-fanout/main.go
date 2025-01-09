package main

import (
	"fmt"
	"runtime"
)

func main() {
	jobs := make(chan int)
	for i := 1; i <= 5; i++ {
		go worker(i, jobs)
	}

	for j := 1; j <= 100; j++ {
		jobs <- j
	}
	close(jobs)
	runtime.Goexit()
}

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Println("worker->", id, "doing the job-->", job)
	}
}
