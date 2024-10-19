package day05

import (
	"fmt"
	"strconv"
	"strings"
)

type SeedRange struct {
	start int
	range   int
}

func getSeeds(seedsChunk string) []int {
	seeds := make([]int, 0)
	seedLines := strings.TrimPrefix(seedsChunk, "seeds: ")
	for _, line := range strings.Split(seedLines, "\n") {
		lineTrimed := strings.TrimSpace(line)
		lineSplitted := strings.Split(lineTrimed, " ")
		duo := make([]int,0, 2)
		for _, seed := range lineSplitted {
			seedInt, err := strconv.Atoi(seed)
			if err != nil {
				fmt.Println("error converting seed to int")
			}
			duo = append(duo, seedInt)
			if len(duo) == 2 {
				for i := duo[0]; i <= duo[0]+duo[1]-1; i++ {
					seeds = append(seeds, i)
				}
				duo = make([]int,0, 2)
			}
		}
	}
	return seeds
}
func Part2(file string) int {
	fmt.Println("Part II: Starting ")
	chunks := strings.Split(file, "\n\n")
	seedsChunk := chunks[0]
	chunks = chunks[1:]
	seeds := getSeeds(seedsChunk)

	fmt.Println("line 40")

	//-- END get all seeds
	for _, chunk := range chunks {
		// ignore first line
		maps := make([][]int, 0)
		lines := strings.Split(chunk, "\n")[1:]
		for _, line := range lines {
			fmt.Println("inside 48")
			var destRange, sourceRange, rangeLength string
			fmt.Sscanf(line, "%s %s %s", &destRange, &sourceRange, &rangeLength)
			destRangeInt, _ := strconv.Atoi(destRange)
			sourceRangeInt, _ := strconv.Atoi(sourceRange)
			rangeLengthInt, _ := strconv.Atoi(rangeLength)
			maps = append(maps, []int{destRangeInt, sourceRangeInt, rangeLengthInt})
		}
		newSeeds := make([]int, len(seeds))
		for i, seed := range seeds {
			fmt.Println("inside line 58")
			for _, m := range maps {
				fmt.Println("inside map", m)
				if seed >= m[1] && seed <= m[1]+m[2]-1 {
					mappedValue := m[0] + (seed - m[1])
					newSeeds[i] = mappedValue
					break
				} else {
					newSeeds[i] = seed
				}
			}
		}
		seeds = newSeeds
	}
	minValue := seeds[0]
	for _, seed := range seeds {
		if seed < minValue {
			minValue = seed
		}

	}
	return minValue

}
