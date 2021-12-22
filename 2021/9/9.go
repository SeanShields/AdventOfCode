package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	path := os.Args[1]
	input := readFile(path)

	part1Answer := solvePart1(input)
	fmt.Printf("Part 1 Answer: %v\n", part1Answer)

	part2Answer := solvePart2(input)
	fmt.Printf("Part 2 Answer: %v", part2Answer)
}

func solvePart1(input string) int {
	rows := strings.Split(input, "\r\n")
	low := findLowPoints(rows)
	lowPoints := []int{}
	for _, p := range low {
		lowPoints = append(lowPoints, p.digit)
	}
	return sum(lowPoints) + len(lowPoints)
}

func solvePart2(input string) int {
	rows := strings.Split(input, "\r\n")
	lowPoints := findLowPoints(rows)
	basins := [][]Point{}
	for _, point := range lowPoints {
		basin := Basin{point}
		basin.calculate(rows, point)
		basins = append(basins, basin)
	}
	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})
	return len(basins[0]) * len(basins[1]) * len(basins[2])
}

type Basin []Point

type Point struct {
	digit       int
	r           int
	c           int
	initialized bool
}

func findLowPoints(rows []string) []Point {
	lowPoints := []Point{}
	for r, row := range rows {
		for c, char := range row {
			digit := int(char - '0')
			adjacents := findAdjacents(rows, r, c)
			if isLow(adjacents, digit) {
				lowPoints = append(lowPoints, Point{digit, r, c, false})
			}
		}
	}
	return lowPoints
}

func findAdjacents(rows []string, r int, c int) []Point {
	adjacents := []Point{}
	top := Point{}
	right := Point{}
	bottom := Point{}
	left := Point{}
	if r != 0 {
		top = findPoint(rows, r-1, c)
		top.initialized = true
	}
	if c < len(rows[0])-1 {
		right = findPoint(rows, r, c+1)
		right.initialized = true
	}
	if r < len(rows)-1 {
		bottom = findPoint(rows, r+1, c)
		bottom.initialized = true
	}
	if c != 0 {
		left = findPoint(rows, r, c-1)
		left.initialized = true
	}

	if top.initialized {
		adjacents = append(adjacents, top)
	}
	if right.initialized {
		adjacents = append(adjacents, right)
	}
	if bottom.initialized {
		adjacents = append(adjacents, bottom)
	}
	if left.initialized {
		adjacents = append(adjacents, left)
	}
	return adjacents
}

func findPoint(rows []string, r int, c int) Point {
	for i, row := range rows {
		if i == r {
			for j, char := range row {
				if j == c {
					return Point{int(char - '0'), r, c, false}
				}
			}
		}
	}
	panic("point not found")
}

func (basin *Basin) calculate(rows []string, lowPoint Point) {
	for _, adjacent := range findAdjacents(rows, lowPoint.r, lowPoint.c) {
		if adjacent.digit != 9 && adjacent.digit > lowPoint.digit && !exists(*basin, adjacent) {
			*basin = append(*basin, adjacent)
			basin.calculate(rows, adjacent)
		}
	}
}

func exists(points []Point, point Point) bool {
	for _, p := range points {
		if p.digit == point.digit && p.r == point.r && p.c == point.c {
			return true
		}
	}
	return false
}

func sum(slice []int) int {
	sum := 0
	for _, s := range slice {
		sum += s
	}
	return sum
}

func isLow(points []Point, i int) bool {
	for _, p := range points {
		if i >= p.digit {
			return false
		}
	}
	return true
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
