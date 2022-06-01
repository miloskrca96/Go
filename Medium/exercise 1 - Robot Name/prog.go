package main

import (
	"crypto/rand"
	"fmt"
)

var robotNames map[string]bool = map[string]bool{}

type Robot struct {
	name   string
	turnOn bool
}

func main() {

	for i := 0; i < 9920; i++ {
		r := Robot{}
		r.Name()
	}

	fmt.Println(len(robotNames))

	// r := Robot{}
	// r.Name()
	// fmt.Println(r)
	// r.Name()
	// fmt.Println(r)

	// r.Reset()
	// fmt.Println(r)
}

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nums    = "1234567890"
)

func giveRobotName() string {
	lettersNum := len(letters)
	digitsNum := len(nums)

	lettersArr := make([]byte, 2)
	digitsArr := make([]byte, 3)

	rand.Read(lettersArr)
	for i := 0; i < 2; i++ {
		lettersArr[i] = letters[int(lettersArr[i])%lettersNum]
	}

	rand.Read(digitsArr)
	for i := 0; i < 3; i++ {
		digitsArr[i] = nums[int(digitsArr[i])%digitsNum]
	}

	return string(lettersArr) + string(digitsArr)
}

func (r *Robot) Name() (string, error) {

	if !r.turnOn {
		r.name = giveRobotName()
		r.turnOn = true
		_, ok := robotNames[r.name]
		if !ok {
			robotNames[r.name] = true
		}

		for ok {
			r.name = giveRobotName()
			_, ok := robotNames[r.name]
			if !ok {
				robotNames[r.name] = true
				return r.name, nil
			}
		}
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = "     "
}
