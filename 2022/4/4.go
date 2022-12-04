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
		sections := parseSections(row)
		match := (existsInRange(sections.left1, sections.right1, sections.right2) &&
			existsInRange(sections.left2, sections.right1, sections.right2)) ||
			(existsInRange(sections.right1, sections.left1, sections.left2) &&
				existsInRange(sections.right2, sections.left1, sections.left2))
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
		sections := parseSections(row)
		match := existsInRange(sections.left1, sections.right1, sections.right2) ||
			existsInRange(sections.left2, sections.right1, sections.right2) ||
			existsInRange(sections.right1, sections.left1, sections.left2) ||
			existsInRange(sections.right2, sections.left1, sections.left2)
		if match {
			total++
		}
	}
	return total
}

type sections struct {
	left1  int
	left2  int
	right1 int
	right2 int
}

func parseSections(row string) sections {
	sections := sections{}
	pair := strings.Split(row, ",")
	left := strings.Split(pair[0], "-")
	right := strings.Split(pair[1], "-")
	sections.left1, _ = strconv.Atoi(left[0])
	sections.left2, _ = strconv.Atoi(left[1])
	sections.right1, _ = strconv.Atoi(right[0])
	sections.right2, _ = strconv.Atoi(right[1])
	return sections
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
