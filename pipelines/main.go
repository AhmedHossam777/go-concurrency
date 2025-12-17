package main

import (
	"fmt"
)

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

func sq(dataChannel <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range dataChannel {
			out <- num * num
		}
		close(out)
	}()

	return out
}

func main() {
	fmt.Println("Hello, Pipelines")

	// input:
	nums := []int{1, 2, 3, 7, 6}

	// stage 1: get the data into channel
	dataChannel := sliceToChannel(nums)

	// stage 2: make operation on data and pass it into another channel
	finalChannel := sq(dataChannel)

	// stage 3: print the result
	for num := range finalChannel {
		fmt.Println(num)
	}
}
