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

type node struct {
	name     string
	weight   int
	children []string
}

func solvePart1(input string) string {
	tree := parseNodes(strings.Split(input, "\r\n"))
	for _, node := range tree {
		if !isChild(node, tree) {
			return node.name
		}
	}
	return "Not Found"
}

func solvePart2(input string) string {
	return "Not Implemented"
}

func isChild(current node, tree []node) bool {
	isChild := false
	for _, node := range tree {
		for _, child := range node.children {
			if child == current.name {
				isChild = true
			}
		}
	}
	return isChild
}

func parseNodes(input []string) []node {
	tree := []node{}
	for _, mapping := range input {
		info := strings.Split(mapping, " -> ")
		nameAndWeight := strings.Split(info[0], " ")
		name := nameAndWeight[0]
		weight, _ := strconv.Atoi(
			strings.Replace(strings.Replace(nameAndWeight[1], "(", "", 1), ")", "", 1),
		)

		node := node{name: name, weight: weight}
		hasChildren := len(info) == 2
		if hasChildren {
			node.children = strings.Split(info[1], ", ")
		}
		tree = append(tree, node)
	}
	return tree
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
