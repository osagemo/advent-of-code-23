package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

func (p Point) Add(dir string, steps int) Point {
	diff := directions[dir]
	return Point{p.x + diff.x*steps, p.y + diff.y*steps}
}

var directions = map[string]Point{
	"U": {0, 1},
	"D": {0, -1},
	"L": {-1, 0},
	"R": {1, 0},
}

func Part1(input string) int {
	trenchCorners := parsePerimeter(input)
	perimeterSize := getPerimeterSize(trenchCorners)
	innerArea := polygonArea(trenchCorners)

	areaWithPerimeter := innerArea - perimeterSize/2 + perimeterSize

	return areaWithPerimeter + 1
}

func parsePerimeter(input string) []Point {
	perimeter := []Point{}
	curr := Point{0, 0}
	perimeter = append(perimeter, curr)

	lines := strings.Split(input, "\n")
	for _, line := range lines[0 : len(lines)-1] {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance, e := strconv.Atoi(parts[1])
		if e != nil {
			panic(e)
		}

		curr = curr.Add(direction, distance)
		perimeter = append(perimeter, curr)
	}

	return perimeter
}

func parsePerimeterPart2(input string) []Point {
	perimeter := []Point{}
	curr := Point{0, 0}
	perimeter = append(perimeter, curr)

	lines := strings.Split(input, "\n")
	for _, line := range lines[0 : len(lines)-1] {
		parts := strings.Split(line, " ")
		color := strings.Trim(parts[2], " ()#")
		// "The first five hexadecimal digits encode the distance in meters
		//   as a five-digit hexadecimal number"
		hex := color[0:5]
		distance, e := strconv.ParseInt(hex, 16, 64)
		if e != nil {
			panic(e)
		}

		// "The last hexadecimal digit encodes the direction to dig:
		//   0 means R, 1 means D, 2 means L, and 3 means U."
		directionNumber := color[len(color)-1]
		direction := ""
		switch directionNumber {
		case '0':
			direction = "R"
		case '1':
			direction = "D"
		case '2':
			direction = "L"
		case '3':
			direction = "U"
		}

		curr = curr.Add(direction, int(distance))
		perimeter = append(perimeter, curr)
	}

	return perimeter
}

func getPerimeterSize(perimeterPoints []Point) int {
	size := 0
	prev := perimeterPoints[0]
	for i, c := range perimeterPoints {
		if i == 0 {
			continue
		}
		distance := math.Abs(float64(c.x-prev.x)) + math.Abs(float64(c.y-prev.y))
		size += int(distance)
		prev = c
	}
	return size
}

func polygonArea(vertices []Point) int {
	area := 0
	prev := vertices[0]

	// irregular polygon area
	// https://www.mathopenref.com/coordpolygonarea.html
	for i, c := range vertices {
		if i == 0 {
			continue
		}
		area += (prev.x * c.y) - (prev.y * c.x)
		prev = c
	}

	if area < 0 {
		area = -area
	}
	return area / 2
}

func Part2(input string) int {
	trenchCorners := parsePerimeterPart2(input)
	perimeterSize := getPerimeterSize(trenchCorners)
	area := polygonArea(trenchCorners)

	answer := area + perimeterSize/2 + 1

	return answer
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
