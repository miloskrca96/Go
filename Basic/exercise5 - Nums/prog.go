package main

import (
	"fmt"
)

const (
	oneCarCost uint = 1e4
	tenCarCost uint = 95e3
)

func main() {
	fmt.Println(CalculateWorkingCarsPerHour(1547, 90))
	fmt.Println(CalculateWorkingCarsPerMinute(1105, 90))
	fmt.Println(CalculateCost(21))
	fmt.Println(CalculateCost(37))
}

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	if successRate < 0 || successRate > 100 || productionRate < 0 {
		return 0
	}

	return (successRate / 100) * float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	if successRate < 0 || successRate > 100 || productionRate < 0 {
		return 0
	}

	return int((successRate/100)*float64(productionRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {

	var numOfTenGroup uint = uint(carsCount) / 10
	var rest uint = uint(carsCount) % 10

	return numOfTenGroup*tenCarCost + rest*oneCarCost
}

