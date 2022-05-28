package main

import (
	"fmt"
)

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

func main() {
	speed := 5
 	batteryDrain := 2
 	car := NewCar(speed, batteryDrain)
 	fmt.Println(car)

 	distance := 800
 	raceTrack := NewTrack(distance)
 	fmt.Println(raceTrack)

 	fmt.Println(CanFinish(car, raceTrack))
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	var numberOfLapsToFinsh = track.distance / car.speed
	
	if (car.battery - numberOfLapsToFinsh*car.batteryDrain) >= 0 {
		return true
	}

	return false
}

