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
	lines := strings.Split(input, "\r\n")
	illegal := []string{}
	for _, line := range lines {
		result := processIllegalChunk(line, 0, &[]string{})
		if result != "" {
			illegal = append(illegal, result)
		}
	}
	return sumIllegalCharValues(illegal)
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\r\n")
	scores := []int{}
	for _, line := range lines {
		closing := []string{}
		open := &[]string{}
		result := processIllegalChunk(line, 0, open)
		incomplete := result == ""
		if incomplete {
			for _, o := range *open {
				closing = append([]string{getClosingChar(o)}, closing...)
			}
			scores = append(scores, sumIncompleteCharValues(closing))
		}
	}
	sort.Ints(scores)
	return scores[(len(scores)-1)/2]
}

func processIllegalChunk(line string, start int, open *[]string) string {
	chars := strings.Split(line, "")
	for i := start; i < len(chars); i++ {
		if isOpeningChar(chars[i]) {
			*open = append(*open, chars[i])
			return processIllegalChunk(line, i+1, open)
		} else if chars[i] != getClosingChar((*open)[len(*open)-1]) {
			return chars[i]
		} else {
			*open = (*open)[:len(*open)-1]
		}
	}
	return ""
}

func isOpeningChar(char string) bool {
	switch char {
	case "(":
		return true
	case "[":
		return true
	case "{":
		return true
	case "<":
		return true
	}
	return false
}

func getClosingChar(openingChar string) string {
	switch openingChar {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	case "<":
		return ">"
	}
	panic("could not find closing char")
}

func sumIllegalCharValues(chars []string) int {
	sum := 0
	for _, char := range chars {
		switch char {
		case ")":
			sum += 3
		case "]":
			sum += 57
		case "}":
			sum += 1197
		case ">":
			sum += 25137
		}
	}
	return sum
}

func sumIncompleteCharValues(chars []string) int {
	sum := 0
	for _, char := range chars {
		if sum != 0 {
			sum *= 5
		}
		switch char {
		case ")":
			sum += 1
		case "]":
			sum += 2
		case "}":
			sum += 3
		case ">":
			sum += 4
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
