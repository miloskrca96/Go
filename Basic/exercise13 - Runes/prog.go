package main

import (
	"fmt"
)

var applicationMap map[rune]string = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

func main() {
	fmt.Println(Application("recommended search product ☀🔍"))
	log := "please replace '👎' with '👍'"

	fmt.Println(Replace(log, '👎', '👍'))

	fmt.Println(WithinLimit("hello❗", 10))
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		_, ok := applicationMap[char]
		if ok {
			return applicationMap[char]
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)

	for i := 0; i < len(runes); i++ {
		if runes[i] == oldRune {
			runes[i] = newRune
		}
	}
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	runes := []rune(log)

	if len(runes) > limit {
		return false
	}

	return true
}
