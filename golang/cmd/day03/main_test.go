package main

import (
	"testing"
)

func TestFindFullNumber(t *testing.T) {
	matrix := [][]rune{
		{'a', '1', '2', '3', 'b'},
		{'c', '4', '5', '6', 'd'},
		{'e', '7', '8', '9', 'f'},
	}

	// Test case 1: Number in the middle of the row
	expected1Index := "2-0"
	expected1Num := 123
	actual1Index, actual1Num := findFullNumber(matrix, 0, 1)
	if actual1Index != expected1Index || actual1Num != expected1Num {
		t.Errorf("Test case 1 failed. Expected (%s, %d), but got (%s, %d)", expected1Index, expected1Num, actual1Index, actual1Num)
	}

	// Test case 2: Number at the beginning of the row
	expected2Index := "1-1"
	expected2Num := 456
	actual2Index, actual2Num := findFullNumber(matrix, 1, 1)
	if actual2Index != expected2Index || actual2Num != expected2Num {
		t.Errorf("Test case 2 failed. Expected (%s, %d), but got (%s, %d)", expected2Index, expected2Num, actual2Index, actual2Num)
	}

	// Test case 3: Number at the end of the row
	expected3Index := "3-2"
	expected3Num := 789
	actual3Index, actual3Num := findFullNumber(matrix, 2, 1)
	if actual3Index != expected3Index || actual3Num != expected3Num {
		t.Errorf("Test case 3 failed. Expected (%s, %d), but got (%s, %d)", expected3Index, expected3Num, actual3Index, actual3Num)
	}
}
