package main

import "fmt"

type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

func main() {
	name := "Matthew Sanabria"
	age := 29
	address := map[string]string{"street": "Main St."}

	resident := NewResident(name, age, address)

	fmt.Println(resident)

	resident.Delete()

	fmt.Println(resident)

	name1 := "Matthew Sanabria"
	age1 := 29
	address1 := map[string]string{"street": "Main St."}

	resident1 := NewResident(name1, age1, address1)

	name2 := "Rob Pike"
	age2 := 0
	address2 := make(map[string]string)

	resident2 := NewResident(name2, age2, address2)

	residents := []*Resident{resident1, resident2}

	fmt.Println(Count(residents))
}

// NewResident registers a new resident in this city.
func NewResident(Name string, Age int, Address map[string]string) *Resident {
	return &Resident{
		Name:    Name,
		Age:     Age,
		Address: Address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (resident *Resident) HasRequiredInfo() bool {
	var Name = resident.Name
	var val, ok = resident.Address["street"]

	if Name != "" && ok && val != "" {
		return true
	}

	return false
}

// Delete deletes a resident's information.
func (resident *Resident) Delete() {
	resident.Name = ""
	resident.Age = 0
	resident.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var sum int
	for i := range residents {
		if residents[i].HasRequiredInfo() {
			sum++
		}
	}

	return sum
}
