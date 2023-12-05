package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func Part1(input string) int {
	pointSum := 0
	pointsPerCard := getAmountOfWinningNumbersPerCard(input)
	for _, p := range pointsPerCard {
		points := 0
		for i := 0; i < p; i++ {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
		pointSum += points
	}
	return pointSum
}

func Part2(input string) int {
	pointsPerCard := getAmountOfWinningNumbersPerCard(input)
	copiesOfCards := make([]int, len(pointsPerCard))
	// fill with 1s for the initial copies
	for i := 0; i < len(copiesOfCards); i++ {
		copiesOfCards[i] = 1
	}

	// increment the copies of the N (points) succeeding cards
	// by the number of copies of the winning card
	for i := 0; i < len(copiesOfCards); i++ {
		points := pointsPerCard[i]
		for j := 0; j < points; j++ {
			copiesOfCards[i+j+1] += copiesOfCards[i]
		}
	}

	cardSum := 0
	for _, p := range copiesOfCards {
		cardSum += p
	}
	// now you have a lot of cards, congratulations
	return cardSum
}

func getAmountOfWinningNumbersPerCard(input string) []int {
	points := []int{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		line = strings.Split(line, ":")[1]

		points = append(points, getAmountOfWinningNumbers(line))
	}

	return points
}

func getAmountOfWinningNumbers(cardNumbers string) int {
	seenNumbers := make(map[string]bool)
	duplicateNumbers := make(map[string]bool)

	numbers := strings.Split(cardNumbers, " ")
	// a winning number is a number that occurs more than once.
	// we can ignore the distinction of "winning numbers" and "card numbers" (|)
	// since there are no duplicate card numbers
	for _, number := range numbers {
		if len(number) == 0 || strings.Contains(number, "|") {
			continue
		}

		if seenNumbers[number] {
			duplicateNumbers[number] = true
		}
		seenNumbers[number] = true
	}

	return len(duplicateNumbers)
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
