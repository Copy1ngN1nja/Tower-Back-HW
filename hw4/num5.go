package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	var n int
	fmt.Scan(&n)
	go func() {
		for {
			x := rand.Intn(10)
			ch <- x
			time.Sleep(200 * time.Millisecond)
		}
	}()
	alarm := time.After(time.Duration(n) * time.Second)
	for {
		select {
		case <-alarm:
			fmt.Println("Program is over!")
			return
		case value := <-ch:
			fmt.Printf("Received %d\n", value)
		}
	}
}
