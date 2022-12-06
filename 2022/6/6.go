package main

import (
	"fmt"
	"os"
	"sort"
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
	return locate(input, 4)
}

func solvePart2(input string) int {
	return locate(input, 14)
}

func locate(input string, padding int) int {
	i := padding
	for i < len(input) {
		unique := unique(input[i-padding : i])
		if unique {
			return i
		}
		i++
	}
	panic("failed to locate position")
}

func unique(s string) bool {
	s = sortString(s)
	i := 1
	for i < len(s) {
		if string(s[i]) == string(s[i-1]) {
			return false
		}
		i++
	}
	return true
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
