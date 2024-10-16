package day05

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

//TODO: We just have to change the input conssuption

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
				for i := duo[0]; i <= duo[1]; i++ {
					seeds = append(seeds, i)
				}
				duo = make([]int, 0)
			}
		}
		fmt.Println(seeds)
		// -- END get all seeds

		for _, chunk := range chunks {
			// ignore first line
			maps := make([][]int, 0)
			lines := strings.Split(chunk, "\n")[1:]
			for _, line := range lines {
				var destRange, sourceRange, rangeLength string
				fmt.Sscanf(line, "%s %s %s", &destRange, &sourceRange, &rangeLength)
				destRangeInt, _ := strconv.Atoi(destRange)
				sourceRangeInt, _ := strconv.Atoi(sourceRange)
				rangeLengthInt, _ := strconv.Atoi(rangeLength)
				maps = append(maps, []int{destRangeInt, sourceRangeInt, rangeLengthInt})
			}
			newSeeds := make([]int, len(seeds))
			for i, seed := range seeds {
				for _, m := range maps {
					if seed >= m[1] && seed <= m[1]+m[2] {
						mappedValue := m[0] + (seed - m[1])
						newSeeds[i] = mappedValue
						break
					} else {
						newSeeds[i] = seed
					}
				}
				seeds = newSeeds
			}
		}
		minValue := 0
		for _, seed := range seeds {
			if seed < minValue {
				minValue = seed
			}

		}
		fmt.Println("Part II:", minValue, minValue == 46)

	}
}
