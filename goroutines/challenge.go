package main

import (
	"fmt"
	"sync"
)

var msg string

func printMessage() {
	fmt.Println(msg)
}

func updateMessage(s string, wg *sync.WaitGroup) {
	msg = s
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup

	msg = "hello, world"

	wg.Add(1)
	go updateMessage("hello, universe", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("hello, cosmos", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("hello, dod", &wg)
	wg.Wait()
	printMessage()

}
