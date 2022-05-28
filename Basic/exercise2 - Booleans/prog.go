package main

import "testing"

const (
	knightAwake     bool = false
	archerAwake     bool = false
	prisonerAwake   bool = false
	petDogIsPresent bool = true
)

func main() {
	fmt.Println(CanFastAttack(knightAwake))
	fmt.Println(CanSpy(knightAwake, archerAwake, prisonerAwake))
	fmt.Println(CanSignalPrisoner(archerAwake, prisonerAwake))
	fmt.Println(CanFreePrisoner(knightAwake, archerAwake, prisonerAwake, petDogIsPresent))
}

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return !archerIsAwake && prisonerIsAwake
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	return (petDogIsPresent && !archerIsAwake) || (!petDogIsPresent && prisonerIsAwake && !knightIsAwake && !archerIsAwake)
}

