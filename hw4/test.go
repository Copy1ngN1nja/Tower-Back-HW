package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Function to process a single number
func processNumber(number int, ch chan<- int) {
	// Simulate processing time
	sleep := rand.Intn(500) + 100
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	// Send the processed number to the channel
	ch <- number * number
}

func main() {
	// Initialize a slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Create a channel to send processed numbers
	ch := make(chan int)

	// Initialize a wait group
	var wg sync.WaitGroup

	// Add 5 to the wait group count (one for each goroutine)
	wg.Add(len(numbers))

	// Launch 5 goroutines to process numbers concurrently
	for _, number := range numbers {
		go func(number int) {
			defer wg.Done() // Decrement the wait group count when the goroutine finishes
			processNumber(number, ch)
		}(number)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(ch) // Close the channel after all goroutines have finished
	}()

	// Print the processed numbers
	for v := range ch {
		fmt.Println(v)
	}
}
