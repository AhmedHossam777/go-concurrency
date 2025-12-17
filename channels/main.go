package main

import "fmt"

// Unbuffered Channels (Synchronous)
func main() {
	ch := make(chan int)

	go func() {
		// this will block until someone receives
		ch <- 42
		fmt.Println("Sent!!")
	}()

	// this will block until someone sends
	value := <-ch
	fmt.Println("received!!", value)
}
