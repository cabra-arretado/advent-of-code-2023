package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"unicode"
)

func part1(matrix [][]rune) int {
	//TODO
	return 0
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) || r != '.'
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
		if newRow < 0 || newCol < 0 {
			continue
		}
		if newRow > len(matrix) || newCol > len(matrix[row]) {
			continue
		}
		if isSymbol(matrix[newRow][newCol]) {
			return true
		}
	}
	return false
}

// findFullNumber finds the full number in the matrix
// returns tuple: (row-col of the first digit, entire number, and it touches a symbol bool)
func findFullNumber(matrix [][]rune, row int, col int) (string, int, bool) {
	touched := false
	initIndex := col
	number := "" + string(matrix[row][col])
	i := col + 1
	for i < len(matrix[row]) {
		if unicode.IsDigit(matrix[row][i]) {
			touched = touchesSymbol(matrix, row, i)
			number += string(matrix[row][i])
			i++
		} else {
			break
		}
	}
	i = initIndex - 1
	for i >= 0 {
		if unicode.IsDigit(matrix[row][i]) {
			touched = touchesSymbol(matrix, row, i)
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
	return strconv.Itoa(row) + "-" + strconv.Itoa(initIndex), num, touched
}

func main() {
	//TODO
	matrix := utils.ReadFileAsMatrix("3")
	fmt.Println(rune(matrix[0][0]))
}
