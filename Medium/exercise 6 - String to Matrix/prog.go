package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	elements [][]int
	rowN     int
	columnN  int
}

func main() {
	resultMatrix, _ := New("89 1903 3\n18 3 1\n9 4 800")
	fmt.Println(resultMatrix.elements)

	fmt.Println(resultMatrix.Rows())
	fmt.Println(resultMatrix.Cols())
	resultMatrix.Set(0, 0, 20)
	fmt.Println(resultMatrix.elements)
}

func New(s string) (*Matrix, error) {
	numOfElement := len(strings.Split(strings.Join(strings.Fields(strings.Replace(s, "\n", " ", -1)), " "), " "))
	matrix := strings.Split(s, "\n")
	numOfRows := len(matrix)
	numOfColumns := len(strings.Split(strings.TrimSpace(matrix[0]), " "))

	if numOfElement != numOfRows*numOfColumns {
		return &Matrix{}, errors.New("Invalid input")
	}

	resultMatrix := make([][]int, numOfRows)

	for index, _ := range matrix {
		splitedRow := strings.Split(strings.TrimSpace(matrix[index]), " ")
		resultMatrix[index] = make([]int, numOfColumns)

		var err error
		for j := 0; j < numOfColumns; j++ {
			resultMatrix[index][j], err = strconv.Atoi(splitedRow[j])
			if err != nil {
				return &Matrix{}, errors.New("Invalid input")
			}
		}
	}

	return &Matrix{
		elements: resultMatrix,
		rowN:     numOfRows,
		columnN:  numOfColumns,
	}, nil
}

func (m *Matrix) Cols() (resultCols [][]int) {
	resultCols = make([][]int, m.columnN)

	for j := 0; j < m.columnN; j++ {
		resultCols[j] = make([]int, m.rowN)
		for i := 0; i < m.rowN; i++ {
			resultCols[j][i] = m.elements[i][j]
		}
	}

	return resultCols
}

func (m *Matrix) Rows() (resultRows [][]int) {
	resultRows = make([][]int, m.rowN)

	for i := 0; i < m.rowN; i++ {
		resultRows[i] = make([]int, m.columnN)
		for j := 0; j < m.columnN; j++ {
			resultRows[i][j] = m.elements[i][j]
		}
	}

	return resultRows
}

func (m *Matrix) Set(row, col, val int) bool {
	if row > m.rowN-1 || col > m.columnN-1 || col < 0 || row < 0 {
		return false
	}

	m.elements[row][col] = val
	return true
}

