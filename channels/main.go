package main

import "fmt"

//! Unbuffered Channels (Synchronous)
//? ch := make(chan int) // Unbuffered - capacity 0
//* Key behavior: Sender blocks until receiver is ready, and vice versa. It's like a phone call — both parties must be present.
//func main() {
//	ch := make(chan int)
//
//	go func() {
//		// this will block until someone receives
//		ch <- 42
//		fmt.Println("Sent!!")
//	}()
//
//	// this will block until someone sends
//	value := <-ch
//	fmt.Println("received!!", value)
//}

// ! Buffered Channels (Asynchronous... up to a point)
// ? ch := make(chan int,3) Buffer size 3
// * Key behavior: Sender only blocks when buffer is full. Like a mailbox with limited slots.
func main() {
	ch := make(chan int, 3)

	ch <- 1 // won't block, won't block as the buffer has space
	ch <- 2 // won't block, won't block as the buffer has space
	//ch <- 3 this will block as the buffer will be full you must wait for someone else to receive to make this rum

	val1 := <-ch // 1
	val2 := <-ch // 2

	fmt.Println(val1)
	fmt.Println(val2)
}
