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

	part1 := part1(input)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := part2(input)
	fmt.Printf("Part 2: %v", part2)
}

func part1(input string) int {
	instructions := strings.Split(input, "\r\n")
	horizontal := 0
	depth := 0

	for _, instruction := range instructions {
		if instruction == "" {
			continue
		}

		s := strings.Split(instruction, " ")
		direction := s[0]
		amount, _ := strconv.Atoi(s[1])

		switch direction {
		case "forward":
			horizontal += amount
			break
		case "up":
			depth -= amount
			break
		case "down":
			depth += amount
			break
		default:
			panic("error")
		}
	}

	return horizontal * depth
}

func part2(input string) int {
	instructions := strings.Split(input, "\r\n")
	horizontal := 0
	depth := 0
	aim := 0

	for _, instruction := range instructions {
		if instruction == "" {
			continue
		}

		s := strings.Split(instruction, " ")
		direction := s[0]
		amount, _ := strconv.Atoi(s[1])

		switch direction {
		case "forward":
			horizontal += amount
			depth += aim * amount
			break
		case "up":
			aim -= amount
			break
		case "down":
			aim += amount
			break
		default:
			panic("error")
		}
	}

	return horizontal * depth
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
