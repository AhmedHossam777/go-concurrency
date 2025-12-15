package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchUserDataWithConcurrency(userId int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("fetching user data of user: ", userId)

}

func fetchUserData(userId int) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("fetching user data of user: ", userId)

}
func main() {
	//fetchUserData(1)
	//fetchUserData(2)
	//fetchUserData(3)

	//! this took 3 seconds to be done!!!!\

	//* Using go concurrency !!!!!!!
	var wg sync.WaitGroup
	wg.Add(3)
	go fetchUserDataWithConcurrency(4, &wg)
	go fetchUserDataWithConcurrency(5, &wg)
	go fetchUserDataWithConcurrency(6, &wg)
	// this took only one sec
	wg.Wait()

}
