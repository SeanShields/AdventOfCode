package main

import (
	"fmt"
	"os"
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

// A - rock
// B - paper
// C - scissors

// X - rock
// Y - paper
// Z - scissors

var SHAPE_SCORES = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var WIN_SCENARIOS = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var DRAW_SCENARIOS = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var LOSS_SCENARIOS = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

func solvePart1(input string) int {
	games := strings.Split(input, "\r\n")
	total := 0
	for _, game := range games {
		split := strings.Split(game, " ")
		opponent := split[0]
		me := split[1]
		win := (opponent == "A" && me == "Y") ||
			(opponent == "B" && me == "Z") ||
			(opponent == "C" && me == "X")
		draw := (opponent == "A" && me == "X") ||
			(opponent == "B" && me == "Y") ||
			(opponent == "C" && me == "Z")
		score := SHAPE_SCORES[me]
		if win {
			score += 6
		} else if draw {
			score += 3
		}
		total += score
	}
	return total
}

func solvePart2(input string) int {
	games := strings.Split(input, "\r\n")
	total := 0
	for _, game := range games {
		split := strings.Split(game, " ")
		opponent := split[0]
		outcome := split[1]
		win := outcome == "Z"
		draw := outcome == "Y"
		score := 0
		if win {
			score += SHAPE_SCORES[WIN_SCENARIOS[opponent]] + 6
		} else if draw {
			score += SHAPE_SCORES[DRAW_SCENARIOS[opponent]] + 3
		} else {
			score += SHAPE_SCORES[LOSS_SCENARIOS[opponent]]
		}
		total += score
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
