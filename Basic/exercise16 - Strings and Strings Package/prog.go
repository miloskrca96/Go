package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(WelcomeMessage("Jane"))
	fmt.Println(AddBorder("Welcome!", 10))

	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`

	fmt.Println(CleanupMessage(message))
}

// WelcomeMessage that accepts the name of the customer as a string argument and
// returns the desired message as a string
func WelcomeMessage(name string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(name))
}

// AddBorder adds a border to a welcome message.
func AddBorder(message string, numOfStars int) string {
	stars := strings.Repeat("*", numOfStars)
	return fmt.Sprintf("%s\n%s\n%s", stars, message, stars)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
