package main

import (
	"errors"
	"fmt"
)

type Planet string

const Earth float64 = 31557600

func Age(seconds float64, planet Planet) (float64, error) {
	switch planet {
	case "Mercury":
		return seconds / (Earth * 0.2408467), nil
	case "Venus":
		return seconds / (Earth * 0.61519726), nil
	case "Mars":
		return seconds / (Earth * 1.8808158), nil
	case "Jupiter":
		return seconds / (Earth * 11.862615), nil
	case "Saturn":
		return seconds / (Earth * 29.447498), nil
	case "Uranus":
		return seconds / (Earth * 84.016846), nil
	case "Neptune":
		return seconds / (Earth * 164.79132), nil
	case "Earth":
		return seconds / Earth, nil
	default:
		return 0.00, errors.New("there is no planet with that name")
	}
}

func main() {
	result, error := Age(1e9, "Earth")
	fmt.Println(result, error)

	result, _ = Age(1e9, "Mercury")
	fmt.Println(result)

	result, _ = Age(1e9, "Venus")
	fmt.Println(result)

	result, _ = Age(1e9, "Mars")
	fmt.Println(result)

	result, _ = Age(1e9, "Jupiter")
	fmt.Println(result)

	result, _ = Age(1e9, "Saturn")
	fmt.Println(result)

	result, _ = Age(1e9, "Uranus")
	fmt.Println(result)

	result, _ = Age(1e9, "Neptune")
	fmt.Println(result)

	result, error = Age(1e9, "Neptune1")
	fmt.Println(result, error)
}
