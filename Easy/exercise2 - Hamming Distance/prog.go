package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(Distance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"))
}

// If we compare two strands of DNA and count the differences between them we can see how many mistakes occurred.
// This is known as the "Hamming Distance"
func Distance(a, b string) (int, error) {
	var hammingDistance int

	if len(a) != len(b) {
		return 0, errors.New("Not equal length")
	}

	for i := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
