package main

import (
	"fmt"
	"sync"
	"time"
)

func add_to_ch(value int, ch chan<- int) {
	time.Sleep(time.Duration(100) * time.Millisecond)
	ch <- value * value
}

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(a))
	for _, value := range a {
		go func() {
			defer wg.Done()
			add_to_ch(value, ch)
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	ans := 0
	for value := range ch {
		ans += value
	}
	fmt.Println(ans)
}
