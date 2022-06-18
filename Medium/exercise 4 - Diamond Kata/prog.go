package main

import (
	"errors"
	"fmt"
)

func intAbsolute(n int) int {
	if n < 0 {
		return n * (-1)
	}

	return n
}

func main() {
	resultString, _ := Gen('D')
	fmt.Println(resultString)
}

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("Letter out of range")
	} else {
		var letters []string = []string{}
		var numOfLetters int
		var resultString string

		for i := 'A'; i <= rune(char); i++ {
			letters = append(letters, string(i))
			numOfLetters++
		}

		dimension := 2*numOfLetters - 1

		k := -1

		for i := 0; i < dimension; i++ {
			if i <= numOfLetters-1 {
				k++
			} else {
				k--
			}

			positionFirst := intAbsolute(k - numOfLetters + 1)
			positionSecond := k + numOfLetters - 1

			for j := 0; j < dimension; j++ {
				if j == positionFirst || j == positionSecond {
					resultString += letters[k]
				} else {
					resultString += " "
				}
			}
			resultString += "\n"
		}
		return resultString, nil
	}
}
