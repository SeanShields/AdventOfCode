package main

import (
	"fmt"
	"os"
	"sort"
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
	elves := parseCals(input)
	sort.Ints(elves)
	return elves[len(elves)-1]
}

func solvePart2(input string) int {
	elves := parseCals(input)
	sort.Ints(elves)
	return sum(elves[len(elves)-3:])
}

func parseCals(input string) []int {
	elves := []int{}
	calories := strings.Split(input, "\r\n")
	elf := 0
	for i, calorie := range calories {
		calorie, _ := strconv.Atoi(calorie)
		elf += calorie
		if calorie > 0 {
			elves = append(elves, elf)
		}
		if calorie == 0 {
			elf = 0
		} else if i == len(calories)-1 {
			return elves
		}
	}

	return elves
}

func sum(arr []int) int {
	s := 0
	for _, a := range arr {
		s += a
	}
	return s
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
