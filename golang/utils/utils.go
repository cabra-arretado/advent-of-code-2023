package utils

import (
	"bufio"
	"os"
)

func readFile(file_path string, callback func(*bufio.Scanner) int) int {
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return callback(scanner)
}
