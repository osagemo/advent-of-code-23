package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

type node struct {
	name      string
	neighbors []string
}

//go:embed input.txt
var input string

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	instructions := parseInstructions(lines[0])
	nodes := parseNodes(lines[1:])

	return findNode(nodes, "AAA", "ZZZ", instructions)
}

func Part2(input string) int64 {
	lines := strings.Split(input, "\n")
	instructions := parseInstructions(lines[0])
	nodes := parseNodes(lines[1:])

	startNodes := []string{}
	for name := range nodes {
		if strings.HasSuffix(name, "A") {
			startNodes = append(startNodes, name)
		}
	}

	return findNodes(nodes, startNodes, instructions)
}

// e.g. "LLR" -> [0, 0, 1]
func parseInstructions(instructionLine string) []int {
	instructions := []int{}

	for _, instruction := range instructionLine {
		if instruction == 'R' {
			instructions = append(instructions, 1)
		} else if instruction == 'L' {
			instructions = append(instructions, 0)
		} else {
			panic("Invalid instruction")
		}
	}

	return instructions
}

func parseNodes(lines []string) map[string]node {
	nodeMap := make(map[string]node)

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " = ")
		name := parts[0]
		parts[1] = strings.Trim(parts[1], "()")
		neighbors := strings.Split(parts[1], ", ")

		node := node{name: name, neighbors: neighbors}
		nodeMap[name] = node
	}

	return nodeMap
}

// Iterate over the instruction sequence until we find the destination node
func findNode(nodeMap map[string]node, src string, dest string, instructions []int) int {
	current := nodeMap[src]
	steps := 0
	for found := false; !found; {
		for _, instruction := range instructions {
			steps++
			current = nodeMap[current.neighbors[instruction]]
			if current.name == dest {
				found = true
			}
		}
	}

	return steps
}

type evaluateDestination func([]node) bool

func findNodes(nodeMap map[string]node, src []string, instructions []int) int64 {
	currents := []node{}
	var steps int64 = 0

	for _, name := range src {
		currents = append(currents, nodeMap[name])
	}

	nMapped := 0
	stepsPerSrc := make([]int64, len(src))
	for found := false; !found; {
		for _, instruction := range instructions {
			steps++
			// All nodes found
			if nMapped == len(src) {
				found = true
				break
			}
			for i, current := range currents {
				if stepsPerSrc[i] != 0 {
					continue
				}
				currents[i] = nodeMap[current.neighbors[instruction]]
				// never happens
				if steps > 1 && currents[i].name == src[i] {
					fmt.Println("Found loop!", currents[i].name, "from", src[i], "at", steps)
				}
				if strings.HasSuffix(currents[i].name, "Z") {
					fmt.Println("Found loop!", currents[i].name, "from", src[i], "at", steps)
					stepsPerSrc[i] = steps
					nMapped++
				}
			}
		}
	}

	fmt.Println("found all Zs", src, stepsPerSrc)

	// when do all iterations sync?
	for found := false; !found; {
		if !allEqual(stepsPerSrc) {
			for i, nSteps := range stepsPerSrc {
				stepsPerSrc[i] = nSteps + nSteps
			}
		} else {
			fmt.Println(stepsPerSrc)
			found = true
		}
	}

	// not working, currently syncs at overflow
	return stepsPerSrc[0]
}

func allEqual(numbers []int64) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != numbers[0] {
			return false
		}
	}

	return true
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
