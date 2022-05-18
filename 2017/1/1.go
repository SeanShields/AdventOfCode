package main

import (
	"fmt"
	"os"
	"strconv"
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
	sum := 0
	for i, n := range input {
		num, _ := strconv.Atoi(string(n))
		numToCompareIndex := getNextIndex(i, len(input))
		numToCompare, _ := strconv.Atoi(string(input[numToCompareIndex]))
		if num == numToCompare {
			sum += num
		}
	}
	return sum
}

func solvePart2(input string) int {
	sum := 0
	for i, n := range input {
		num, _ := strconv.Atoi(string(n))
		numToCompareIndex := getHalfwayIndex(i, len(input))
		numToCompare, _ := strconv.Atoi(string(input[numToCompareIndex]))
		if num == numToCompare {
			sum += num
		}
	}
	return sum
}

func getNextIndex(index int, len int) int {
	i := 0
	if index < len-1 {
		i = index + 1
	}
	return i
}

func getHalfwayIndex(index int, len int) int {
	half := len / 2
	for i := 0; i < half; i++ {
		if index == len-1 {
			index = 0
		} else {
			index++
		}
	}
	return index
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
