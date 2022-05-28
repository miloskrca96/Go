package main

import "fmt"

type Rank []bool
type Chessboard map[string]Rank

func main() {
	var board Chessboard = Chessboard{
		"A": Rank{true, false, true, false, false, false, false, true},
		"B": Rank{false, false, false, false, true, false, false, false},
		"C": Rank{false, false, true, false, false, false, false, false},
		"D": Rank{false, false, false, false, false, false, false, false},
		"E": Rank{false, false, false, false, false, true, false, true},
		"F": Rank{false, false, false, false, false, false, false, false},
		"G": Rank{false, false, false, true, false, false, false, false},
		"H": Rank{true, true, true, true, true, true, false, true},
	}

	fmt.Println(CountInRank(board, "H"))
	fmt.Println(CountInFile(board, 1))
	fmt.Println(CountAll(board))
	fmt.Println(CountOccupied(board))

}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var sum int

	val, ok := cb[rank]
	if !ok {
		return 0
	}

	for i := range val {
		if val[i] {
			sum++
		}
	}

	return sum
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	var sum int

	if file < 1 || file > 8 {
		return 0
	} else {
		for _, val := range cb {
			if val[file-1] {
				sum++
			}
		}
	}

	return sum
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var sum int
	for val, _ := range cb {
		sum += len(cb[val])
	}
	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var sum int
	for val, _ := range cb {
		for i := range cb[val] {
			if cb[val][i] {
				sum++
			}
		}
	}
	return sum
}
