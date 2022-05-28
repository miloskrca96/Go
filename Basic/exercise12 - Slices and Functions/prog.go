package main

import "fmt"

func main() {
	cards := FavoriteCards()
	fmt.Println(cards)
	fmt.Println(GetItem([]int{1, 2, 4, 1}, 2))
	fmt.Println(GetItem([]int{1, 2, 4, 1}, 22))

	index := -1
	newCard := 6
	cardsSetItem := SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(cardsSetItem)

	slice := []int{3, 2, 6, 4, 8}
	cardsPrepend := PrependItems(slice)
	fmt.Println(cardsPrepend)

	cardsNew := RemoveItem([]int{3, 2, 6, 4, 8}, 4)
	fmt.Println(cardsNew)
}

// Function that return slice with 2, 6, 9 cards
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// Function that return card at index position
func GetItem(slice []int, index int) int {
	if index < 0 || index > len(slice)-1 {
		return -1
	}

	return slice[index]
}

// Exchange the card at position index with the new card
// provided and return the adjusted stack
func SetItem(slice []int, index, newCard int) []int {
	if index < 0 || index > len(slice)-1 {
		slice = append(slice, newCard)
	} else {
		slice[index] = newCard
	}

	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, items ...int) []int {
	for i := len(items) - 1; i >= 0; i-- {
		slice = append([]int{items[i]}, slice...)
	}

	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= 0 && index < len(slice) {
		slice = append(slice[:index], slice[index+1:]...)
	}

	return slice
}

