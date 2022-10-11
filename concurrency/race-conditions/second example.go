package main

import (
	"fmt"
	"strings"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

var wg sync.WaitGroup
var balanceLock sync.Mutex

func main() {
	var bankBalance int
	var totalWeeks = 52
	incomes := []Income{
		{Source: "Software Engineer", Amount: 1000},
		{Source: "Part time", Amount: 500},
		{Source: "Gifts", Amount: 100},
		{Source: "Investments", Amount: 200},
	}
	fmt.Printf("Total incomes: #%d\n", len(incomes))
	fmt.Printf("Initial bank balance $%d\n", bankBalance)
	wg.Add(len(incomes))
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= totalWeeks; week++ {
				balanceLock.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balanceLock.Unlock()
				fmt.Printf("Week #%d incomes: %d.00 from %s:\n", week, income.Amount, strings.ToLower(income.Source))
			}
		}(i, income)
	}
	wg.Wait()

	fmt.Printf("Total bank balance after #%dweeks is:%d.00\n", totalWeeks, bankBalance)
}
