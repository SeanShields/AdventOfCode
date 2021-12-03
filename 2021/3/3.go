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

func solvePart1(input string) int64 {
	gamma := ""
	epsilon := ""
	numbers := strings.Split(input, "\r\n")

	i := 0
	for i < len(numbers[0]) {
		num0s := 0
		num1s := 0
		for _, number := range numbers {
			if number[i] == '0' {
				num0s++
			} else {
				num1s++
			}
		}

		if num0s > num1s {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}

		i++
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	return g * e
}

func solvePart2(input string) int64 {
	numbers := strings.Split(input, "\r\n")
	o2Rating := calcRating(numbers, 0, true)
	co2Rating := calcRating(numbers, 0, false)
	return o2Rating * co2Rating
}

func calcRating(numbers []string, i int, isO2 bool) int64 {
	keep := []string{}
	num0s := 0
	num1s := 0
	for _, number := range numbers {
		if number[i] == '0' {
			num0s++
		} else {
			num1s++
		}
	}

	keepWithLeading0 := false
	keepWithLeading1 := false
	if isO2 {
		keepWithLeading0 = num0s > num1s
		keepWithLeading1 = num1s >= num0s
	} else {
		keepWithLeading0 = num0s <= num1s
		keepWithLeading1 = num1s < num0s
	}
	for _, number := range numbers {
		if (keepWithLeading0 && number[i] == '0') || (keepWithLeading1 && number[i] == '1') {
			keep = append(keep, number)
		}
	}

	if len(keep) == 1 {
		rating, _ := strconv.ParseInt(keep[0], 2, 64)
		return rating
	}

	return calcRating(keep, i+1, isO2)
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
