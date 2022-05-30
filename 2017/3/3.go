package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	path := os.Args[1]
	input := readFile(path)

	part1Answer := solvePart1(input)
	fmt.Printf("Part 1 Answer: %v\n", part1Answer)

	part2Answer := solvePart2(input)
	fmt.Printf("Part 2 Answer: %v", part2Answer)
}

func solvePart1(input string) float64 {
	current := coords{0, 0, 1}
	square := 1
	moves := 0
	port, _ := strconv.Atoi(input)
	counter := 1
	for square < port {
		if square+moves+1 > port {
			moves = port - square
		} else if counter == 1 || counter == 3 {
			moves++
		}
		switch counter {
		case 1:
			current.x += moves
		case 2:
			current.y += moves
		case 3:
			current.x -= moves
		case 4:
			current.y -= moves
			counter = 0
		}
		square += moves
		counter++
	}
	return manhattanDistance(coords{0, 0, 1}, current)
}

func solvePart2(input string) int {
	square := 1
	current := coords{0, 0, square}
	visited := []coords{current}
	moves := 0
	counter := 1
	for {
		if counter == 1 || counter == 3 {
			moves++
		}

		for m := 0; m < moves; m++ {
			if current.num > 277678 {
				return current.num
			}
			switch counter {
			case 1:
				current.x++
			case 2:
				current.y++
			case 3:
				current.x--
			case 4:
				current.y--
			}
			current.num = sumAdjacentNumbers(getAdjacentCoords(current, visited))
			visited = append(visited, current)
			square++
		}

		if counter == 4 {
			counter = 0
		}
		counter++
	}
}

func sumAdjacentNumbers(c []coords) int {
	sum := 0
	for _, coord := range c {
		sum += coord.num
	}
	return sum
}

func getAdjacentCoords(c coords, visited []coords) []coords {
	adjacents := []coords{}

	// top
	topCoords := coords{c.x, c.y + 1, 0}
	topCoords.num = findNumber(topCoords, visited)
	if topCoords.num != 0 {
		adjacents = append(adjacents, topCoords)
	}

	// right
	rightCoords := coords{c.x + 1, c.y, 0}
	rightCoords.num = findNumber(rightCoords, visited)
	if rightCoords.num != 0 {
		adjacents = append(adjacents, rightCoords)
	}

	// bottom
	bottomCoords := coords{c.x, c.y - 1, 0}
	bottomCoords.num = findNumber(bottomCoords, visited)
	if bottomCoords.num != 0 {
		adjacents = append(adjacents, bottomCoords)
	}

	// left
	leftCoords := coords{c.x - 1, c.y, 0}
	leftCoords.num = findNumber(leftCoords, visited)
	if leftCoords.num != 0 {
		adjacents = append(adjacents, leftCoords)
	}

	// top-right
	topRightCoords := coords{c.x + 1, c.y + 1, 0}
	topRightCoords.num = findNumber(topRightCoords, visited)
	if topRightCoords.num != 0 {
		adjacents = append(adjacents, topRightCoords)
	}

	// bottom-right
	bottomRightCoords := coords{c.x + 1, c.y - 1, 0}
	bottomRightCoords.num = findNumber(bottomRightCoords, visited)
	if bottomRightCoords.num != 0 {
		adjacents = append(adjacents, bottomRightCoords)
	}

	// bottom-left
	bottomLeftCoords := coords{c.x - 1, c.y - 1, 0}
	bottomLeftCoords.num = findNumber(bottomLeftCoords, visited)
	if bottomLeftCoords.num != 0 {
		adjacents = append(adjacents, bottomLeftCoords)
	}

	// top-left
	topLeftCoords := coords{c.x - 1, c.y + 1, 0}
	topLeftCoords.num = findNumber(topLeftCoords, visited)
	if topLeftCoords.num != 0 {
		adjacents = append(adjacents, topLeftCoords)
	}

	return adjacents
}

func findNumber(c coords, cc []coords) int {
	for _, coord := range cc {
		if c.x == coord.x && c.y == coord.y {
			return coord.num
		}
	}
	return 0
}

func manhattanDistance(a coords, b coords) float64 {
	return math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
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

type coords struct {
	x   int
	y   int
	num int
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
