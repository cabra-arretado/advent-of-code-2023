package utils

import (
	"fmt"
	"bufio"
	"os"
)

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

func PrintHelloWorld() {
	fmt.Println("Hello World")
}
