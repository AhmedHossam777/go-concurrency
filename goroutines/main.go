package main

import (
	"fmt"
	"sync"
)

func processTransaction(id int) {
	fmt.Printf("Processing transaction %d\n", id)
	// Simulate work...
	fmt.Printf("Completed transaction %d\n", id)
}

// ?let's imagine that this process tranaction takes time to complete so if
// ?we wanted to do it 5 times we will need to wait the taime it takes * 5

// * the solution of that is using go routines
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done() // this decrement the counter when goroutine is done
			processTransaction(id)
		}(i)
	}

	wg.Wait()
	fmt.Println("all processes is done")
}
