package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

var maxOfColors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// Example game:
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func Part1(input string) int {
	// ignore set dividers
	input = strings.ReplaceAll(input, ";", ",")

	idSum := 0
	for i, game := range strings.Split(input, "\n") {
		if game == "" {
			continue
		}
		gameId := i + 1
		valid := true
		sets := strings.Split(game, ": ")[1]

		if valid = allCubeCountsLessThanMax(sets); valid {
			idSum += gameId
		}
	}

	return idSum
}

func Part2(input string) int {
	// ignore set dividers
	input = strings.ReplaceAll(input, ";", ",")

	powerSum := 0
	for _, game := range strings.Split(input, "\n") {
		if game == "" {
			continue
		}
		maxCounts := map[string]int{"red": 0, "green": 0, "blue": 0}
		sets := strings.Split(game, ": ")[1]

		for color, count := range parseAllCubeCounts(sets) {
			if count > maxCounts[color] {
				maxCounts[color] = count
			}
		}
		power := 1
		for _, count := range maxCounts {
			power *= count
		}
		powerSum += power
	}
	return powerSum
}

func allCubeCountsLessThanMax(set string) bool {
	valid := true
	for _, cubeCount := range strings.Split(set, ", ") {
		if color, count := parseCubeCount(cubeCount); count > maxOfColors[color] {
			valid = false
			break
		}
	}
	return valid
}

func parseAllCubeCounts(set string) map[string]int {
	cubeCounts := map[string]int{}
	for _, cubeCount := range strings.Split(set, ", ") {
		color, count := parseCubeCount(cubeCount)
		cubeCounts[color] = count
	}
	return cubeCounts
}

func parseCubeCount(cubeCount string) (color string, count int) {
	cubeCount = strings.Trim(cubeCount, " ")

	parts := strings.Split(cubeCount, " ")
	if len(parts) != 2 {
		panic("Error while parsing cube count")
	}
	color = parts[1]
	count, e := strconv.Atoi(parts[0])
	if e != nil {
		panic(e)
	}
	return
}

func main() {
	input := strings.ReplaceAll(input, "\r\n", "\n")
	fmt.Println("Day 1")
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input))
	fmt.Println(time.Since(start))
}
