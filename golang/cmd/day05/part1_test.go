package day05

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)


func TestDay05(t *testing.T) {
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

	for i, chunk := range chunks {
		// ignore first line
		maps := make([][]int, 0)
		fmt.Printf("Chunk %d\n", i)
		lines := strings.Split(chunk, "\n")[1:]
		for _, line := range lines {
			var destRange, sourceRange, rangeLength string
			fmt.Sscanf(line, "%s %s %s", &destRange, &sourceRange, &rangeLength)
			fmt.Printf("destRange: %s, sourceRange: %s, rangeLength: %s\n", destRange, sourceRange, rangeLength)
			maps = append(maps, []int{int(destRange[0]), int(sourceRange[0]), int(rangeLength[0])})
		}
	}

}
