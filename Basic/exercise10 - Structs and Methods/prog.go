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

	car.Drive()
	fmt.Println(car)
	fmt.Println(car.DisplayDistance())
	fmt.Println(car.DisplayBattery())

	trackDistance := 120
	fmt.Println(car.CanFinish(trackDistance))
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// DisplayDistance method on Car to return the distance
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery method on Car to return the battery percentage
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish checks if a car is able to finish a certain track.
func (car *Car) CanFinish(trackDistance int) bool {
	var numberOfLapsToFinsh = trackDistance / car.speed

	if (car.battery - numberOfLapsToFinsh*car.batteryDrain) >= 0 {
		return true
	}

	return false
}
