package main

import "fmt"

func main() {
	config := make(chan map[string]string, 1)
	readyConfig := make(chan struct{})

	go loadConfiguration(config, readyConfig)
	<-readyConfig

	cfg := <-config
	fmt.Println("Database URL:", cfg["database_url"])

}

func loadConfiguration(config chan<- map[string]string, ready chan<- struct{}) {
	cfg := map[string]string{
		"database_url": "postgres://localhost:5432/app",
	}

	config <- cfg
	close(ready) // the signal to remove the block
}
