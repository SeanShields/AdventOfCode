package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Args[1]
	input := readFile(path)

	part1Answer := solvePart1(input)
	fmt.Printf("Part 1 Answer: %v\n", part1Answer)

	part2Answer := solvePart2(input)
	fmt.Printf("Part 2 Answer: %v", part2Answer)
}

type coords struct {
	x int
	y int
}

func solvePart1(input string) int {
	total := 1
	visitedCoords := []coords{{0, 0}}

	for _, move := range input {
		coord := visitedCoords[len(visitedCoords)-1]
		switch move {
		case '^':
			coord.y += 1
		case 'v':
			coord.y -= 1
		case '>':
			coord.x += 1
		case '<':
			coord.x -= 1
		}

		if !coordsExist(visitedCoords, coord) {
			total++
		}
		visitedCoords = append(visitedCoords, coord)
	}
	return total
}

func solvePart2(input string) int {
	total := 1
	santaCoords := []coords{{0, 0}}
	roboSantaCoords := []coords{{0, 0}}

	for i, move := range input {
		coord := coords{}
		isSanta := i%2 == 0
		if isSanta {
			coord = santaCoords[len(santaCoords)-1]
		} else {
			coord = roboSantaCoords[len(roboSantaCoords)-1]
		}

		switch move {
		case '^':
			coord.y += 1
		case 'v':
			coord.y -= 1
		case '>':
			coord.x += 1
		case '<':
			coord.x -= 1
		}

		if !coordsExist(santaCoords, coord) && !coordsExist(roboSantaCoords, coord) {
			total++
		}

		if isSanta {
			santaCoords = append(santaCoords, coord)
		} else {
			roboSantaCoords = append(roboSantaCoords, coord)
		}
	}
	return total
}

func coordsExist(coords []coords, coord coords) bool {
	for _, c := range coords {
		if c.x == coord.x && c.y == coord.y {
			return true
		}
	}
	return false
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
