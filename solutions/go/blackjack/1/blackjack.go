package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	result:= 0
    switch {
        case card == "ace":
        result = 11
           case card == "two":
        result = 2
           case card == "three":
        result = 3
           case card == "four":
        result = 4
           case card == "five":
        result = 5
           case card == "six":
        result = 6
           case card == "seven":
        result = 7
           case card == "eight":
        result = 8
           case card == "nine":
        result = 9
 		case card == "ten", card == "jack", card == "queen", card == "king":
    	result = 10
        default:
    	result = 0
    }
    return result
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    scorePlayer := ParseCard(card1) + ParseCard(card2)
    result := ""

    if card1 == "ace" && card2 == "ace" {
        result = "P"
    } else if scorePlayer == 21 {
		if ParseCard(dealerCard) == 10 || ParseCard(dealerCard) == 11 {
            result = "S"
        }else{
            result = "W"
        }
    } else if scorePlayer >= 17 && scorePlayer <= 20{
        result = "S"
    } else if scorePlayer >= 12 && scorePlayer <= 16 {
    	if ParseCard(dealerCard) >= 7 {
            result = "H"
        } else {
    		result = "S"}
    }else {
        result = "H"
            }


    return result
}
