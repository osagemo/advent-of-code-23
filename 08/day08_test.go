package main

import (
	"fmt"
	"testing"
)

const input1 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

const input2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ`

const input3 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestDay8Part1(t *testing.T) {
	expected1 := 2
	expected2 := 6

	result1 := Part1(input1)
	if result1 != expected1 {
		fmt.Printf("got %v, expected %v\n", result1, expected1)
		t.Fail()
	}

	result2 := Part1(input2)
	if result2 != expected2 {
		fmt.Printf("got %v, expected %v\n", result2, expected2)
		t.Fail()
	}
}

func TestDay8Part2(t *testing.T) {
	result := Part2(input3)
	expected := 6

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}
