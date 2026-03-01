package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	
	done := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()
	
	for {
		select {
		case <-done:
			fmt.Println("done")
			return
		case t := <-ticker.C:
			fmt.Println("Tick at: ", t)
		}
	}
	
}
