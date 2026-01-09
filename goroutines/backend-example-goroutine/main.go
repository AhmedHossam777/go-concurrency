package main

import (
	"fmt"
	"sync"
	"time"
)

type UserProfile struct {
	User    string
	Orders  []string
	Friends []string
}

func fetchUser(userID int) string {
	time.Sleep(100 * time.Millisecond) // Simulate DB
	return fmt.Sprintf("User_%d", userID)
}

func fetchOrders(userID int) []string {
	time.Sleep(150 * time.Millisecond) // Simulate DB
	return []string{"Order1", "Order2"}
}

func fetchFriends(userID int) []string {
	time.Sleep(120 * time.Millisecond) // Simulate API call
	return []string{"Friend1", "Friend2"}
}

func GetUserProfile(userId int) UserProfile {
	var userProfile UserProfile
	
	var wg sync.WaitGroup
	
	wg.Add(3)
	
	go func() {
		defer wg.Done()
		userProfile.User = fetchUser(userId)
	}()
	
	go func() {
		defer wg.Done()
		userProfile.Friends = fetchFriends(userId)
	}()
	
	go func() {
		defer wg.Done()
		userProfile.Orders = fetchOrders(userId)
	}()
	
	wg.Wait()
	return userProfile
}

func main() {
	start := time.Now()
	profile := GetUserProfile(42)
	fmt.Printf(
		"Got profile in %v:  %+v\n", time.Since(start), profile,
	) //Got profile in 370.756402ms:  {User:User_42 Orders:[Order1 Order2] Friends:[Friend1 Friend2]}
	
	//? after adding goroutines
	//?Got profile in 150.4933ms:  {User:User_42 Orders:[Order1 Order2] Friends:[Friend1 Friend2]}
}
