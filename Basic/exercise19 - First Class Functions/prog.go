package main

import (
	"errors"
	"testing"
)

type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func predicate(r Record) bool {
	return r.Day == 1
}

func main() {
	p1 := DaysPeriod{From: 1, To: 30}
	p2 := DaysPeriod{From: 31, To: 60}

	records := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	fmt.Println(CategoryExpenses(records, p1, "entertainment"))
	// => 0, error(unknown category entertainment)

	fmt.Println(CategoryExpenses(records, p1, "rent"))
	// => 1300, nil

	fmt.Println(CategoryExpenses(records, p2, "rent"))
	// => 0, nil

	fmt.Println(Filter(records, predicate))
	fmt.Println(Filter(records, ByDaysPeriod(p1)))
	fmt.Println(Filter(records, ByCategory("groceries")))
}

// Filter function to filter records according to a criteria given by a function
func Filter(records []Record, predicate func(Record) bool) []Record {
	var finalRecords []Record

	for i := range records {
		if predicate(records[i]) {
			finalRecords = append(finalRecords, records[i])
		}
	}

	return finalRecords
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		}

		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		if r.Category == c {
			return true
		}

		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var sum float64
	for i := range in {
		if in[i].Day >= p.From && in[i].Day <= p.To {
			sum += in[i].Amount
		}
	}

	return sum
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var sum float64
	var catBool bool = false
	for i := range in {

		if in[i].Category == c {
			catBool = true
			if in[i].Day >= p.From && in[i].Day <= p.To {
				sum += in[i].Amount
			}
		}
	}

	if !catBool {
		return 0, errors.New("unknown category " + c)
	}

	return sum, nil
}

