package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(LargestSeriesProduct("73167176531330624919225119674426574742355349194934", 6))
	fmt.Println(LargestSeriesProduct("0123456789", 5))
	fmt.Println(LargestSeriesProduct("1234a5", 2))
	fmt.Println(LargestSeriesProduct("12345", -2))
	fmt.Println(LargestSeriesProduct("29", 2))
}

func maxArr(arr []int) (max int64) {
	max = int64(arr[0])

	for i := 1; i < len(arr); i++ {
		if int64(arr[i]) > max {
			max = int64(arr[i])
		}
	}

	return
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var productsArr []int

	if len(digits) < span || span < 0 {
		return 0, errors.New("Error")
	}

	for i := 0; i < len(digits)-span+1; i++ {
		productSimple := 1

		for j := 0; j < span; j++ {
			num, err := strconv.Atoi(string(digits[i+j]))

			if err != nil {
				return 0, errors.New("Error")
			}
			productSimple *= num
		}

		productsArr = append(productsArr, productSimple)
	}

	return maxArr(productsArr), nil
}
