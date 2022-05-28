package main

import (
	"fmt"
)

const (
	defaultTime int     = 2
	noodleLayer int     = 50
	sauceLayer  float64 = 0.2
)

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTime(layers, 3))
	fmt.Println(PreparationTime(layers, 0))

	fmt.Println(Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}))

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}

	AddSecretIngredient(friendsList, myList)
	fmt.Println(myList)

	quantities := []float64{1.2, 3.6, 10.5}
	scaledQuantities := ScaleRecipe(quantities, 4)
	fmt.Println(scaledQuantities)
}

//PreparationTime function accepts a slice of layers and the average preparation time per layer in minutes.
//The function should return the estimate for the total preparation time based on the number of layers
func PreparationTime(layers []string, time int) int {
	if time <= 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

//Quantities function takes a slice of layers.
//The function will then determine the quantity of noodles and sauce needed to make your meal
func Quantities(layers []string) (int, float64) {
	var expNoodles int
	var expSauce float64

	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			expNoodles += noodleLayer
		case "sauce":
			expSauce += sauceLayer
		}
	}
	return expNoodles, expSauce
}

//AddSecretIngredient function accepts two slices of ingredients.
//The first parameter is the list your friend sent you, the second is the ingredient list of your own recipe
//The last element in your ingredient list is always "?".
//The function should replace it with the last item from your friends list
func AddSecretIngredient(list1, list2 []string) {
	list2[len(list2)-1] = list1[len(list1)-1]
}

//ScaleRecipe function takes two parameters.A slice of amounts needed for 2 portions.
//and the number of portions you want to cook.
func ScaleRecipe(sliceAmounts []float64, numberOfPortion int) []float64 {
	var scaledQuantities []float64
	for i := 0; i < len(sliceAmounts); i++ {
		var amountNeeded = sliceAmounts[i] * float64(numberOfPortion) / 2
		scaledQuantities = append(scaledQuantities, amountNeeded)
	}

	return scaledQuantities
}
