package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Worker(id int, jobs <-chan int, quit <-chan struct{}) {
	for {
		select {
		case <-quit:
			fmt.Println("Shutting down!")
			return
		case job := <-jobs:
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	jobs := make(chan int, 10)
	quit := make(chan struct{})

	// Start workers
	for i := 1; i <= 3; i++ {
		go Worker(i, jobs, quit)
	}

	// Listen for OS signals
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)

	// Send some jobs
	go func() {
		for i := 1; i <= 100; i++ {
			jobs <- i
			time.Sleep(200 * time.Millisecond)
		}
	}()

	// Wait for shutdown signal
	<-signChan
	fmt.Println("\nReceived shutdown signal...")
	close(quit) // Signal all workers to stop

	time.Sleep(1 * time.Second) // Give workers time to finish
	fmt.Println("Shutdown complete")
}
