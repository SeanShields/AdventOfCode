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

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}

func solvePart1(input string) int {
	floor := 0

	for _, paren := range input {
		if paren == '(' {
			floor++
		} else if paren == ')' {
			floor--
		} else {
			continue
		}
	}

	return floor
}

func solvePart2(input string) int {
	floor := 0

	for index, paren := range input {
		if floor == -1 {
			return index
		} else if paren == '(' {
			floor++
		} else if paren == ')' {
			floor--
		} else {
			continue
		}
	}

	panic("position not found")
}
