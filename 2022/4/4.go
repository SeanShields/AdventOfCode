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
	total := 0
	rows := strings.Split(input, "\r\n")
	for _, row := range rows {
		pair := strings.Split(row, ",")
		left := strings.Split(pair[0], "-")
		right := strings.Split(pair[1], "-")
		left1, _ := strconv.Atoi(left[0])
		left2, _ := strconv.Atoi(left[1])
		right1, _ := strconv.Atoi(right[0])
		right2, _ := strconv.Atoi(right[1])
		match := (existsInRange(left1, right1, right2) &&
			existsInRange(left2, right1, right2)) ||
			(existsInRange(right1, left1, left2) &&
				existsInRange(right2, left1, left2))
		if match {
			total++
		}
	}
	return total
}

func solvePart2(input string) int {
	total := 0
	rows := strings.Split(input, "\r\n")
	for _, row := range rows {
		pair := strings.Split(row, ",")
		left := strings.Split(pair[0], "-")
		right := strings.Split(pair[1], "-")
		left1, _ := strconv.Atoi(left[0])
		left2, _ := strconv.Atoi(left[1])
		right1, _ := strconv.Atoi(right[0])
		right2, _ := strconv.Atoi(right[1])
		match := existsInRange(left1, right1, right2) ||
			existsInRange(left2, right1, right2) ||
			existsInRange(right1, left1, left2) ||
			existsInRange(right2, left1, left2)
		if match {
			total++
		}
	}
	return total
}

func existsInRange(num int, start int, end int) bool {
	for start <= end {
		if num == start {
			return true
		}
		start++
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
