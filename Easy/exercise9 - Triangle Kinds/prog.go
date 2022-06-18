package main

import "fmt"

type Kind int

const (
	Equ = iota // equilateral
	Iso		   // isosceles
	Sca		   // scalene
	NaT        // not a triangle
)

func main() {
	fmt.Println(KindFromSides(3, 3, 3))
	fmt.Println(KindFromSides(3, 4, 4))
	fmt.Println(KindFromSides(3, 4, 5))
	fmt.Println(KindFromSides(1, 1, 3))
}

func isTriangle(a, b, c float64) bool {
	if a > 0 && b > 0 && c > 0 && a+b >= c && a+c >= b && b+c >= a {
		return true
	}

	return false
}

func KindFromSides(a, b, c float64) Kind {
	var k Kind = NaT

	if isTriangle(a, b, c) {
		if a == b && a == c {
			k = Equ
		} else if a == b || a == c || b == c {
			k = Iso
		} else if a != b && a != c && b != c {
			k = Sca
		}
	}

	return k
}
