package day05

import (
	"fmt"
	"reflect"

	// "strconv"
	// "strings"
	"testing"
)

var givenExample = `seeds: 79 14 55 13

	seed-to-soil map:
	50 98 2
	52 50 48

	soil-to-fertilizer map:
	0 15 37
	37 52 2
	39 0 15

	fertilizer-to-water map:
	49 53 8
	0 11 42
	42 0 7
	57 7 4

	water-to-light map:
	88 18 7
	18 25 70

	light-to-temperature map:
	45 77 23
	81 45 19
	68 64 13

	temperature-to-humidity map:
	0 69 1
	1 0 69

	humidity-to-location map:
	60 56 37
	56 93 4`

func compareSlices(slice1 []int, slice2 []int) bool {
	return reflect.DeepEqual(slice1, slice2)
}

func TestGetSeeds(t *testing.T) {
	seeds := getSeeds("seeds: 79 14 55 13")
	expected := []SeedRange{{79, 93}, {55, 68}}
	fmt.Println("Seeds: ", seeds)
	fmt.Println("Expected: ", expected)
	fmt.Println("\033[33mPasses?\033[0m", reflect.DeepEqual(seeds, expected))
	secondExample := `seeds: 79 14 55 13
18 9 199 231`

	seeds2 := getSeeds(secondExample)
	expected2 := []SeedRange{{79, 93}, {55, 68}, {18, 27}, {199, 430}}
	fmt.Println("Seeds2: ", seeds2)
	fmt.Println("Expected2: ", expected2)
	fmt.Println("\033[33mPasses?\033[0m", reflect.DeepEqual(seeds2, expected2))
}

func TestGetInstructions(t *testing.T) {
	instructions := getInstructions("50 98 2\n52 50 48")
	expected := []Instructions{{50, 98, 2}, {52, 50, 48}}
	fmt.Println("Instructions: ", instructions)
	fmt.Println("Expected: ", expected)
	fmt.Println("Passes?", reflect.DeepEqual(instructions, expected))
}

// func TestPartII(t *testing.T) {
// 	part2 := Part2(givenExample)
// 	fmt.Println("Part II: ", part2)
// 	fmt.Println("Passes?", part2 == 46)
// }
