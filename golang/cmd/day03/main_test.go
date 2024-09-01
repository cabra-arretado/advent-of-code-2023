package main

import (
	// "strings"
	"testing"
)

func TestFindFullNumber(t *testing.T) {
	matrix := [][]rune{
		{'*', '1', '2', '3', 'b', '2'},
		{'c', '4', '5', '6', 'd', 'g'},
		{'5', 'k', 'i', 'u', '4', 'a'},
	}

	// Test case 1: Number in the middle of the row
	expected1Index := "0-1"
	expected1Num := 123
	touchesSymbol1 := true
	actual1Index, actual1Num, touches:= findFullNumber(matrix, 0, 3)
	if actual1Index != expected1Index || actual1Num != expected1Num || touchesSymbol1 != touches  {
		t.Errorf("Test case 1 failed. Expected (%s, %d), but got (%s, %d)", expected1Index, expected1Num, actual1Index, actual1Num)
	}

	// Test case 2: Number at the beginning of the row
	expected2Index := "1-1"
	expected2Num := 456
	actual2Index, actual2Num, touches2 := findFullNumber(matrix, 1, 1)
	touchesSymbol2 := true
	if actual2Index != expected2Index || actual2Num != expected2Num || touchesSymbol2 != touches2 {
		t.Errorf("Test case 2 failed. Expected (%s, %d), but got (%s, %d)", expected2Index, expected2Num, actual2Index, actual2Num)
	}

	// Test case 3: Number alone in the first column
	expected3Index := "2-0"
	expected3Num := 5
	touchesSymbol3 := false
	actual3Index, actual3Num, touches3 := findFullNumber(matrix, 2, 0)
	if actual3Index != expected3Index || actual3Num != expected3Num || touchesSymbol3 != touches3 {
		t.Errorf("Test case 3 failed. Expected (%s, %d), but got (%s, %d)", expected3Index, expected3Num, actual3Index, actual3Num)
	}

	// Test case 4: Number alone in the last column
	expected4Index := "2-4"
	expected4Num := 4
	touchesSymbol4 := false
	actual4Index, actual4Num, touches4 := findFullNumber(matrix, 2, 4)
	if actual3Index != expected3Index || actual3Num != expected3Num || touchesSymbol4 != touches4 {
		t.Errorf("Test case 4 failed. Expected (%s, %d), but got (%s, %d)", expected4Index, expected4Num, actual4Index, actual4Num)
	}
}

func TestTouchesSymbol(t *testing.T) {
	matrix := [][]rune{
		{'a', '(', '2'},
		{'c', '4', '5'},
		{'5', '*', 'i'},
		{'5', '.', 'i'},
		{'5', '.', 'i'},
	}

	// Test case 1: Symbol in the middle of the row
	// expected1 := true
	// actual1 := touchesSymbol(matrix, 0, 0)
	// if actual1 != expected1 {
	// 	t.Errorf("Test case 1 failed. Expected %t, but got %t", expected1, actual1)
	// }
	// expected2 := true
	// actual2 := touchesSymbol(matrix, 0, 0)
	// if actual2 != expected2 {
	// 	t.Errorf("Test case 2 failed. Expected %t, but got %t", expected2, actual2)
	// }
	//TODO: This is the problem
	expected3 := false
	actual3 := touchesSymbol(matrix, 4, 2)
	if actual3 != expected3 {
		t.Errorf("Test case 3 failed. Expected %t, but got %t", expected3, actual3)
	}
}

// func TestDay1(t *testing.T) {
// 	// Test case 1
// providedExample := 
// `467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..`
// 	splited := strings.Split(providedExample, "\n")
// 	matrix := make([][]rune, len(splited))
// 	for i, s := range splited {
// 		matrix[i] = []rune(s)
// 	}
// 	expected := 4361
// 	actual := part1(matrix)
//
// 	if actual != expected {
// 		t.Errorf("Test case 1 failed. Expected %d, but got %d", expected, actual)
// 	}
// }
//
