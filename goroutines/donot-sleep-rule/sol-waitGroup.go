package main

import (
	"fmt"
	"sync"
)

var config map[string]string

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		loadConfigs()
	}()

	wg.Wait()
	fmt.Println("Database URL:", config["database_url"])

}

func loadConfigs() {
	config = map[string]string{
		"database_url": "postgres://localhost:5432/app",
	}
}
