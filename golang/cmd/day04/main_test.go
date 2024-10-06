package day04

import (
	"fmt"
	"strings"
	"testing"
)

// 	givenExample:= `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func Test1(t *testing.T) {
	fmt.Println("First test")

	fmt.Println(processLine("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"))
}

func BreakNumbers(s string) (slice []string, slice2 []string) {
	fullString := strings.Split(s, "|")
	rawLotery := fullString[0]
	rawGivenNumbers := fullString[1]
	lotery := gotLoterryNumbers(rawLotery)
	givenNumbers := gotGivenNumbers(rawGivenNumbers)
	return lotery, givenNumbers
}

func gotLoterryNumbers(s string) (slice []string) {
	fullString := strings.Split(s, ":")
	loteryNumbers := strings.Fields(fullString[1])
	return loteryNumbers
}
func gotGivenNumbers(s string) (slice []string) {
	givenNumbers := strings.Fields(s)
	return givenNumbers
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func processLine(rawLine string) (result int) {
	loteryNumbers, givenNumbers := BreakNumbers(rawLine)
	generalTotal := 0
	inTotal := 0
	for _, number := range loteryNumbers {
		total := 0
		if Contains(givenNumbers, number) {
			if total == 0 {
				total += 1
			} else {
				total *= 2
			}
		}
		inTotal += total
	}

	generalTotal += inTotal
	return generalTotal

}
