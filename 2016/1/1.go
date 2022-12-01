package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
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

type vector struct {
	x int
	y int
}

func solvePart1(input string) int {
	instructions := strings.Split(input, ", ")
	current := vector{0, 0}
	direction := "N"
	for _, instruction := range instructions {
		newDirection, distance := parseInstruction(instruction)
		switch direction {
		case "N":
			if newDirection == "R" {
				direction = "E"
				current.x += distance
			} else {
				direction = "W"
				current.x -= distance
			}
			break
		case "E":
			if newDirection == "R" {
				direction = "S"
				current.y -= distance
			} else {
				direction = "N"
				current.y += distance
			}
			break
		case "S":
			if newDirection == "R" {
				direction = "W"
				current.x -= distance
			} else {
				direction = "E"
				current.x += distance
			}
			break
		case "W":
			if newDirection == "R" {
				direction = "N"
				current.y += distance
			} else {
				direction = "S"
				current.y -= distance
			}
			break
		}
	}
	return manhattanDistance(vector{0, 0}, current)
}

func solvePart2(input string) int {
	instructions := strings.Split(input, ", ")
	current := vector{0, 0}
	visited := []vector{current}
	direction := "N"
	for _, instruction := range instructions {
		newDirection, distance := parseInstruction(instruction)
		switch direction {
		case "N":
			if newDirection == "R" {
				direction = "E"
				current.x += distance
			} else {
				direction = "W"
				current.x -= distance
			}
			break
		case "E":
			if newDirection == "R" {
				direction = "S"
				current.y -= distance
			} else {
				direction = "N"
				current.y += distance
			}
			break
		case "S":
			if newDirection == "R" {
				direction = "W"
				current.x -= distance
			} else {
				direction = "E"
				current.x += distance
			}
			break
		case "W":
			if newDirection == "R" {
				direction = "N"
				current.y += distance
			} else {
				direction = "S"
				current.y -= distance
			}
			break
		}
	}
	panic("Could not find visited vector")
}

func parseInstruction(instruction string) (string, int) {
	direction := regexp.MustCompile("[RL]").FindString(instruction)
	d := regexp.MustCompile("[0-9].*").FindString(instruction)
	distance, _ := strconv.Atoi(d)
	return direction, distance
}

func getVisited(current vector, destination vector, direction string) []vector {
	visited := []vector{}

	switch direction {
	case "N":
		break
	case "E":

		break
	case "S":

		break
	case "W":

		break
	}

	return visited
}

func manhattanDistance(a vector, b vector) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func vectorExists(vector vector, vectors []vector) bool {
	for _, v := range vectors {
		if vector.x == v.x && vector.y == v.y {
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
