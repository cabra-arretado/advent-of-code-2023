package day04
//
// import (
// 	"fmt"
// 	"strings"
// )
//
// // 	givenExample:= `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// // Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
// // Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// // Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
// // Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
// // Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
//
// func Part2(lines []string) int {
// 	fmt.Println("First test")
// 	total := 0
// 	for _, line := range lines {
// 		total += processLine(line)
// 	}
// 	fmt.Println(total)
// 	return total
// }
//
// type Card struct {
// 	cardNumber    string
// 	quantity      int
// 	loteryNumbers []string
// 	givenNumbers  []string
// }
//
// func processCard(rawString string) Card {
// 	s := strings.Split(rawString, ":")
// 	cardNumber := strings.Fields(s[0])[1]
// 	fullString := strings.Split(s[1], "|")
// 	rawLotery := fullString[0]
// 	rawGivenNumbers := fullString[1]
// 	lotery := gotLoterryNumbers(rawLotery)
// 	givenNumbers := gotGivenNumbers(rawGivenNumbers)
// 	card := Card{cardNumber, 1, lotery, givenNumbers}
// 	return card
// }
//
// func Contains2(s []string, e string) bool {
// 	for _, a := range s {
// 		if a == e {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func processLine2(rawLine string) (result int) {
// 	loteryNumbers, givenNumbers := BreakNumbers(rawLine)
// 	generalTotal := 0
// 	total := 0
// 	for _, number := range loteryNumbers {
// 		if Contains(givenNumbers, number) {
// 			if total == 0 {
// 				total += 1
// 			} else {
// 				total *= 2
// 			}
// 		}
// 	}
// 	generalTotal += total
// 	return generalTotal
// }
