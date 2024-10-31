package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(i int, context context.Context, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case info, good := <-ch:
			if !good {
				fmt.Printf("Worker %d found, that channel is closed!\n", i)
				return
			}
			fmt.Printf("Worked %d got %d!\n", i, info)
		case <-context.Done():
			return
		}
	}
}

func main() {
	ch := make(chan int)
	var n int
	fmt.Scan(&n)
	var wg sync.WaitGroup
	context, cancel := context.WithCancel(context.Background())
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i, context, ch, &wg)
	}
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		fmt.Println("\nOk, finish :)")
		cancel()
		close(ch)
	}()
	for {
		f := false
		select {
		case <-context.Done():
			f = true
		default:
			x := rand.Intn(10)
			ch <- x
			time.Sleep(200 * time.Millisecond)
		}
		if f {
			break
		}
	}
	wg.Wait()
	fmt.Println("That's all, the workers've finished!")
}
