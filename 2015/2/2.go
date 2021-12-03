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

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}

func solvePart1(input string) int {
	dimensions := strings.Split(input, "\r\n")
	total := 0

	for _, dimension := range dimensions {
		if dimension == "" {
			continue
		}

		lwh := strings.Split(dimension, "x")
		l, _ := strconv.Atoi(lwh[0])
		w, _ := strconv.Atoi(lwh[1])
		h, _ := strconv.Atoi(lwh[2])
		total += (2 * l * w) + (2 * w * h) + (2 * h * l)
		arr := []int{l, w, h}
		sort.Ints(arr)
		total += arr[0] * arr[1]
	}

	return total
}

func solvePart2(input string) int {
	dimensions := strings.Split(input, "\r\n")
	total := 0

	for _, dimension := range dimensions {
		if dimension == "" {
			continue
		}

		lwh := strings.Split(dimension, "x")
		l, _ := strconv.Atoi(lwh[0])
		w, _ := strconv.Atoi(lwh[1])
		h, _ := strconv.Atoi(lwh[2])
		perimeters := []int{2 * (l + w), 2 * (w + h), 2 * (h + l)}
		sort.Ints(perimeters)
		total += perimeters[0] + (l * w * h)
	}

	return total
}
