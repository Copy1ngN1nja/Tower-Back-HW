package main

import (
	"fmt"
	"sync"
	"time"
)

func make_squares(value int, ch chan<- int) {
	time.Sleep(time.Duration(500) * time.Millisecond)
	ch <- value * value
}

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(a))

	for _, value := range a {
		go func(value int) {
			defer wg.Done()
			make_squares(value, ch)
		}(value)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for value := range ch {
		fmt.Printf("%d ", value)
		time.Sleep(700 * time.Millisecond)
	}
}
