package main

import (
	"fmt"
)

func main() {
	fmt.Println(Difference(10))
	fmt.Println(Difference(7))
}

func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Find the difference between the square of the sum and
//  the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
