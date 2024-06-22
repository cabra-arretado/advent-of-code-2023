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
// returns tuple: (row-col of the first digit, number)
func findFullNumber(matrix [][]rune, row int, col int) (string, int) {
	initIndex := col
	number := "" + string(matrix[row][col])
	i := col + 1
	for i < len(matrix[row]) {
		if unicode.IsDigit(matrix[row][i]) {
			number += string(matrix[row][i])
			i++
		} else {
			break
		}
	}
	i = initIndex - 1
	for i >= 0 {
		if unicode.IsDigit(matrix[row][i]) {
			number = string(matrix[row][i]) + number
			initIndex = i
			i--
		} else {
			break
		}
	}
	num, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(row)+ "-" + strconv.Itoa(initIndex) , num
}

func main() {
	matrix := utils.ReadFileAsMatrix("3")
	fmt.Println(rune(matrix[0][0]))
}
