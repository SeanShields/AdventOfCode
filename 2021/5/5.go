package main

import (
	"fmt"
	"os"
	"strconv"
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
	lines := strings.Split(input, "\r\n")
	graph := Graph{}

	for _, line := range lines {
		start, end := parseLine(line)

		if isHorizontal(start, end) {
			graph.plotHorizontalLine(start.x, start.y, end.y)
		} else if isVertical(start, end) {
			graph.plotVerticalLine(start.x, end.x, start.y)
		}
	}

	return graph.sum(2)
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\r\n")
	graph := Graph{}

	for _, line := range lines {
		start, end := parseLine(line)

		if isHorizontal(start, end) {
			graph.plotHorizontalLine(start.x, start.y, end.y)
		} else if isVertical(start, end) {
			graph.plotVerticalLine(start.x, end.x, start.y)
		} else {
			graph.plotDiagonalLine(start.x, start.y, end.x, end.y)
		}
	}

	return graph.sum(2)
}

type Graph []Coords
type Coords struct {
	x             int
	y             int
	intersections int
}

func (graph *Graph) plotHorizontalLine(x int, y0 int, y1 int) {
	start := min(y0, y1)
	end := max(y0, y1)
	y := start
	for y <= end {
		if !graph.exists(x, y) {
			*graph = append(*graph, Coords{x, y, 0})
		}
		y++
	}
}

func (graph *Graph) plotVerticalLine(x0 int, x1 int, y int) {
	start := min(x0, x1)
	end := max(x0, x1)
	x := start
	for x <= end {
		if !graph.exists(x, y) {
			*graph = append(*graph, Coords{x, y, 0})
		}
		x++
	}
}

func (graph *Graph) plotDiagonalLine(x0 int, y0 int, x1 int, y1 int) {
	if x0 < x1 && y0 > y1 {
		// add to x, sub from y
		graph.plotDiagonalDown(x0, y0, x1, y1)
	} else if x0 > x1 && y0 < y1 {
		graph.plotDiagonalDown(x1, y1, x0, y0)
	} else if x0 < x1 && y0 < y1 {
		graph.plotDiagonalUp(x0, y0, x1, y1)
	} else if x0 > x1 && y0 > y1 {
		graph.plotDiagonalUp(x1, y1, x0, y0)
	} else {
		panic("no line found for diagonal")
	}
}

func (graph *Graph) plotDiagonalDown(x0 int, y0 int, x1 int, y1 int) {
	x := x0
	y := y0
	for {
		if !graph.exists(x, y) {
			*graph = append(*graph, Coords{x, y, 0})
		}
		if x == x1 && y == y1 {
			return
		}
		x++
		y--
	}
}

func (graph *Graph) plotDiagonalUp(x0 int, y0 int, x1 int, y1 int) {
	x := x0
	y := y0
	for {
		if !graph.exists(x, y) {
			*graph = append(*graph, Coords{x, y, 0})
		}
		if x == x1 && y == y1 {
			return
		}
		x++
		y++
	}
}

func getSlope(start Coords, end Coords) int {
	return (end.y - start.y) / (end.x - start.x)
}

func max(x int, y int) int {
	if x > y {
		return x
	} else if y > x {
		return y
	} else {
		return x
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	} else if y < x {
		return y
	} else {
		return x
	}
}

func parseLine(line string) (Coords, Coords) {
	coords := strings.Split(line, " -> ")
	start := toInt(strings.Split(coords[0], ","))
	end := toInt(strings.Split(coords[1], ","))
	return Coords{start[0], -start[1], 0}, Coords{end[0], -end[1], 0}
}

func (graph *Graph) exists(x int, y int) bool {
	for i, c := range *graph {
		if c.x == x && c.y == y {
			(*graph)[i].intersections++
			return true
		}
	}
	return false
}

func (graph Graph) sum(threshold int) int {
	sum := 0
	for _, i := range graph {
		if i.intersections >= threshold-1 {
			sum++
		}
	}
	return sum
}

func isHorizontal(coords1 Coords, coords2 Coords) bool {
	return coords1.x == coords2.x
}

func isVertical(coords1 Coords, coords2 Coords) bool {
	return coords1.y == coords2.y
}

func toInt(slice []string) []int {
	ret := []int{}
	for _, s := range slice {
		i, _ := strconv.Atoi(s)
		ret = append(ret, i)
	}
	return ret
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
