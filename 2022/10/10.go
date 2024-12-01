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
	strengths := map[int]int{
		20:  0,
		60:  0,
		100: 0,
		140: 0,
		180: 0,
		220: 0,
	}
	x := 1
	q := []int{}
	for cycle, instruction := range strings.Split(input, "\r\n") {
		// we we hit an add operation, add it to the queue
		// with 3 as the number of cycles to wait before
		// add to x with the amount

		if instruction != "noop" {
			v, _ := strconv.Atoi(strings.Split(instruction, " ")[1])
			// q = append(q, 0, 0, v)
			q = append(q, 0, 0, v)
			copy(q[1:], q[0:])
			q[0] = x
		}

		if len(q) > 0 {
			popped := q[0]
			q = q[0:]
			if popped != 0 {
				x += popped
			}
		}

		_, exists := strengths[cycle+1]
		if exists {
			strengths[cycle+1] = x
		}

		fmt.Printf("%v\n", strengths)
		cycle++
	}

	return sumStrengths(strengths)
}

func solvePart2(input string) string {
	return "Not Implemented"
}

func sumStrengths(strengths map[int]int) int {
	total := 0
	for _, strength := range strengths {
		total += strength
	}
	return total
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
