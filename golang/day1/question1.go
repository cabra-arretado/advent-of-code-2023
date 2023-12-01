package main

import (
	"bufio"
	"fmt"
	"strings"
)

func processLineQ1(s string) int {
	return 0
}

func callbackQuestion1(scanner *bufio.Scanner) int {
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += processLineQ1(line)
	}
	return total
}

func answerQ1(file_path string) int {
	return readFile(file_path, callbackQuestion1)
}

func testQuestion1() {
	arr := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	testLines := strings.Split(arr, "\n")
}
