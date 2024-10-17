package day05

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func compareSlices(slice1 []int, slice2 []int) bool {
	return reflect.DeepEqual(slice1, slice2)
}

func TestDay05PartII(t *testing.T) {
	givenExample := `seeds: 79 14 55 13

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

	chunks := strings.Split(givenExample, "\n\n")
	seedsChunk := chunks[0]
	chunks = chunks[1:]

	//Let's get all the seeds
	seeds := make([]int, 0)
	seedLines := strings.TrimPrefix(seedsChunk, "seeds: ")
	fmt.Println("seedLines", seedLines)
	// for _, line := range seedLines {
	// 	if (len(seedNumbers) % 2) != 0 {
	// 		panic(fmt.Sprintf("Error: seeds are not in pairs: %v\n", seedNumbers))
	// 	}
	// }

	actualSeeds := seeds
	expectedSeeds := []int{79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67}
	fmt.Println("actual", actualSeeds)
	fmt.Println("expected", expectedSeeds)
	compare := compareSlices(actualSeeds, expectedSeeds)
	fmt.Println("compare", compare)

}
