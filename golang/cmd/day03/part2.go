package day03

import (
	"strconv"
	"unicode"
)

func Part2(matrix [][]rune) int {
	total := 0
	for i := 0; i < len(matrix); i++ {
		for ii := 0; ii < len(matrix[i]); ii++ {
			if matrix[i][ii] == '*' {
				total += goArround(matrix, i, ii)
			}
		}
	}
	return total
}

func goArround(matrix [][]rune, row int, col int) int {
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
	touchedIndexList := make(map[string]int)
	for i := 0; i < len(directions); i++ {
		newRow := row + directions[i][0]
		newCol := col + directions[i][1]
		if newRow < 0 || newCol < 0 {
			continue
		}
		if newRow >= len(matrix) || newCol >= len(matrix[0]) {
			continue
		}
		if unicode.IsDigit(matrix[newRow][newCol]) {
			index, number := findFullNumber2(matrix, newRow, newCol)
			touchedIndexList[index] = number
		}
	}
	if len(touchedIndexList) != 2 {
		return 0
	}
	acc := 1
	for _, value := range touchedIndexList {
		acc *= value
	}
	return acc
}

func findFullNumber2(matrix [][]rune, row int, col int) (string, int) {
	initCol := col
	number := ""
	i := col
	for i < len(matrix[row]) {
		if unicode.IsDigit(matrix[row][i]) {
			number += string(matrix[row][i])
			i++
		} else {
			break
		}
	}
	i = initCol - 1
	for i >= 0 {
		if unicode.IsDigit(matrix[row][i]) {
			number = string(matrix[row][i]) + number
			initCol = i
			i--
		} else {
			break
		}
	}
	num, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(row) + "-" + strconv.Itoa(initCol), num
}
