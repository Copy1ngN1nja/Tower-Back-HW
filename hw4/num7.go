package main

import (
	"fmt"
	"sync"
)

func setMap(m map[int]int, key, val int, wg *sync.WaitGroup, mute *sync.Mutex) {
	defer wg.Done()
	mute.Lock()
	m[key] = val
	mute.Unlock()
}

func main() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	var mute sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go setMap(m, i, i, &wg, &mute)
	}
	wg.Wait()
	for key, val := range m {
		fmt.Println(key, val)
	}
}
