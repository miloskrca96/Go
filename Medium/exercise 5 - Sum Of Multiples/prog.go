package main

import "fmt"

func main() {
	fmt.Println(SumMultiples(4, 3, 0))
	fmt.Println(SumMultiples(1e3, 3, 5))
	fmt.Println(SumMultiples(1e4, 43, 47))
	fmt.Println(SumMultiples(150, 5, 6, 8))
}

func SumMultiples(limit int, divisors ...int) int {
	var sum int
	m := make(map[int]int)

	for _, val := range divisors {
		multiplier := 1
		res := 0

		for res < limit-val && val > 0 {
			res = val * multiplier
			if _, ok := m[res]; !ok {
				sum += res
			}

			m[res] = res
			multiplier++
		}
	}
	return sum
}
