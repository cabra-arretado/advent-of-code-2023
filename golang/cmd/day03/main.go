package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"unicode"
)

func part1(matrix [][]rune) int {
	return 0
}

// findFullNumber finds the full number in the matrix
func findFullNumber(matrix [][]rune, row int, col int) (string, int) {
	initIndex := col
	number := ""
	i := col + 1
	for i < len(matrix) {
		if unicode.IsDigit(matrix[row][i]) {
			number += string(matrix[row][i])
			col++
		} else {
			break
		}
	}
	for col > 0 {
		if unicode.IsDigit(matrix[row][col]) {
			number = string(matrix[row][col]) + number
			initIndex = col
			col--
		} else {
			break
		}
	}
	num, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(initIndex) + "-" + strconv.Itoa(row), num
}

func main() {
	matrix := utils.ReadFileAsMatrix("3")
	fmt.Println(rune(matrix[0][0]))
}
