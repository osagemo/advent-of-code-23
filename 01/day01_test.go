package main

import (
	"fmt"
	"testing"
)

const input1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const input2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestDay1Part1(t *testing.T) {
	result := Part1(input1)
	expected := 142

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay1Part2(t *testing.T) {
	result := Part2(input2)
	expected := 281

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}
