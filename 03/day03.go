package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type point struct {
	x int
	y int
}

// all directions
var acceptedDirections = []point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1},
}

func Part1(input string) int {
	uniquePartNumbers := map[*int]struct{}{} // "set"
	partNumberSum := 0

	// parse lines
	numberLocations, symbolLocations := parseNumberAndSymbolLocations(input)

	// find part numbers neighbouring symbols
	for point := range symbolLocations {
		neighbouringNumbers := getAdjacentNumbers(point, numberLocations)
		for _, num := range neighbouringNumbers {
			uniquePartNumbers[num] = struct{}{}
		}
	}
	// calculate sum of part numbers
	for num := range uniquePartNumbers {
		partNumberSum += *num
	}
	return partNumberSum
}

func Part2(input string) int {
	gearRatioSum := 0
	// parse lines
	numberLocations, symbolLocations := parseNumberAndSymbolLocations(input)

	// find gears and sum "gear ratios"
	// "A gear is any * symbol that is adjacent to exactly two part numbers."
	for point, symbol := range symbolLocations {
		if symbol == '*' {
			neighbouringNumbers := getAdjacentNumbers(point, numberLocations)
			if len(neighbouringNumbers) == 2 {
				gearRatio := 1
				for _, partNumber := range neighbouringNumbers {
					gearRatio *= *partNumber
				}
				gearRatioSum += gearRatio
			}
		}
	}

	return gearRatioSum
}

func getAdjacentNumbers(p point, numberLocations map[point]*int) []*int {
	adjacentNumbers := []*int{}

	for _, dif := range acceptedDirections {
		adjacentpoint := point{p.x + dif.x, p.y + dif.y}
		adjacentNumber, ok := numberLocations[adjacentpoint]
		if ok {
			if !slices.Contains(adjacentNumbers, adjacentNumber) {
				adjacentNumbers = append(adjacentNumbers, adjacentNumber)
			}
		}
	}

	return adjacentNumbers
}

// Example lines
// 467..114..
// ...*......
func parseNumberAndSymbolLocations(input string) (
	numberLocations map[point]*int,
	symbolLocations map[point]rune,
) {
	numberLocations = map[point]*int{}
	symbolLocations = map[point]rune{}

	// find part numbers and symbols in each input line
	// store their coordinates
	for y, line := range strings.Split(input, "\n") {
		// build currentNumber for each consecutive digit
		// int pointer due to multiple coordinates referencing the same number
		currentNumber := new(int)

		for x, c := range line {
			digit, e := strconv.Atoi(string(c))
			isDigit := e == nil

			if isDigit {
				*currentNumber = appendDigit(*currentNumber, digit)
			}

			// end of consecutive digits
			if (!isDigit || x == len(line)-1) && *currentNumber != 0 {
				numLength := digitLen(*currentNumber)
				// store all x coordinates spanning the number
				startingX := x - numLength
				for xD := 0; xD < numLength; xD++ {
					numberLocations[point{startingX + xD, y}] = currentNumber
				}
				currentNumber = new(int)
			}
			if c == '.' {
				continue
			}
			if !isDigit { // symbol
				symbolLocations[point{x, y}] = c
			}
		}
	}
	return
}

func digitLen(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func appendDigit(number int, digit int) int {
	if number == 0 {
		number = digit
	} else {
		number = number*10 + digit
	}
	return number
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
