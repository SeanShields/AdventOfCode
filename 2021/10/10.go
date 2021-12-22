package main

import (
	"fmt"
	"os"
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
	//sum := 0
	for _, line := range lines {
		illegal := ""
		for _, char := range line {

		}
	}
	return 0
}

func solvePart2(input string) string {
	return "Not Implemented"
}

func sumIllegalCharValues(chars string) int {
	sum := 0
	for _, char := range chars {
		switch char {
		case ')':
			sum += 3
		case ']':
			sum += 57
		case '}':
			sum += 1197
		case '>':
			sum += 25137
		}
	}
	return sum
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
