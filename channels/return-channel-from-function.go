package main

import "fmt"

func generator(nums ...int) <-chan int {
	out := make(chan int, len(nums))

	go func() {
		defer close(out)
		for _, num := range nums {
			out <- num
		}
	}()

	return out
}

func main() {
	channel := generator(1, 2, 3, 4, 5, 6)

	for res := range channel {
		fmt.Println(res)
	}
}
