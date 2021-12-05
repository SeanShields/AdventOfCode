package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// fabric := [100000]int{}
	// gridSize = 1000000
	testInstructions := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	testFabric := []int{}
	rowCounter := 1
	for i := 0; i < len(testInstructions); i++ {
		number := parseNumber(testInstructions[i])
		startXPosition, _ := strconv.Atoi(parseStartPosition(testInstructions[i]))
		startYPosition, _ := strconv.Atoi(parseEndPosition(testInstructions[i]))
		xSize, _ := strconv.Atoi(parseXSize(testInstructions[i]))
		ySize, _ := strconv.Atoi(parseYSize(testInstructions[i]))
		startIndex := getStartIndex(startXPosition, startYPosition)
		markNumberOnFabric(testFabric, startIndex, number, xSize, ySize)
		rowCounter++
	}

	fmt.Println(testFabric)

	// fmt.Printf("result: %d", getTwoOrMoreClaims(testFabric))
}

func markNumberOnFabric(fabric []int, startIndex int, number int, xSize int, ySize int) {
	rowCounter := startIndex
	for i := startIndex; i < len(fabric); i += counter {
		index := i

		fabric[index] = number
	}
}

func atFabricEdge(fabric []int, index int) bool {
	return len(fabric) % fabric[index]
	for {

	}
}

func getStartIndex(startXPosition, startYPosition int) int {
	return startXPosition + startYPosition*8
}

func parseNumber(instruction string) int {
	r, _ := regexp.Compile("#\\d*")
	number, _ := strconv.Atoi(strings.ReplaceAll(r.FindString(instruction), "#", ""))
	return number
}

func parseStartPosition(instruction string) string {
	r, _ := regexp.Compile("\\d*,\\d*")
	return strings.Split(r.FindString(instruction), ",")[0]
}

func parseEndPosition(instruction string) string {
	r, _ := regexp.Compile("\\d*,\\d*")
	return strings.Split(r.FindString(instruction), ",")[1]
}

func parseXSize(instruction string) string {
	r, _ := regexp.Compile("\\d*x\\d*")
	return strings.Split(r.FindString(instruction), "x")[0]
}

func parseYSize(instruction string) string {
	r, _ := regexp.Compile("\\d*x\\d*")
	return strings.Split(r.FindString(instruction), "x")[1]
}

func getTwoOrMoreClaims(fabric []int) int {
	count := 0
	for _, position := range fabric {
		if position >= 2 {
			count++
		}
	}
	return count
}
