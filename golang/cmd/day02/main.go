package main

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"strings"
)

var limits map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func part1(slice []string) int {
	var total int
	for i, line := range slice {
		if processLine1(line) {
			total += i + 1
		}
	}
	return total
}

func part2(slice []string) int {
	return 0
}

func processLine1(line string) bool {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	games := extractGames(line)
	for _, game := range games {
		for _, pair := range strings.Split(game, ",") {
			color, number := extractColorAndNumber(pair)
			if number > limits[color] {
				return false
			}
		}
	}
	return true
}

func extractGames(line string) []string {
	arr := strings.Split(line, ":")
	games := strings.Split(arr[1], ";")
	for i, game := range games {
		trimmedGame := strings.TrimSpace(game)
		games[i] = trimmedGame
	}
	return games
}

func extractColorAndNumber(pair string) (string, int) {
	pair = strings.TrimSpace(pair)
	arr := strings.Split(pair, " ")
	numberString := arr[0]
	color := arr[1]
	number, err := strconv.Atoi(numberString)
	if err != nil {
		panic(err)
	}
	return color, number
}

func main() {
	slice := utils.ReadFileAsSlice("2")
	fmt.Println("Day 2")
	fmt.Println("Part 1 answer:", part1(slice))
	fmt.Println("Part 2 answer:", part2(slice))
}
