package main

import (
	"fmt"
	"math/rand"
	"time"
)

func queryReplica(replicaID int) <-chan string {
	result := make(chan string)

	go func() {
		delay := time.Duration(rand.Intn(500)) * time.Millisecond
		time.Sleep(delay)
		result <- fmt.Sprintf(
			"Response from replica %d (took %v)", replicaID,
			delay,
		)
	}()

	return result
}

func queryWithRedundancy() string {
	select {
	case r := <-queryReplica(1):
		return r

	case r := <-queryReplica(2):
		return r

	case r := <-queryReplica(3):
		return r

	case <-time.After(1 * time.Second):
		return fmt.Sprintf("All replicas timed out")
	}
}

func main() {
	result := queryWithRedundancy()
	fmt.Println(result)
}

// results:
// aldod@archlinux  ~/Documents/go-concurrency/racing-api   master  go run anti-pattern.go
// racing api, print the first response to process
// aldod@archlinux  ~/Documents/go-concurrency/racing-api   master  go run anti-pattern.go
// Response from replica 2 (took 21ms)
// aldod@archlinux  ~/Documents/go-concurrency/racing-api   master  go run anti-pattern.go
// Response from replica 3 (took 80ms)
// aldod@archlinux  ~/Documents/go-concurrency/racing-api   master  go run anti-pattern.go
// Response from replica 1 (took 106ms)
// aldod@archlinux  ~/Documents/go-concurrency/racing-api   master  go run anti-pattern.go
// Response from replica 1 (took 18ms)
// aldod@archlinux  ~/Documents/go-concurrency/racing-api   master  go run anti-pattern.go
// Response from replica 2 (took 76ms)
// aldod@archlinux  ~/Documents/go-concurrency/racing-api   master  go run anti-pattern.go
// Response from replica 3 (took 78ms)
// aldod@archlinux  ~/Documents/go-concurrency/racing-api   master  go run anti-pattern.go
// Response from replica 2 (took 217ms)
