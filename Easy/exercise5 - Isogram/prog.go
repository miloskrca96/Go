package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(IsIsogram("six-year-old"))
	fmt.Println(IsIsogram("lumberjacks"))
	fmt.Println(IsIsogram("background"))
	fmt.Println(IsIsogram("downstream"))
	fmt.Println(IsIsogram("isograms"))
}

// Determine if a word or phrase is an isogram.
func IsIsogram(word string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	processedString := strings.ToLower(reg.ReplaceAllString(word, ""))

	for i := range processedString {
		if strings.Count(processedString, string(processedString[i])) > 1 {
			return false
		}
	}

	return true
}
