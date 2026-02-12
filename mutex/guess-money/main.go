package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank balance
	var bankBalance int
	var mu sync.Mutex
	
	// print out the starting values
	fmt.Printf("initial account balance : $%d.00", bankBalance)
	fmt.Println()
	// define weekly revenue
	incomes := []Income{
		{
			Source: "Main Job",
			Amount: 500,
		},
		{
			Source: "Gifts",
			Amount: 10,
		},
		{
			Source: "Part-time job",
			Amount: 50,
		},
		{
			Source: "investments",
			Amount: 100,
		},
	}
	
	wg.Add(len(incomes))
	//  loop through 52 weeks and print out how much he has made
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			
			for week := 1; week <= 52; week++ {
				mu.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				
				fmt.Printf(
					"On week %d, you earned $%d.00 from %s\n", week,
					income.Amount, income.Source,
				)
				mu.Unlock()
			}
			
		}(i, income)
	}
	
	wg.Wait()
	// print out final balance
	fmt.Printf("final bank balance: $%d.00", bankBalance)
	fmt.Println()
}

// ? we put the mutex inside the inner for loop as the race condition is
// ? happening because there is multiple goroutines trying to access the
// ? balance variable so we need to use the mutex when modifying that balance
