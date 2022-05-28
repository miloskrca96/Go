package main

import (
	"fmt"
)

func main() {
	units := Units()
	fmt.Println(units)

	bill := NewBill()
	fmt.Println(bill)

	okAdd := AddItem(bill, units, "carrot", "dozen")
	fmt.Println(okAdd)
	fmt.Println(bill)

	okRemove := RemoveItem(bill, units, "carrot", "quarter_of_a_dozen")
	fmt.Println(okRemove)
	fmt.Println(bill)

	newBill := map[string]int{"carrot": 12, "grapes": 3}
	qty, ok := GetItem(newBill, "grapes")
	fmt.Println(qty)
	fmt.Println(ok)
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, ok := units[unit]
	if ok {
		bill[item] += units[unit]
		return true
	}

	return false

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, okItem := bill[item]
	_, okUnit := units[unit]

	if !okItem || !okUnit || bill[item]-units[unit] < 0 {
		return false
	}

	bill[item] -= units[unit]
	if bill[item] == 0 {
		delete(bill, item)
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, okItem := bill[item]
	if !okItem {
		return 0, false
	}

	return bill[item], true
}

