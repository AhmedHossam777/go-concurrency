package main

import "fmt"

func producer(out chan<- int) {
	defer close(out)
	for i := 0; i < 10; i++ {
		out <- i
	}
}

func consumer(in <-chan int) {
	for val := range in {
		fmt.Println("process: ", val)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)

	consumer(ch)
}

// ! Note
// ? What happens with go (current code):
// * producer runs concurrently in a separate goroutine
// * consumer runs in the main goroutine
// * Both execute simultaneously: producer sends values while the consumer receives them
// * Program completes successfully
// *
// ? What happens without go:
// * producer(ch) runs in the main goroutine
// * It tries to send out <- 0 to an unbuffered channel
// * No one is receiving yet (consumer hasn't started)
// * Deadlock — program blocks forever and panics
