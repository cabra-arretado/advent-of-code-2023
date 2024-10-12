package day04

import (
	"strconv"
	"strings"
)

func Part2(lines []string) int {

	total := 0
	hashMap := make(map[int]*Card)
	for _, line := range lines {
		card := processCard2(line)
		hashMap[card.cardNumber] = &card
	}
	queue := []int{}
	for _, card := range hashMap {
		queue = append(queue, card.cardNumber)
	}

	instances := make(map[int]int)

	for len(queue) > 0 {
		currentCardNumber := queue[0]
		queue = queue[1:]
		instances[currentCardNumber]++
		card := hashMap[currentCardNumber]
		numMatches := card.numberOfMatches
		if numMatches > 0 {
			for i := 1; i < numMatches+1; i++ {
				nextCardNumber := currentCardNumber + i
				if _, ok := hashMap[nextCardNumber]; ok {
					queue = append(queue, nextCardNumber)
				} else {
					break
				}
			}
		}
	}
	for _, count := range instances {
		total += count
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
	card := Card{
		cardNumber:      cardNumber,
		numberOfMatches: matches,
		loteryNumbers:   lotery,
		givenNumbers:    givenNumbers,
		copies:          1,
	}
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

