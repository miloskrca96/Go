package main

import "fmt"

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(TotalBirdCount(birdsPerDay))

	fmt.Println(BirdsInWeek(birdsPerDay, 1))

	birdsPerDayNew := []int{2, 5, 0, 7, 4, 1}
	fmt.Println(FixBirdCountLog(birdsPerDayNew))
}

//Function TotalBirdCount that accepts a slice of ints that contains the bird count per day.
//It should return the total number of birds that you counted
func TotalBirdCount(birdsPerDay []int) int {
	var sum int = 0
	for i := range birdsPerDay {
		sum += birdsPerDay[i]
	}
	return sum
}

// BirdsInWeek that accepts a slice of bird counts per day and a week number.
// It returns the total number of birds that you counted in that specific week
func BirdsInWeek(birdsPerDay []int, week int) int {
	var sum int = 0

	for i := 0; i < len(birdsPerDay); i++ {
		if week-1 == i/7 {
			sum += birdsPerDay[i]
		}
	}
	return sum
}

//FixBirdCountLog that takes a slice of birds counted per day as an argument
//and returns the slice after correcting the counting mistake
func FixBirdCountLog(birdsPerDay []int) []int {
	var resultBirdsPerDay = birdsPerDay
	for i := range resultBirdsPerDay {
		if i%2 == 0 {
			resultBirdsPerDay[i]++
		}
	}
	return resultBirdsPerDay
}
