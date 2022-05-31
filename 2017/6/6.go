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
	blocks := toInt(strings.Split(input, "\t"))
	states := [][]int{}
	i := 0
	for {
		state := redistribute(blocks)
		i++
		if exists(state, states) {
			return i
		}
		s := []int{}
		for _, st := range state {
			s = append(s, st)
		}
		states = append(states, s)
	}
}

func solvePart2(input string) string {
	return "Not Implemented"
}

func redistribute(blocks []int) []int {
	index := indexAtLargest(blocks)
	redistributionAmount := blocks[index]
	blocks[index] = 0

	i := index + 1
	if i >= len(blocks) {
		i = 0
	}

	for redistributionAmount > 0 {
		blocks[i]++
		if i == len(blocks)-1 {
			i = 0
		} else {
			i++
		}
		redistributionAmount--
	}
	return blocks
}

func exists(x []int, y [][]int) bool {
	for _, yy := range y {
		if sliceToInt(x) == sliceToInt(yy) {
			return true
		}
	}
	return false
}

func indexAtLargest(blocks []int) int {
	largestBlock := 0
	indexOfLargest := len(blocks) - 1
	for i, block := range blocks {
		if block > largestBlock {
			largestBlock = block
			indexOfLargest = i
		} else if block == largestBlock {
			if i < indexOfLargest {
				indexOfLargest = i
			}
		}
	}
	return indexOfLargest
}

func toInt(arr []string) []int {
	ret := []int{}
	for _, s := range arr {
		i, _ := strconv.Atoi(s)
		ret = append(ret, i)
	}
	return ret
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
