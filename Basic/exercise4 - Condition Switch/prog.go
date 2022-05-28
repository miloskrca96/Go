package main

import "fmt"

var cards map[string]int = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"king":  10,
	"queen": 10,
}

func main() {
	value := ParseCard("two")
	fmt.Println(value)
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	val, ok := cards[card]
	if ok {
		return val
	} else {
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	valCard1, _ := cards[card1]
	valCard2, _ := cards[card2]
	valDealerCard, _ := cards[dealerCard]

	sumOfCards := valCard1 + valCard2

	switch {
	case sumOfCards == 22:
		return "P"
	case sumOfCards == 21 && (valDealerCard == 11 || valDealerCard == 10):
		return "S"
	case sumOfCards == 21 && (valDealerCard != 11 || valDealerCard != 10):
		return "W"
	case sumOfCards >= 17 && sumOfCards <= 20:
		return "S"
	case sumOfCards >= 12 && sumOfCards <= 16 && valDealerCard >= 7:
		return "H"
	case sumOfCards >= 12 && sumOfCards <= 16 && valDealerCard < 7:
		return "S"
	default:
		return "H"
	}
}
