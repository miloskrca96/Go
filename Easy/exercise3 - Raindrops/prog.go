package main

import (
	"fmt"
	"strconv"
)

var numStringConversions map[int]string = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func main() {
	fmt.Println(Convert(28))
	fmt.Println(Convert(30))
	fmt.Println(Convert(34))
}

func Convert(number int) string {
	var resultString string
	var isFactor bool = false

	for key, val := range numStringConversions {
		if number%key == 0 {
			resultString += val
			isFactor = true
		}
	}
	if isFactor {
		return resultString
	}

	return strconv.Itoa(number)
}
