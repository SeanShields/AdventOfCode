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

	part1 := part1(input)
	fmt.Printf("Part 1: %v\n", part1)

	part2 := part2(input)
	fmt.Printf("Part 2: %v", part2)
}

func part1(input string) string {
	larger := 0
	depths := strings.Split(input, "\r\n")

	for i := 0; i < len(depths); i++ {
		if i == 0 || depths[i] == "" {
			continue
		}

		currentDepth, _ := strconv.Atoi(depths[i])
		previousDepth, _ := strconv.Atoi(depths[i-1])
		increasing := currentDepth > previousDepth
		if increasing {
			larger++
		}
	}

	return strconv.Itoa(larger)
}

func part2(input string) string {
	larger := 0
	depthsInput := strings.Split(input, "\r\n")
	depths := toInt(depthsInput)

	for i := 0; i < len(depths); i++ {
		if i == 0 || i+3 >= len(depths) || depths[i] == 0 {
			continue
		}

		lastIndexInCurrentWindow := i + 3
		firstIndexInPreviousWindow := i - 1
		lastIndexInPreviousWindow := i + 2
		currentWindow := depths[i:lastIndexInCurrentWindow]
		previousWindow := depths[firstIndexInPreviousWindow:lastIndexInPreviousWindow]

		increasing := sum(currentWindow) > sum(previousWindow)
		if increasing {
			larger++
		}
	}

	return strconv.Itoa(larger)
}

func toInt(arr []string) []int {
	ret := []int{}
	for _, s := range arr {
		i, _ := strconv.Atoi(s)
		ret = append(ret, i)
	}
	return ret
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
