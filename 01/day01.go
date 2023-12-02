package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func Part1(input string) int {
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		if s == "" {
			continue
		}

		nums := []int{}
		for _, c := range s {
			n, e := strconv.Atoi(string(c))
			if e == nil {
				nums = append(nums, n)
			}
		}
		sum += nums[0]*10 + nums[len(nums)-1]
	}
	return sum
}

// Slower
func Part1_regex(input string) int {
	r, _ := regexp.Compile(`[^0-9\r\n]`)
	numbers := r.ReplaceAllString(input, "")
	sum := 0
	for _, s := range strings.Split(numbers, "\n") {
		if s == "" {
			continue
		}

		n1, e := strconv.Atoi(string(s[0]))
		if e != nil {
			panic(e)
		}
		n2, e := strconv.Atoi(string(s[len(s)-1]))
		if e != nil {
			panic(e)
		}
		sum += n1*10 + n2
	}

	return sum
}

func Part2(input string) int {
	numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, s := range numbers {
		r, _ := regexp.Compile("(" + s + ")")
		input = r.ReplaceAllStringFunc(input, func(match string) string {
			// Stupid :)
			return match + strconv.Itoa(i) + match
		})
	}

	return Part1(input)
}

func main() {
	fmt.Println("Day 1")
	start := time.Now()
	fmt.Println("Part 1: ", Part1(input))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println("Part 2: ", Part2(input))
	fmt.Println(time.Since(start))
}
