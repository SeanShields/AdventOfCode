package main

import (
	"fmt"
	"os"
	"regexp"
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
	total := 0
	rucksacks := strings.Split(input, "\r\n")
	for _, rucksack := range rucksacks {
		first := rucksack[:len(rucksack)/2]
		second := rucksack[len(rucksack)/2:]
		match := regexp.MustCompile("[" + first + "]").FindString(second)
		total += getPriority(match)
	}
	return total
}

func solvePart2(input string) string {
	return "Not Implemented"
}

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getPriority(item string) int {
	priority := 1
	for _, letter := range letters {
		if string(letter) == item {
			return priority
		}
		priority++
	}
	panic("Failed to get priority")
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
