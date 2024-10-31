package day05

import (
	"fmt"
	"strconv"
	"strings"
)

type SeedRange struct {
	start int
	end   int
}

type Instructions struct {
	destinationStart int
	sourceStart      int
	sourceEnd        int
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
				seedRange := SeedRange{start: duo[0], end: duo[0] + duo[1]}
				duo = make([]int, 0, 2)
				seeds = append(seeds, seedRange)
			}
		}
	}
	return seeds
}

func getInstructions(chunk string) []Instructions {
	chunckLines := strings.Split(chunk, "\n")
	chunckLines = chunckLines[1:]
	instructions := make([]Instructions, 0)
	for _, line := range chunckLines {
		lineTrimed := strings.TrimSpace(line)
		lineSplitted := strings.Split(lineTrimed, " ")
		destinationStart, err := strconv.Atoi(lineSplitted[0])
		if err != nil {
			fmt.Println("error converting destinationStart to int")
		}
		sourceStart, err := strconv.Atoi(lineSplitted[1])
		if err != nil {
			fmt.Println("error converting sourceStart to int")
		}
		rangeLength, err := strconv.Atoi(lineSplitted[2])
		if err != nil {
			fmt.Println("error converting rangeLength to int")
		}
instructions = append(instructions, Instructions{destinationStart: destinationStart, sourceStart: sourceStart, rangeLength: rangeLength, sourceEnd: sourceStart + rangeLength})
	}
	return instructions
}

func Part2(file string) int {
	// chunks := strings.Split(file, "\n\n")
	// seeds := getSeeds(chunks[0])
	// chunks = chunks[1:]
	// newSeeds := make([]SeedRange, 0)
	// for _, chunk := range chunks {
	// 	for _, seedRange := range seeds {
	// 		for _, instruction := range getInstructions(chunk) {
	// 		}
	// 	}
	// }
	return 0
}
