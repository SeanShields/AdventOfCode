package main

import (
	"fmt"
	"math"
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
	head := vector{0, 0}
	tail := vector{0, 0}
	visited := []vector{tail}
	instructions := strings.Split(input, "\r\n")
	for _, instruction := range instructions {
		split := strings.Split(instruction, " ")
		direction := split[0]
		length, _ := strconv.Atoi(split[1])
		for length > 0 {
			switch direction {
			case "U":
				head.x++
				break
			case "R":
				head.y++
				break
			case "D":
				head.x--
				break
			case "L":
				head.y--
				break
			}
			if !touching(head, tail) {
				tail = moveToHead(head, tail)
				exists := false
				for _, v := range visited {
					if tail.x == v.x && tail.y == v.y {
						exists = true
					}
				}
				if !exists {
					visited = append(visited, tail)
				}
			}
			length--
		}
	}
	return len(visited)
}

func solvePart2(input string) int {
	vectors := [10]vector{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}
	visited := []vector{{0, 0}}
	instructions := strings.Split(input, "\r\n")
	for _, instruction := range instructions {
		split := strings.Split(instruction, " ")
		direction := split[0]
		length, _ := strconv.Atoi(split[1])
		for l := length; l > 0; l-- {
			for v, _ := range vectors {
				if v == 0 {
					switch direction {
					case "U":
						vectors[v].x++
						break
					case "R":
						vectors[v].y++
						break
					case "D":
						vectors[v].x--
						break
					case "L":
						vectors[v].y--
						break
					}
				} else if !touching(vectors[v-1], vectors[v]) {
					next := moveToHead(vectors[v-1], vectors[v])
					vectors[v] = next
				}

				if v == 9 {
					exists := false
					for _, s := range visited {
						if vectors[v].x == s.x && vectors[v].y == s.y {
							exists = true
						}
					}
					if !exists {
						visited = append(visited, vectors[v])
					}
				}
			}
		}
		fmt.Printf("%v: %v\n\n", instruction, vectors)
	}
	return len(visited)
}

type vector struct {
	x int
	y int
}

func moveToHead(head vector, tail vector) vector {
	potentials := []vector{
		{tail.x, tail.y + 1},
		{tail.x + 1, tail.y + 1},
		{tail.x + 1, tail.y},
		{tail.x + 1, tail.y - 1},
		{tail.x, tail.y - 1},
		{tail.x - 1, tail.y - 1},
		{tail.x - 1, tail.y},
		{tail.x - 1, tail.y + 1},
	}

	var shortest vector
	for i, v := range potentials {
		if i == 0 {
			shortest = v
			continue
		}

		if manhattanDistance(v, head) < manhattanDistance(shortest, head) {
			shortest = v
		}
	}

	return shortest
}

func touching(head vector, tail vector) bool {
	if tail.x == head.x && tail.y == head.y {
		return true
	}

	potentials := []vector{
		{tail.x, tail.y + 1},
		{tail.x + 1, tail.y + 1},
		{tail.x + 1, tail.y},
		{tail.x + 1, tail.y - 1},
		{tail.x, tail.y - 1},
		{tail.x - 1, tail.y - 1},
		{tail.x - 1, tail.y},
		{tail.x - 1, tail.y + 1},
	}

	for _, p := range potentials {
		if p.x == head.x && p.y == head.y {
			return true
		}
	}
	return false
}

func manhattanDistance(a vector, b vector) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func max(arr []int) int {
	s := 0
	for _, a := range arr {
		if a > s {
			s = a
		}
	}
	return s
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
