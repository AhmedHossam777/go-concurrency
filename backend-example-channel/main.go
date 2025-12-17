package main

import (
	"fmt"
	"time"
)

func RequestGenerator(request chan<- int) {
	for i := 0; i < 5; i++ {
		requestNumber := i
		fmt.Printf("📨 Received request number: %d\n", requestNumber)

		request <- requestNumber
		time.Sleep(100 * time.Millisecond)

	}

	close(request) // you need to close the opened channel
}

func RequestProcessor(requests <-chan int, result chan<- int) {
	for req := range requests {
		time.Sleep(100 * time.Millisecond)
		result <- req
	}
	close(result)
}

func main() {
	requests := make(chan int, 4)
	results := make(chan int, 4)

	go RequestGenerator(requests)
	go RequestProcessor(requests, results)

	for result := range results {
		fmt.Printf("✅ %s\n", result)
	}
}
