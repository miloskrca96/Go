package main

import (
	"fmt"
)

func main() {
	fmt.Println(ShareWith("Jane"))
	fmt.Println(ShareWith("Bob"))
	fmt.Println(ShareWith(""))
}

// ShareWith tells you how to share with the given name
func ShareWith(name string) string {
	var finalName string = name
	if name == "" {
		finalName = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", finalName)
}
