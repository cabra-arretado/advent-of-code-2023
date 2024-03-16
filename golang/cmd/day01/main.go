package main

import (
	"advent-of-code-2023/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// callbackQ1 is a callback function for the readFile function for the question 1
func callbackQ1(scanner *bufio.Scanner) int {
	var total int
	for scanner.Scan() {
		line := scanner.Text()
		total += processLine1(line)
	}
	return total
}

func processLine1(line string) int {
	slice := []rune(line)
	var total []rune
	for _, char := range slice {
		if unicode.IsDigit(char) {
			total = append(total, char)
			break
		}
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if unicode.IsDigit(slice[i]) {
			total = append(total, slice[i])
			break
		}
	}
	totalString := string(total)
	totalInt, err := strconv.Atoi(totalString)
	if err != nil {
		panic(err)
	}
	return totalInt
}

func callbackQ2(scanner *bufio.Scanner) int {
	var total int
	for scanner.Scan() {
		line := scanner.Text()
		line = substituteNumbers(line)
		total += processLine1(line)
	}
	return total
}

var numSubs = map[string]string{
	"zero":  "zero0zero",
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

func substituteNumbers(s string) string {
	for old, new := range numSubs {
		s = strings.ReplaceAll(s, old, new)
	}
	return s
}

func main() {
	filePath := "../inputs/day1_input.txt"

	fmt.Println("Day 1")
	fmt.Println("Question 1", utils.ReadFile(filePath, callbackQ1))
	fmt.Println("Question 2", utils.ReadFile(filePath, callbackQ2))
}
