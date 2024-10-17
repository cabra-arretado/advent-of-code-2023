package day05

import (
	"fmt"
	"testing"
	"reflect"
	"strings"
	"strconv"
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
	seedLines := strings.Trim(seedsChunk, "seeds: ")
	duo := make([]int, 0)
	for _, line := range strings.Split(seedLines, "\n") {
		for _, seed := range strings.Split(line, " ") {
			seedNumber, err := strconv.Atoi(seed)
			if err != nil {
				panic(fmt.Sprintf("Error converting seed to int: %v\n", seed))
			}
			if len(duo) < 2 {
				duo = append(duo, seedNumber)
			} else {
				start := duo[0]
				length := duo[1] - duo[0] + 1
				for i := start; i <= length; i++ {
					seeds = append(seeds, i)
				}
				duo = make([]int, 0)
			}
		}
	}
	actualSeeds := seeds
	expected := []int{79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67}
	fmt.Println("actual", actualSeeds)
	fmt.Println("expected", expected)
	compare := compareSlices(actualSeeds, expected)
	fmt.Println("compare", compare)

}
