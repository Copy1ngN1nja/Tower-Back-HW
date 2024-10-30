package main

import (
	"context"
	"fmt"
	"time"
)

func worker(context context.Context) {
	for {
		select {
		case <-context.Done():
			fmt.Println("Goroutine is finished")
			return
		default:
			fmt.Println("Goroutine is working")
			time.Sleep(400 * time.Millisecond)
		}
	}
}

func main() {
	context, cancel := context.WithCancel(context.Background())
	go worker(context)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(500 * time.Second)
}
