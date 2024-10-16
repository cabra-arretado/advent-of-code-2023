package day04

import (
	"advent-of-code-2023/utils"
	"fmt"
)

func Day04() {
	slice := utils.ReadFileAsSlice("4")
	fmt.Println("Part1")

	// ---- Part 1 ----
	result := Part1(slice)
	fmt.Println("Part I:", result)

	// // ---- Part 2 ----
	// result = Part2(slice)
	// fmt.Println("Part II:", result)
}
