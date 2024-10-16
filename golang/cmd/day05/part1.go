package day05

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(file string) int {
	//read file

	chunks := strings.Split(file, "\n\n")
	seedsChunk := chunks[0]
	chunks = chunks[1:]

	//Let's get all the seeds
	seeds := make([]int, 0)
	seedLines := strings.Trim(seedsChunk, "seeds: ")
	for _, line := range strings.Split(seedLines, "\n") {
		for _, seed := range strings.Split(line, " ") {
			seedNumber, err := strconv.Atoi(seed)
			if err != nil {
				panic(fmt.Sprintf("Error converting seed to int: %v\n", seed))
			}
			seeds = append(seeds, seedNumber)
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
	minValue := seeds[0]
	for _, seed := range seeds {
		if seed < minValue {
			minValue = seed
		}

	}
	return minValue

}
