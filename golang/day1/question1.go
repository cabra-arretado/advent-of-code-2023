package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func findFirstDigit(r []rune, dir int) int {
	var init_idx int
	if dir < 0 {
		init_idx = len(r)
		for i := init_idx; 0 > i; i-- {
			fmt.Println(r[i])
			if unicode.IsNumber(r[i]) {
				return convertRunetoInt(r[i])
			}
		}
	} else {
		init_idx = 0
		for i := init_idx; i < len(r); i++ {
			fmt.Println(r[i])
			if unicode.IsNumber(r[i]) {
				return convertRunetoInt(r[i])
			}
		}
	}
	return -1
}

func convertRunetoInt(r rune) int {
	num, err := strconv.Atoi(string(r))
	if err != nil {
		panic("cannot convert runer")
	}
	fmt.Println("convert", num)
	return num
}

func processLineQ1(s string) int {
	var total int
	r := []rune(s)
	total += findFirstDigit(r, 1)
	fmt.Println("total", total)
	total *= 10
	total += findFirstDigit(r, -1)
	fmt.Println("total", total)
	return total
}

func callbackQuestion1(scanner *bufio.Scanner) int {
	total := 0
	scanner.Split(bufio.ScanLines)

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
	arr := `1abc2`
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet`
	testLines := strings.Split(arr, "\n")
	for i := range testLines {
		fmt.Println("line", testLines[i])
		fmt.Println("answer", processLineQ1(testLines[i]))
		
	}
	// expected := []int{12, 38, 15, 77}
	// i := 1
	// print(testLines[i])
	// fmt.Println("Test", i, processLineQ1(testLines[i]) == expected[i], processLineQ1(testLines[i]))
}
