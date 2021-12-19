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
	entries := []string{}
	for _, line := range lines {
		entries = append(entries, line)
	}

	count := 0
	for _, e := range entries {
		notes := strings.Split(e, " | ")[1]
		for _, note := range strings.Split(notes, " ") {
			if exists([]int{2, 3, 4, 7}, len(note)) {
				count++
			}
		}
	}

	return count
}

func solvePart2(input string) string {
	return "Not Implemented"
}

func exists(slice []int, i int) bool {
	for _, s := range slice {
		if s == i {
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
