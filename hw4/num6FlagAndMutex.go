package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(mute *sync.Mutex, flag *bool) {
	for {
		mute.Lock()
		if *flag {
			mute.Unlock()
			fmt.Println("Goroutine is finished")
			return
		}
		mute.Unlock()
		fmt.Println("Goroutine is working")
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	var mute sync.Mutex
	flag := false
	go worker(&mute, &flag)
	time.Sleep(3 * time.Second)
	mute.Lock()
	flag = true
	mute.Unlock()
	time.Sleep(500 * time.Millisecond)
}
