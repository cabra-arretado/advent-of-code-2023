package day05

import (
	"fmt"
	"strconv"
	"strings"
)

type SeedRange struct {
	seedStart int
	length    int
}

type Instructions struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

func getSeeds(seedsChunk string) []SeedRange {
	seeds := make([]SeedRange, 0)
	seedLines := strings.TrimPrefix(seedsChunk, "seeds: ")
	for _, line := range strings.Split(seedLines, "\n") {
		lineTrimed := strings.TrimSpace(line)
		lineSplitted := strings.Split(lineTrimed, " ")
		duo := make([]int, 0, 2)
		for _, seed := range lineSplitted {
			seedInt, err := strconv.Atoi(seed)
			if err != nil {
				fmt.Println("error converting seed to int")
			}
			duo = append(duo, seedInt)
			if len(duo) == 2 {
				seedRange := SeedRange{seedStart: duo[0], length: duo[1]}
				duo = make([]int, 0, 2)
				seeds = append(seeds, seedRange)
			}
		}
	}
	return seeds
}

func Part2(file string) int {

	return 0
}
