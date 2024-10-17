package day05

import (
	"fmt"
	"strconv"
	"strings"
)

func Part2(file string) int {
	fmt.Println("Part II: Starting ")
	chunks := strings.Split(file, "\n\n")
	seedsChunk := chunks[0]
	chunks = chunks[1:]

	//Let's get all the seeds
	seeds := make([]int, 0)
	seedLines := strings.TrimPrefix(seedsChunk, "seeds: ")
	for _, line := range strings.Split(seedLines, "\n") {
		fmt.Println("inside line 19")
		seedNumbers := strings.Fields(line)
		if (len(seedNumbers) % 2) != 0 {
			panic(fmt.Sprintf("Error: seeds are not in pairs: %v\n", seedNumbers))
		}
		for i := 0; i < len(seedNumbers); i += 2 {
			seedStart, err := strconv.Atoi(seedNumbers[i])
			if err != nil {
				panic(fmt.Sprintf("Error converting seed to int: %v\n", seedNumbers[i]))
			}
			length, err := strconv.Atoi(seedNumbers[i+1])
			if err != nil {
				panic(fmt.Sprintf("Error converting seed to int: %v\n", seedNumbers[i+1]))
			}
			for j := seedStart; j <= seedStart+length-1; j++ {
				seeds = append(seeds, j)
			}
			fmt.Println("inside line 36")
		}
	}

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
