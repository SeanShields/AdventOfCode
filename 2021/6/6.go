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
	fish := strings.Split(input, ",")

	lanternfish := toInt(fish)
	add := 0
	i := 0
	for i < 80 {
		for f, fish := range lanternfish {
			lanternfish[f]--
			if fish == 0 {
				add++
				lanternfish[f] = 6
			}
		}

		for j := 0; j < add; j++ {
			lanternfish = append(lanternfish, 8)
		}

		add = 0
		i++
	}

	return len(lanternfish)
}

func solvePart2(input string) int64 {
	fish := strings.Split(input, ",")

	lanternfish := toInt(fish)
	counts := make(map[int]int64)
	for _, fish := range lanternfish {
		counts[fish]++
	}

	i := 0
	for i < 256 {
		zeroes := counts[0]
		f := 0
		for f <= 8 {
			if f == 8 {
				counts[f] = zeroes
				f++
				continue
			}
			counts[f] = counts[f+1]
			if f == 6 {
				counts[f] += zeroes
			}
			f++
		}
		i++
	}

	return sum(counts)
}

func sum(counts map[int]int64) int64 {
	var sum int64
	f := 0
	for f <= 8 {
		sum += counts[f]
		f++
	}
	return sum
}

func toInt(slice []string) []int {
	ret := []int{}
	for _, s := range slice {
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
