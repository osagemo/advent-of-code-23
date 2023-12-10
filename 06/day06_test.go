package main

import (
	"fmt"
	"testing"
)

const input1 = `Time:      7  15   30
Distance:  9  40  200`

func TestDay6Part1(t *testing.T) {
	result := Part1(input1)
	expected := 288

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay6Part2(t *testing.T) {
	result := Part2(input1)
	expected := 71503

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}
