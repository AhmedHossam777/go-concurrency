package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for i, word := range words {
		wg.Add(1)
		go printSomething(fmt.Sprintf("%d: %s", i, word), &wg)
	}

	wg.Wait()
	wg.Add(1)
	printSomething("second", &wg)

}
