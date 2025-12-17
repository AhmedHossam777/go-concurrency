package main

import (
	"fmt"
	"sync"
	"time"
)

type UserProfile struct {
	UserId  int
	Orders  []string
	Friends []string
}

func FetchUser(id int) int {
	time.Sleep(100 * time.Millisecond) // Simulate DB
	fmt.Println("user id: ", id)
	return id
}

func FetOrders(userID int) []string {
	time.Sleep(150 * time.Millisecond) // Simulate DB
	return []string{"Order1", "Order2"}
}

func FetUser(userID int) []string {
	time.Sleep(120 * time.Millisecond) // Simulate API call
	return []string{"Friend1", "Friend2"}
}

func GetUserProfile(userId int) UserProfile {
	var Profile UserProfile
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		Profile.UserId = FetchUser(userId)
	}()

	go func() {
		defer wg.Done()
		Profile.Orders = FetOrders(userId)
	}()

	go func() {
		defer wg.Done()
		Profile.Friends = FetUser(userId)
	}()

	wg.Wait() // wait until all three are done

	return Profile
}

// ! Note: we cannot use wg inside the main function, this will lead the program to end immediately
func main() {
	start := time.Now()
	profile := GetUserProfile(90)

	fmt.Printf("Got profile in %v:  %+v\n", time.Since(start), profile)
	// Got profile in 150.559883ms:  {UserId:90 Orders:[Order1 Order2] Friends:[Friend1 Friend2]}
	//? as you see we get the result in 150ms instead of 370 ms
}
