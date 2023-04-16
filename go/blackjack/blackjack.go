package blackjack

// Player decisions
const (
	stand = "S"
	win   = "W"
	split = "P"
	hit   = "H"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
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
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sumCards := sumCardsValues(card1, card2)

	switch {
	case isAce(card1) && isAce(card2):
		return split
	case sumCards == 21:
		if isAce(dealerCard) || ParseCard(dealerCard) == 10 {
			return stand
		} else {
			return win
		}
	case sumCards >= 17 && sumCards <= 20:
		return stand
	case sumCards >= 12 && sumCards <= 16:
		if ParseCard(dealerCard) >= 7 {
			return hit
		} else {
			return stand
		}
	default:
		return hit
	}
}

// sumCardsValues returns the sum of two given cards values
func sumCardsValues(card1, card2 string) int {
	return ParseCard(card1) + ParseCard(card2)
}

// isAce determine if a given card is an Ace or not 
func isAce(card string) bool {
	return card == "ace"
}
