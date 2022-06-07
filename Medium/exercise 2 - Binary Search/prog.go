package main

import (
	"fmt"
)

func main() {
	var arrayNum []int = []int{19, 22, 34, 41, 51, 64, 73, 81, 90, 100, 112}

	fmt.Println(SearchInts(arrayNum, 1))
	fmt.Println(SearchInts(arrayNum, 124))
	fmt.Println(SearchInts(arrayNum, 60))
	fmt.Println(SearchInts(arrayNum, 41))
}

func SearchInts(list []int, key int) int {
	start := 0
	end := len(list) - 1
	current := (start + end) / 2

	if len(list) == 0 {
		return -1
	}

	for key != list[current] {
		if start >= end {
			return -1
		}

		if key > list[current] {
			start = current + 1
		} else {
			end = current - 1
		}

		current = (start + end) / 2
	}

	return current
}
