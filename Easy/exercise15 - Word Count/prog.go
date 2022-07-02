package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Frequency map[string]int

func main() {
	str := "\"That's the password: 'PASSWORD 123'!\", cried the Special Agent.\nSo I fled."
	fmt.Println(WordCount(str))

	str = "car: carpet as java: javascript!!&@$%^&"
	fmt.Println(WordCount(str))
}

func Split(r rune) bool {
	return r == ' ' || r == '.' || r == '!' || r == '?' || r == ','
}

func WordCount(phrase string) Frequency {
	var words Frequency = map[string]int{}

	reg, _ := regexp.Compile("[^a-zA-Z0-9_'.,!? ]+")
	processedString := reg.ReplaceAllString(phrase, "")
	splitedString := strings.FieldsFunc(processedString, Split)

	for i := range splitedString {

		if splitedString[i][0] == '\'' {
			splitedString[i] = splitedString[i][1:]
		}

		wordLen := len(splitedString[i]) - 1

		if splitedString[i][wordLen] == '\'' {
			splitedString[i] = splitedString[i][:wordLen]
		}

		words[strings.ToLower(splitedString[i])]++
	}

	return words
}
