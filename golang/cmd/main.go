package main

import (
	"advent-of-code-2023/cmd/day03"
	"advent-of-code-2023/cmd/day04"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the day: (e.g. 01, 10, 20) ")
	input, _ := reader.ReadString('\n') // Read input until newline
	input = strings.TrimSpace(input)
	switch input {
	case "03":
		day03.Day03()
	case "04":
		day04.Day04()
	case "05":
		day04.Day04()
	default:
		fmt.Println("Day not found")
	}
}
