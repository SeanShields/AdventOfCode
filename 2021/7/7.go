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
	c := strings.Split(input, ",")
	positions := toInt(c)

	fuelOptions := []int{}
	for _, position := range positions {
		fuel := 0
		for _, p := range positions {
			fuel += max(position, p) - min(position, p)
		}
		fuelOptions = append(fuelOptions, fuel)
	}

	sort.Ints(fuelOptions)
	return fuelOptions[0]
}

func solvePart2(input string) string {
	return "Not Implemented"
}

func toInt(slice []string) []int {
	ret := []int{}
	for _, s := range slice {
		i, _ := strconv.Atoi(s)
		ret = append(ret, i)
	}
	return ret
}

func max(x int, y int) int {
	if x > y {
		return x
	} else if y > x {
		return y
	} else {
		return x
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	} else if y < x {
		return y
	} else {
		return x
	}
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
