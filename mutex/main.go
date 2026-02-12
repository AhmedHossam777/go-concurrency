package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	msg = s
	mu.Unlock()
}

func main() {
	msg = "hello, world"

	var mutex sync.Mutex
	// here is the problem of the race condition
	// two goroutines is trying to modofy the same data
	// we realy don't know which goroutine will
	wg.Add(2)
	go updateMessage("hello, universe!", &mutex)
	go updateMessage("hello, cosmos!", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
