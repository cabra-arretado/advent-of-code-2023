package day04

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test2(t *testing.T) {
	givenExample := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	fmt.Println("PartII: ", Part2(strings.Split(givenExample, "\n")))
}

func Part2(lines []string) int {
	total := 0
	hashMap := make(map[int]Card)
	queue := []Card{}
	for _, line := range lines {
		card := processCard2(line)
		hashMap[card.cardNumber] = card
		queue = append(queue, card)
	}
	for _, ea := range queue {
		if ea.numberOfMatches > 0 {
			addCardsToFollowingCard(queue, ea, hashMap)
		}
	}

	for _, ea := range hashMap {
		total += ea.copies
	}

	return total
}

type Card struct {
	cardNumber      int
	numberOfMatches int
	loteryNumbers   []string
	givenNumbers    []string
	copies          int
}

func processCard2(rawString string) Card {
	s := strings.Split(rawString, ":")
	cardNumber, err := strconv.Atoi(strings.Fields(s[0])[1])
	if err != nil {
		panic("Could not convert number")
	}
	fullString := strings.Split(s[1], "|")
	rawLotery := fullString[0]
	rawGivenNumbers := fullString[1]
	lotery := getNumbers(rawLotery)
	givenNumbers := getNumbers(rawGivenNumbers)
	matches := getMatchNumber(lotery, givenNumbers)
	// fmt.Println("Card: ", cardNumber, "Matches: ", matches)
	card := Card{cardNumber, matches, lotery, givenNumbers, 1}
	return card
}

func getNumbers(s string) []string {
	// That was bugging when using the first one
	s = strings.Trim(s, " ")
	return strings.Fields(s)
}

func getMatchNumber(loteryNumbers []string, givenNumbers []string) int {
	count := 0
	for _, lotery := range loteryNumbers {
		for _, given := range givenNumbers {
			if lotery == given {
				count++
			}
		}
	}
	return count
}

func addCardsToFollowingCard(queue []Card, currentCard Card, instances map[int]Card) {
	if currentCard.numberOfMatches <= 0 {
		return
	}
	for i := 1; i < currentCard.numberOfMatches+1; i++ {
		nextCardNumber := currentCard.cardNumber + i
		if _, ok := instances[nextCardNumber]; ok {
			nextCard := instances[nextCardNumber]
			nextCard.copies++
			queue = append(queue, nextCard)
		}
	}
}
