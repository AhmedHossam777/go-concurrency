package main

import (
	"fmt"
	"time"
)

func slowAPICall() chan string {
	result := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		result <- "Api response"
	}()

	return result
}

func fetchWithTimeout(timeout time.Duration) (string, error) {
	select {
	case result := <-slowAPICall():
		return result, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("request timed out after %v", timeout)
	}
}

func main() {
	fmt.Println("calling api with timeout")

	// This will timeout, as the timeout is one second and our api call takes 2 seconds
	result, err := fetchWithTimeout(1 * time.Second)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
