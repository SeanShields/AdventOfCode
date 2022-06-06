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
	instructions := strings.Split(input, "\n")
	registers := make(map[string]int)
	for _, instruction := range instructions {
		operationAndCondition := strings.Split(instruction, " if ")
		operation := strings.Split(operationAndCondition[0], " ")
		condition := strings.Split(operationAndCondition[1], " ")
		register := operation[0]
		direction := operation[1]
		amount, _ := strconv.Atoi(operation[2])
		conditionRegister := condition[0]
		conditionOperator := condition[1]
		conditionAmount, _ := strconv.Atoi(condition[2])

		conditionResult := false
		registerToCompare := registers[conditionRegister]
		switch conditionOperator {
		case ">":
			conditionResult = registerToCompare > conditionAmount
		case "<":
			conditionResult = registerToCompare < conditionAmount
		case ">=":
			conditionResult = registerToCompare >= conditionAmount
		case "<=":
			conditionResult = registerToCompare <= conditionAmount
		case "==":
			conditionResult = registerToCompare == conditionAmount
		case "!=":
			conditionResult = registerToCompare != conditionAmount
		}

		registerAmount := registers[register]
		if conditionResult {
			if direction == "inc" {
				registerAmount += amount
			} else {
				registerAmount -= amount
			}
		}
		registers[register] = registerAmount
	}

	largest := 0
	for _, register := range registers {
		if register > largest {
			largest = register
		}
	}
	return largest
}

func solvePart2(input string) string {
	return "Not Implemented"
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
