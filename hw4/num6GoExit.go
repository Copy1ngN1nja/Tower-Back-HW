package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker() {
	for {
		if time.Now().Second()%10 == 0 {
			fmt.Println("Goroutine is finished")
			runtime.Goexit()
		}
		fmt.Println("Goroutine is working")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go worker()
	time.Sleep(11 * time.Second)
}
