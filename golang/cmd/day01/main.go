package main

// In this one there are two different ways to consume the file.
// The first one is using a callback function and the second one is using a matrix of runes.

import (
	"advent-of-code-2023/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
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

func processLine1Matrix(slice []rune) int {
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

func matrixQ1(matrix [][]rune) int {
	var total int
	for _, slice := range matrix {
		total += processLine1Matrix(slice)
	}
	return total
}
func matrixQ2(matrix [][]rune) int {
	var total int
	for _, slice := range matrix {
		line := string(slice)
		line = substituteNumbers(line)
		total += processLine1(line)
	}
	return total
}

func main() {
	filePath := "../inputs/day1_input.txt"
	start1 := time.Now()
	resultQ1 := utils.ReadFile(filePath, callbackQ1)

	resultQ2 := utils.ReadFile(filePath, callbackQ2)
	elapsed1 := time.Since(start1)
	fmt.Println("Question 1:", resultQ1)
	fmt.Println("Question 2:", resultQ2)
	fmt.Println("Execution time with Callback:", elapsed1)

	fmt.Println(("-------------------"))

	matrix := utils.ReadFileAsMatrix(filePath)
	start2 := time.Now()
	resultQ1Matrix := matrixQ1(matrix)
	resultQ2Matrix := matrixQ2(matrix)
	elapsed2 := time.Since(start2)

	fmt.Println("Question 1:", resultQ1Matrix)
	fmt.Println("Question 2:", resultQ2Matrix)
	fmt.Println("Execution time with Matrix:", elapsed2)
}
