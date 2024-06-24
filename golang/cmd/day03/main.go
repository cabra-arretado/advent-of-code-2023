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

func touchesSymbol(matrix [][]rune, row int, col int) bool {
	directions := [][]int{
		{-1, -1}, // up-left
		{-1, 0},  // up
		{-1, 1},  // up-right
		{1, 0},   // down
		{1, 1},   // down-right
		{1, -1},  // down-left
		{0, -1},  // left
		{0, 1},   // right
	}
	for i := 0; i < len(directions); i++ {
		newRow := row + directions[i][0]
		newCol := col + directions[i][1]
		if unicode.IsDigit(matrix[newRow][newCol]) || matrix[newRow][newCol] == '.' {
			continue
		}
		if unicode.IsSymbol(matrix[newRow][newCol]) {
			return true
		}
	}
	return false
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
	return strconv.Itoa(row) + "-" + strconv.Itoa(initIndex), num
}

func main() {
	matrix := utils.ReadFileAsMatrix("3")
	fmt.Println(rune(matrix[0][0]))

}
