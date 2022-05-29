package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Score("cabbage"))
}

// Given a word, compute the Scrabble score for that word.
func Score(word string) int {
	var score int

	for i := range word {
		switch strings.ToUpper(string(word[i])) {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			score++
		case "D", "G":
			score += 2
		case "B", "C", "M", "P":
			score += 3
		case "F", "H", "V", "W", "Y":
			score += 4
		case "K":
			score += 5
		case "J", "X":
			score += 8
		case "Q", "Z":
			score += 10
		}
	}

	return score
}


