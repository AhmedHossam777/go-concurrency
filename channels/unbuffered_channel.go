package main

import "fmt"

// ? Key behavior: Sender only blocks when buffer is full.
// ? Like a mailbox with limited slots.
func main() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 10
	ch <- 10
	//ch <- 10 // this will cause a deadlock error

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
