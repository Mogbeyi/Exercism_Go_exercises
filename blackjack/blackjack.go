package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}

	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	cardSum := card1Value + card2Value

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case cardSum == 21:
		if dealerCard != "ace"  || dealerCardValue != 0 || dealerCard != "ten"{
			return "W"
		}
		return "S"
	case cardSum == 17 || cardSum == 18 || cardSum == 19 || cardSum == 20:
		return "S"
	case cardSum == 12 || cardSum == 13 || cardSum == 14 || cardSum == 15 || cardSum == 16:
		if dealerCardValue >= 7 {
			return "H"
		}
		return "S"
	case cardSum <= 11:
		return "H"
	}
	return ""
}
