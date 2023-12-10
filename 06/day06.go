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

func Part1(input string) int {
	times := []int{}
	distances := []int{}

	input = strings.ReplaceAll(input, "Time:      ", "")
	input = strings.ReplaceAll(input, "Distance:  ", "")

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if line == "" {
			continue
		}

		tokens := strings.Fields(line)
		for _, token := range tokens {
			num, e := strconv.Atoi(token)
			if e != nil {
				panic(e)
			}
			if i == 0 {
				times = append(times, num)
			} else {
				distances = append(distances, num)
			}
		}
	}

	productWinners := 1
	for i := range times {
		productWinners *= waysToWin(times[i], distances[i])
	}

	return productWinners
}

func Part2(input string) int {
	time := 0
	distance := 0

	input = strings.ReplaceAll(input, "Time:      ", "")
	input = strings.ReplaceAll(input, "Distance:  ", "")

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if line == "" {
			continue
		}

		tokens := strings.Fields(line)
		wholeNumber := ""
		for _, token := range tokens {
			wholeNumber += token
		}
		num, e := strconv.Atoi(wholeNumber)
		if e != nil {
			panic(e)
		}
		if i == 0 {
			time = num
		} else {
			distance = num
		}
	}

	return waysToWin(time, distance)
}

func waysToWin(time int, distance int) int {
	numWinners := 0
	for j := 0; j < time; j++ {
		if j*(time-j) > distance {
			numWinners++
		}
	}
	return numWinners
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
