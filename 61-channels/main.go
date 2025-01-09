package main

import (
	"fmt"
	"time"
)

func main() {
	pingCh, pongCh := make(chan string), make(chan string)
	go ping(pingCh, pongCh)
	pong(pingCh, pongCh)
}

func ping(pingCh, pongCh chan string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 100)
		pingCh <- "ping"
		fmt.Println(<-pongCh)
	}

}
func pong(pingCh, pongCh chan string) {
	for i := 1; i <= 10; i++ {
		fmt.Println(<-pingCh)
		pongCh <- "pong"
	}
}
