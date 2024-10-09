package day04

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type Card struct {
	quantity      int
	loteryNumbers []string
	givenNumbers  []string
}

func Test2(t *testing.T) {
	givenExample := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

}

func Part2(lines []string) int {
	cards := make(map[int]Card)
	total := 0
	// Initial processing
	for _, line := range lines {
		cardNumber, card := processCard(line)
		cards[cardNumber] = card
	}
	// Processing the cards per si
	// for cardNumber, card := range cards {
	// 	// 1-Countains2
	// 	// 2-AddMoreCards
	//
	// }
	fmt.Println(total)
	return total
}
func addMoreCards(cards map[int]Card, initialCardNumber int, quantity int) {
	for i := 1; i < quantity; i++ {
		card, exists := cards[i+initialCardNumber]
		if exists {
			card.quantity += 1
		} else {
			panic("Card does not exist")
		}
	}

}

func processCard(rawString string) (number int, c Card) {
	s := strings.Split(rawString, ":")
	cardNumberString := strings.Fields(s[0])[1]
	cardNumber, err := strconv.Atoi(cardNumberString)
	if err != nil {
		fmt.Println("Could not convert card number to int")
	}
	fullString := strings.Split(s[1], "|")
	rawLotery := fullString[0]
	rawGivenNumbers := fullString[1]
	lotery := gotLoterryNumbers(rawLotery)
	givenNumbers := gotGivenNumbers(rawGivenNumbers)
	card := Card{1, lotery, givenNumbers}
	return cardNumber, card
}

func Contains2(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func processLine2(loteryNumbers []string, givenNumbers []string) (result int) {
	generalTotal := 0
	total := 0
	for _, number := range loteryNumbers {
		if Contains2(givenNumbers, number) {
			total += 1
		}
	}
	generalTotal += total
	return generalTotal
}
