package main

import (
	_ "embed"
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type almanacMap struct {
	source string
	dest   string
	ranges []almanacRange
}

type almanacRange struct {
	sourceStart int
	destStart   int
	length      int
}

func Part1(input string) int {
	seeds := parseSeedsPart1(input)
	typeToAlmanacMap := mapSourceToAlmanacMaps(input, false)

	locations := []int{}
	for _, seed := range seeds {
		seedLocation := findPart1("seed", seed, typeToAlmanacMap)
		locations = append(locations, seedLocation)
	}

	// parse almanac
	return slices.Min(locations)
}

// Function Must Locate...
func Part2(input string) int {
	seedRanges := parseSeedsPart2(input)
	reverseTypeToAlmanacMap := mapSourceToAlmanacMaps(input, true)

	// Brute force the lowest possible location by mapping any location number
	// and checking if it is within any of the seed ranges...
	for possibleLocation := 0; possibleLocation < math.MaxInt32; possibleLocation++ {
		mappedSeed := findPart1("location", possibleLocation, reverseTypeToAlmanacMap)
		for _, seedRange := range seedRanges {
			if withinRange(mappedSeed, seedRange) {
				return possibleLocation
			}
		}
	}

	return -1
}

func parseSeedsPart1(input string) []int {
	seedBlock := strings.Split(input, "\n\n")[0]
	seeds := []int{}
	for _, num := range strings.Split(strings.Split(seedBlock, ": ")[1], " ") {
		seed, e := strconv.Atoi(num)
		if e != nil {
			panic(e)
		}
		seeds = append(seeds, seed)
	}

	return seeds
}

func parseSeedsPart2(input string) []almanacRange {
	seedRanges := []almanacRange{}
	isRange := false
	prev := 0

	seedBlock := strings.Split(input, "\n\n")[0]
	for _, num := range strings.Split(strings.Split(seedBlock, ": ")[1], " ") {
		val, e := strconv.Atoi(num)
		if e != nil {
			panic(e)
		}

		if isRange {
			seedRange := almanacRange{sourceStart: prev, length: val}
			seedRanges = append(seedRanges, seedRange)
		}

		prev = val
		isRange = !isRange
	}

	return seedRanges
}

func mapSourceToAlmanacMaps(input string, reverse bool) map[string]almanacMap {
	typeToAlmanacMap := map[string]almanacMap{}
	blocks := strings.Split(input, "\n\n")

	for _, block := range blocks[1:] {
		blockLines := strings.Split(block, "\n")
		header := strings.Split(strings.Split(blockLines[0], " ")[0], "-")
		if len(header) < 2 {
			continue
		}

		if reverse {
			header[0], header[2] = header[2], header[0]
		}
		source, destination := header[0], header[2]
		almanacMap := almanacMap{source: source, dest: destination}

		for _, line := range blockLines[1:] {
			lineParts := strings.Split(line, " ")
			if reverse {
				lineParts[0], lineParts[1] = lineParts[1], lineParts[0]
			}
			sourceStart, e1 := strconv.Atoi(lineParts[1])
			destStart, e2 := strconv.Atoi(lineParts[0])
			length, e3 := strconv.Atoi(lineParts[2])

			if e := errors.Join(e1, e2, e3); e != nil {
				panic(e)
			}

			almanacRange := almanacRange{
				sourceStart: sourceStart,
				destStart:   destStart,
				length:      length,
			}
			almanacMap.ranges = append(almanacMap.ranges, almanacRange)
		}

		typeToAlmanacMap[source] = almanacMap
	}

	return typeToAlmanacMap
}

func findPart1(dest string, value int, maps map[string]almanacMap) int {
	almanacMap, ok := maps[dest]
	if !ok {
		return value
	}

	target := 0
	for _, almanacRange := range almanacMap.ranges {
		if withinRange(value, almanacRange) {
			target = almanacRange.destStart + (value - almanacRange.sourceStart)
		}
	}
	if target == 0 {
		target = value
	}

	return findPart1(almanacMap.dest, target, maps)
}

func withinRange(value int, almanacRange almanacRange) bool {
	return value >= almanacRange.sourceStart && value < almanacRange.sourceStart+almanacRange.length
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
