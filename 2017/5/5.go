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

	part1Answer := solvePart1(input)
	fmt.Printf("Part 1 Answer: %v\n", part1Answer)

	part2Answer := solvePart2(input)
	fmt.Printf("Part 2 Answer: %v", part2Answer)
}

func solvePart1(input string) int {
	steps := 0
	instructions := toInt(strings.Split(input, "\r\n"))
	position := 0
	for {
		if position < 0 || position >= len(instructions) {
			return steps
		}
		current := position
		position += instructions[position]
		instructions[current]++
		steps++
	}
}

func solvePart2(input string) int {
	steps := 0
	instructions := toInt(strings.Split(input, "\r\n"))
	position := 0
	for {
		if position < 0 || position >= len(instructions) {
			return steps
		}
		current := position
		position += instructions[position]
		if instructions[current] >= 3 {
			instructions[current]--
		} else {
			instructions[current]++
		}
		steps++
	}
}

func toInt(arr []string) []int {
	ret := []int{}
	for _, s := range arr {
		i, _ := strconv.Atoi(s)
		ret = append(ret, i)
	}
	return ret
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
