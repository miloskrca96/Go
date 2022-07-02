package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	resultStr := ""
	splitedStr := strings.FieldsFunc(s, Split)

	for i := range splitedStr {
		reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
		formatedStr := reg.ReplaceAllString(splitedStr[i], "")
		resultStr += string(formatedStr[0])
	}
	return strings.ToUpper(resultStr)
}

func Split(r rune) bool {
	return r == ' ' || r == '-' || r == '_'
}

func main() {
	fmt.Println(Abbreviate("Complementary metal-oxide semiconductor"))
	fmt.Println(Abbreviate("GNU Image Manipulation Program"))
	fmt.Println(Abbreviate("First In, First Out"))
	fmt.Println(Abbreviate("The Road _Not_ Taken"))
	fmt.Println(Abbreviate("In Real Life"))
}
