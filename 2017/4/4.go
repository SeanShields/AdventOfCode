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
	valid := 0
	phrases := strings.Split(input, "\r\n")
	for _, phrase := range phrases {
		words := strings.Split(phrase, " ")
		if hasDuplicateWords(words) {
			continue
		}
		valid++
	}
	return valid
}

func solvePart2(input string) string {
	return "Not Implemented"
}

func hasDuplicateWords(words []string) bool {
	for i, x := range words {
		for o, y := range words {
			if i != o && x == y {
				return true
			}
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
