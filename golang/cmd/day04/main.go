package day04

import (
	"advent-of-code-2023/utils"
	"fmt"
)

func Day03() {
	matrix := utils.ReadFileAsMatrix("3")

	// ---- Part 1 ----
	result := Part1(matrix)
	fmt.Println(result)
	// fmt.Println("pass?", result == 544664)

	// ---- Part 2 ----
	result = Part2(matrix)
	fmt.Println(result)
	// fmt.Println("pass?", result == 84495585)
}
