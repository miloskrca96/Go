package main

import (
	"fmt"
)

var carLicence map[string]bool = map[string]bool{
	"car":   true,
	"truck": true,
}

func main() {
	fmt.Println(NeedsLicense("car"))
	fmt.Println(ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf"))
	fmt.Println(CalculateResellPrice(1000, 1))
	fmt.Println(CalculateResellPrice(1000, 5))
	fmt.Println(CalculateResellPrice(1000, 15))
}

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	val, ok := carLicence[kind]
	if ok {
		return val
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var resultString string
	if option1 < option2 {
		resultString = option1 + " is clearly the better choice."
	} else {
		resultString = option2 + " is clearly the better choice."
	}

	return resultString
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var finalPrice float64
	if age < 3 {
		finalPrice = originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		finalPrice = originalPrice * 0.7
	} else if age >= 10 {
		finalPrice = originalPrice * 0.5
	}

	return finalPrice
}
