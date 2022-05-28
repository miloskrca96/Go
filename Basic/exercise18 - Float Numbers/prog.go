package main

import (
	"fmt"
)

func main() {
	fmt.Println(InterestRate(200.75))
	fmt.Println(Interest(200.75))
	fmt.Println(AnnualBalanceUpdate(200.75))

	balance := 200.75
	targetBalance := 214.88
	fmt.Println(YearsBeforeDesiredBalance(balance, targetBalance))
}

// Function to calculate the interest rate based on the specified balance
func InterestRate(interestRate float64) float32 {
	switch {
	case interestRate < 0:
		return 3.213
	case interestRate >= 0 && interestRate < 1e3:
		return 0.5
	case interestRate >= 1e3 && interestRate < 5e3:
		return 1.621
	case interestRate >= 5e3:
		return 2.475
	default:
		return 0
	}
}

// Function to calculate the interest based on the specified balance
func Interest(interestRate float64) float64 {
	return float64(InterestRate(interestRate)) * interestRate / 100
}

// Function to calculate the annual balance update, taking into account the interest rate
func AnnualBalanceUpdate(interestRate float64) float64 {
	return Interest(interestRate) + interestRate
}

// YearsBeforeDesiredBalance() function to calculate the minimum number of years
// required to reach the desired balance, taking into account that each year, interest is added to the balance
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var i = balance
	var numberOfYears int

	for i <= targetBalance {
		i = AnnualBalanceUpdate(i)
		numberOfYears++
	}

	return numberOfYears
}
