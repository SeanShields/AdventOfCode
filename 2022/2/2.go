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

var WIN_SCORE = 6
var DRAW_SCORE = 3
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
		opponent, me := parseGame(game)
		win := WIN_SCENARIOS[opponent] == me
		draw := DRAW_SCENARIOS[opponent] == me
		score := SHAPE_SCORES[me]
		if win {
			score += WIN_SCORE
		} else if draw {
			score += DRAW_SCORE
		}
		total += score
	}
	return total
}

func solvePart2(input string) int {
	games := strings.Split(input, "\r\n")
	total := 0
	for _, game := range games {
		opponent, outcome := parseGame(game)
		win := outcome == "Z"
		draw := outcome == "Y"
		score := 0
		if win {
			score += SHAPE_SCORES[WIN_SCENARIOS[opponent]] + WIN_SCORE
		} else if draw {
			score += SHAPE_SCORES[DRAW_SCENARIOS[opponent]] + DRAW_SCORE
		} else {
			score += SHAPE_SCORES[LOSS_SCENARIOS[opponent]]
		}
		total += score
	}
	return total
}

func parseGame(game string) (string, string) {
	split := strings.Split(game, " ")
	return split[0], split[1]
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
