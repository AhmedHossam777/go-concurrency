package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "message from channel one"
	}()
	
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "message from channel two"
	}()
	
	// ! without the select statement: this will return an error
	// for i := 0; i < 2; i++ {
	// 	msg1 := <-ch1
	// 	fmt.Println(msg1)
	//
	// 	msg2 := <-ch1
	// 	fmt.Println(msg2)
	// }
	
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
