package main

import (
	"fmt"
	"testing"
)

const input1 = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestDay3Part1(t *testing.T) {
	result := Part1(input1)
	expected := 4361

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay3Part2(t *testing.T) {
	result := Part2(input1)
	expected := 467835
	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay3SymbolLocations(t *testing.T) {
	expectedSymbolLocations := map[rune][]point{
		'*': {
			{3, 1},
			{3, 4},
			{5, 8},
		},
		'#': {
			{6, 3},
		},
		'+': {
			{5, 5},
		},
		'$': {
			{3, 8},
		},
	}
	_, symbolLocations := parseNumberAndSymbolLocations(input1)
	for expectedSymbol, locations := range expectedSymbolLocations {
		for _, point := range locations {
			symbol, ok := symbolLocations[point]
			if !ok || symbol != expectedSymbol {
				fmt.Printf("Expected symbol %s at location %v\n", string(expectedSymbol), point)
			}
		}
	}
}
func TestDay3NumberLocation(t *testing.T) {
	expectedSymbolLocation := map[int][]point{
		592: {
			{2, 6},
			{3, 6},
			{4, 6},
		},
	}
	numberLocations, _ := parseNumberAndSymbolLocations(input1)
	for expectedNumber, locations := range expectedSymbolLocation {
		var prevNumber *int
		for _, point := range locations {
			number, ok := numberLocations[point]
			// Number is correctly mapped
			if !ok || *number != expectedNumber {
				fmt.Printf("Expected number %d at location %v", expectedNumber, point)
				fmt.Printf(" got: %v\n", *number)
				t.Fail()
			}
			// All coordinates to number point to the same address
			if prevNumber != nil && prevNumber != number {
				fmt.Printf("Expected address %p for number %v", prevNumber, *number)
				t.Fail()
			}
			prevNumber = number
		}
	}
}
