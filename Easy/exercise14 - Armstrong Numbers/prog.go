package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	num := 153
	fmt.Println(IsNumber(num))
	num = 190
	fmt.Println(IsNumber(num))
	num = 9
	fmt.Println(IsNumber(num))
	num = 10
	fmt.Println(IsNumber(num))
}

func IsNumber(n int) bool {
	var sum float64 = 0

	digitsNum := len(strconv.Itoa(n))
	num := n

	for n > 0 {
		sum += math.Pow(float64(n%10), float64(digitsNum))
		n /= 10
	}

	return float64(num) == sum
}
