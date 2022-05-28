package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Welcome("Milos"))
	fmt.Println(HappyBirthday("Milos", 25))
	fmt.Println(AssignTable("Christiane", 100, "Frank", "on the left", 23.7834298))
}

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	if age < 0 {
		return "incorrect age number"
	}
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var message string
	var tableNum string = strconv.Itoa(table)

	if table < 10 {
		tableNum = "00" + tableNum
	} else if table > 10 && table < 100 {
		tableNum = "0" + tableNum
	}

	message = fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %s. "+
		"Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.",
		name, tableNum, direction, distance, neighbor)
	return message
}
