package utils

import (
	"bufio"
	"fmt"
	"os"
)

const baseFilePath string = "../inputs/day"

// ReadFile reads a file and returns a scanner object
func ReadFile(file_path string, callback func(*bufio.Scanner) int) int {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return callback(scanner)
}

// ReadFileAsMatrix reads a file and returns a matrix of runes
func ReadFileAsMatrix(dayNumber string) [][]rune {
	base := baseFilePath + dayNumber + "_input.txt"
	fmt.Println(base)
	file, err := os.Open(base)
	if err != nil {
		fmt.Println(base)
		panic(err)
	}
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		matrix = append(matrix, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return matrix
}

func PrintHelloWorld() {
	fmt.Println("Hello World")
}
