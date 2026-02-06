package main

import (
	"fmt"
	"time"
)

// var config map[string]string

func main() {
	go loadConfig()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Database URL:", config["database_url"])

}

func loadConfig() {
	// Simulate loading from file/network
	time.Sleep(50 * time.Millisecond) // Usually takes 50ms...

	config = map[string]string{
		"database_url": "postgres://localhost:5432/app",
	}
}
