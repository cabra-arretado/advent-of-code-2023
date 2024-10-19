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
	fmt.Println("Part II: Starting ")
	chunks := strings.Split(file, "\n\n")
	seedsChunk := chunks[0]
	chunks = chunks[1:]
	seeds := getSeeds(seedsChunk)
	fmt.Println("finish processing the seeds")
	//-- END get all seeds
	for _, chunk := range chunks {
		// ignore first line
		processedInstructions := make([]Instructions, 0)
		lines := strings.Split(chunk, "\n")[1:]
		for _, line := range lines {
			var destRange, sourceRange, rangeLength string
			fmt.Sscanf(line, "%s %s %s", &destRange, &sourceRange, &rangeLength)
			destRangeInt, err := strconv.Atoi(destRange)
			if err != nil {
				fmt.Println("error converting destRange to int")
			}
			sourceRangeInt, _ := strconv.Atoi(sourceRange)
			if err != nil {
				fmt.Println("error converting sourceRange to int")
			}
			rangeLengthInt, _ := strconv.Atoi(rangeLength)
			if err != nil {
				fmt.Println("error converting rangeLength")
			}
			processedInstructions = append(processedInstructions, Instructions{destRangeInt, sourceRangeInt, rangeLengthInt})
		}
		fmt.Println("finish processing the instructions")
		newSeeds := make([]SeedRange, len(seeds))
		for _, seed := range seeds {
			for _, instruction := range processedInstructions {
				// Range all out of the map
				if seed.seedStart+seed.length < instruction.sourceStart || seed.seedStart > instruction.sourceStart+instruction.rangeLength-1 {
					newSeeds = append(newSeeds, seed)
				}
				// Range all in the map
				if seed.seedStart >= instruction.sourceStart && seed.seedStart+seed.length <= instruction.sourceStart+instruction.rangeLength-1 {
					newSeeds = append(newSeeds, SeedRange{seedStart: instruction.destinationStart + (seed.seedStart - instruction.sourceStart), length: seed.length})
				}
				// TODO : Seed range starts out but ends in the map
				if seed.seedStart < instruction.sourceStart && seed.seedStart+seed.length > instruction.sourceStart+instruction.rangeLength-1 {
					//to add the first part without the mapping
					newSeeds = append(newSeeds, SeedRange{seedStart: seed.seedStart, length: instruction.sourceStart - seed.seedStart})
					//to add the second part with the mapping
					newSeeds = append(newSeeds, SeedRange{seedStart: instruction.destinationStart, length: instruction.rangeLength})
				}
				// TODO : Seed range starts in the map but ends out
				if seed.seedStart >= instruction.sourceStart && seed.seedStart+seed.length > instruction.sourceStart+instruction.rangeLength-1 {
					//to add the first part with the mapping
					newSeeds = append(newSeeds, SeedRange{seedStart: instruction.destinationStart + (seed.seedStart - instruction.sourceStart), length: instruction.rangeLength - (seed.seedStart - instruction.sourceStart)})
					//to add the second part without the mapping
					newSeeds = append(newSeeds, SeedRange{seedStart: instruction.destinationStart + instruction.rangeLength, length: seed.seedStart + seed.length - (instruction.sourceStart + instruction.rangeLength)})

				}
				fmt.Println("finish processing the seeds with instruction")

			}
		}
		seeds = newSeeds
	}
	minValue := seeds[0].seedStart
	for _, seed := range seeds {
		if seed.seedStart < minValue {
			minValue = seed.seedStart
		}

	}
	return minValue

}
