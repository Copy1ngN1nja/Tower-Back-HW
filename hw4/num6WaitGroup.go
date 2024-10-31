package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("Goroutine is working")
		time.Sleep(400 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(&wg)
	wg.Wait()
	fmt.Println("Goroutine is finished")
}
