package main

import (
	"errors"
	"testing"
)

type Classification string

const (
	ClassificationDeficient Classification = "Deficient"
	ClassificationPerfect   Classification = "Perfect"
	ClassificationAbundant  Classification = "Abundant"
	ClassificationInvalid   Classification = "Invalid"
)

var ErrOnlyPositive = errors.New("Not positive integer")

func main() {
	fmt.Println(Classify(6))
	fmt.Println(Classify(28))
	fmt.Println(Classify(12))
	fmt.Println(Classify(24))
	fmt.Println(Classify(8))
}

func Classify(num int64) (Classification, error) {
	if num <= 0 {
		return ClassificationInvalid, ErrOnlyPositive
	} else {
		var sum int

		for i := 1; i <= int(num)/2; i++ {
			if int(num)%i == 0 {
				sum += i
			}
		}

		switch {
		case sum == int(num):
			return ClassificationPerfect, nil
		case sum > int(num):
			return ClassificationAbundant, nil
		default:
			return ClassificationDeficient, nil

		}
	}
}
