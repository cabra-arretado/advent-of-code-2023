package main

import (
	"fmt"
	"bufio"
	"advent-of-code-2023/utils"
)

// callbackQ1 is a callback function for the readFile function for the question 1
func callbackQ1 (scanner *bufio.Scanner) int {
	var total int
	for scanner.Scan() {
		total += 1
	}
	return total
}

func main() {
	filePath := "../inputs/day1_input.txt"

	fmt.Println(utils.ReadFile(filePath, callbackQ1))
}
