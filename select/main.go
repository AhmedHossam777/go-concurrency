package main

import "fmt"

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	go func() {
		anotherChannel <- "cow"
	}()

	// this select statement will block until one of it's statement runs
	select {
	case messageFromMyChannel := <-myChannel:
		fmt.Println(messageFromMyChannel)
	case messageFromAnotherChannel := <-anotherChannel:
		fmt.Println(messageFromAnotherChannel)
	}
}
