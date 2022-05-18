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
	lines := strings.Split(input, "\r\n")
	differences := []int{}
	for _, line := range lines {
		nums := toInt(strings.Split(line, "\t"))
		sort.Ints(nums)
		differences = append(differences, nums[len(nums)-1]-nums[0])
	}
	return sum(differences)
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\r\n")
	differences := []int{}
	for _, line := range lines {
		nums := toInt(strings.Split(line, "\t"))
		for i, x := range nums {
			for ii, y := range nums {
				if i == ii || x == 0 || y == 0 {
					continue
				}
				if x%y == 0 {
					differences = append(differences, x/y)
				}
			}
		}
	}
	return sum(differences)
}

func sum(arr []int) int {
	s := 0
	for _, a := range arr {
		s += a
	}
	return s
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
