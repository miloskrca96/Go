package main

import "fmt"

const OvenTime = 40
const oneLayerTime = 2

func main() {
	fmt.Println(RemainingOvenTime(30))
	fmt.Println(PreparationTime(2))
	fmt.Println(ElapsedTime(3, 20))
}

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven > OvenTime {
		return 0
	}

	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	if numberOfLayers < 0 {
		return 0
	}
	return oneLayerTime * numberOfLayers
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	if numberOfLayers < 0 || actualMinutesInOven < 0 {
		return 0
	}

	return actualMinutesInOven + numberOfLayers*oneLayerTime
}
