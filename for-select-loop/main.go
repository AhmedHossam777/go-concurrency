package main

import "fmt"

func main() {
	charChannel := make(chan string, 3)
	chars := [3]string{"a", "b", "c"}

	for _, char := range chars {
		select {
		case charChannel <- char:

		}
	}

	close(charChannel)

	//* the following lines will return only the first character as the channel get out the data like queue
	//result := <-charChannel
	//fmt.Println(result)

	//* to get the all data we will use a for loop
	for result := range charChannel {
		fmt.Println(result) // a  b  c
	}

	fmt.Println("hello,this is \"for select\" concurrency pattern")
}
