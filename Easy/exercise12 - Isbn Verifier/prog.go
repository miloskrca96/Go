package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(IsValidISBN("35982P15088"))
}

func IsValidISBN(isbn string) bool {
	sum := 0
	mul := 10
	num := 0

	for i := range isbn {
		if isbn[i] >= 'A' && isbn[i] <= 'Z' && isbn[i] != 'X' {
			return false
		}

		if isbn[i] == 'X' {
			sum += DoAddition(10, &mul, &num)

			if i != len(isbn)-1 {
				return false
			}
		}

		if isbn[i] >= '0' && isbn[i] <= '9' {
			res, _ := strconv.Atoi(string(isbn[i]))
			sum += DoAddition(res, &mul, &num)
		}
	}

	return sum%11 == 0 && num == 10
}

func DoAddition(num int, num1, num2 *int) int {
	result := num * *(num1)
	*(num1)--
	*(num2)++

	return result
}
