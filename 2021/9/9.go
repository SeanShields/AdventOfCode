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
	rows := strings.Split(input, "\r\n")

	lowPoints := []int{}
	for r, row := range rows {
		for c, char := range row {
			top := ""
			if r != 0 {
				row := rows[r-1]
				for p, cc := range row {
					if p == c {
						top = string(cc)
					}
				}
			}

			right := ""
			if c < len(row)-1 {
				right = string(row[c+1])
			}

			bottom := ""
			if r < len(rows)-1 {
				row := rows[r+1]
				for p, cc := range row {
					if p == c {
						bottom = string(cc)
					}
				}
			}

			left := ""
			if c != 0 {
				left = string(row[c-1])
			}

			digit := int(char - '0')
			adjacent := []int{}
			if top != "" {
				t, _ := strconv.Atoi(top)
				adjacent = append(adjacent, t)
			}
			if right != "" {
				r, _ := strconv.Atoi(right)
				adjacent = append(adjacent, r)
			}
			if bottom != "" {
				b, _ := strconv.Atoi(bottom)
				adjacent = append(adjacent, b)
			}
			if left != "" {
				l, _ := strconv.Atoi(left)
				adjacent = append(adjacent, l)
			}

			if isLow(adjacent, digit) {
				lowPoints = append(lowPoints, int(char-'0'))
			}
		}
	}

	return sum(lowPoints) + len(lowPoints)
}

func solvePart2(input string) string {
	return "Not Implemented"
}

func sum(slice []int) int {
	sum := 0
	for _, s := range slice {
		sum += s
	}
	return sum
}

func isLow(slice []int, i int) bool {
	for _, s := range slice {
		if i >= s {
			return false
		}
	}
	return true
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
