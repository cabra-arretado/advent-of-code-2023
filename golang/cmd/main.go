package main

import (
	"advent-of-code-2023/cmd/day03"
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
	default:
		fmt.Println("Day not found")
	}
}
