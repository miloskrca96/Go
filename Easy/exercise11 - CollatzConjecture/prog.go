package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(CollatzConjecture(0))
	fmt.Println(CollatzConjecture(12))
	fmt.Println(CollatzConjecture(31))
}

func CollatzConjecture(n int) (int, error) {
	numOfSteps := 0

	if n <= 0 {
		return 0, errors.New("Error")
	}

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}

		numOfSteps++
	}

	return numOfSteps, nil
}
