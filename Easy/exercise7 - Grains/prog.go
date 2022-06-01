package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	sum := Total()
	square, _ := Square(10)

	fmt.Println(sum)
	fmt.Println(square)
}

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Num cannot be smaller than 1 and greater than 64")
	}

	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	return (1 << 64) - 1
}
