package main

import "fmt"

// buffered channels -> make(chan int)

func main() {
	ch := make(chan int)
	
	fmt.Println("sent to channel ch")
	go func() {
		ch <- 90 // here we send data into channel,
		// so this will block the goroutine until this data get revievied
	}()
	
	// here we are receiving data from the channel,
	// so this will block until the data is sent from the channel
	value := <-ch
	fmt.Println("data received: ", value)
	
	fmt.Println("hello")
}

//! without the go routine that contain the channel, we will get deadlock error
//? that's because ch <- 90 tries to send to the channel on the main
//? goroutine and the reciever in the same goroutines but comes after that
//? and that gourotines as we said is blocked, so this will cause the deadlock
