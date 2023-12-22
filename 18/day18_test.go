package main

import (
	"fmt"
	"testing"
)

const input1 = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

func TestDay18Part1(t *testing.T) {
	result := Part1(input1)
	expected := 62

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay18PerimeterSize(t *testing.T) {
	perimeter := parsePerimeter(input1)
	result := getPerimeterSize(perimeter)
	expected := 38

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay18Part2(t *testing.T) {
	result := Part2(input1)
	expected := 952408144115

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}
