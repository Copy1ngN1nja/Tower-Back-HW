package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []int{1, 2, 3, 4}
	input := make(chan int)
	output := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, val := range a {
			input <- val
		}
		close(input)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range input {
			output <- val * 2
		}
		close(output)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range output {
			fmt.Println(val)
		}
	}()
	wg.Wait()
}
