package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
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
	lines := strings.Split(input, "\r\n")
	caves := initCaves(lines)
	sort.Slice(caves, func(i, j int) bool {
		return caves[i].name == "start" || caves[j].name == "end"
	})
	// fmt.Printf("caves: %v\n", caves)

	paths := [][]string{}
	path := getPath(caves, "start")

	fmt.Printf("path: %v\n", path)

	return len(paths)
}

func solvePart2(input string) string {
	return "Not Implemented"
}

type Caves []Cave
type Cave struct {
	name    string
	edges   []string
	visited bool
}

func getPath(caves Caves, vertex string) []string {
	path := []string{vertex}
	for _, cave := range caves {
		edges := cave.getEdges(caves)
		sort.Slice(edges, func(i, j int) bool {
			return edges[i].name == "end"
		})
	}
	return path
}

func (cave Cave) getEdges(caves Caves) Caves {
	edges := Caves{}
	for _, e := range cave.edges {
		for _, cave := range caves {
			if e == cave.name {
				edges = append(edges, cave)
			}
		}
	}
	return edges
}

func initCaves(lines []string) Caves {
	caves := Caves{}
	for _, l := range lines {
		points := strings.Split(l, "-")
		for _, p := range points {
			exists := false
			for _, c := range caves {
				if p == c.name {
					exists = true
				}
			}
			if !exists {
				adjacents := getAdjacents(lines, p)
				cave := Cave{p, adjacents, false}
				caves = append(caves, cave)
			}
		}
	}
	return caves
}

func getAdjacents(points []string, name string) []string {
	adjacents := []string{}
	for _, c := range points {
		z := strings.Split(c, "-")
		for _, b := range z {
			if b == name {
				a := strings.Replace(c, name, "", 1)
				x := strings.Replace(a, "-", "", 1)
				adjacents = append(adjacents, x)
			}
		}
	}
	return adjacents
}

func (cave Cave) isSmall() bool {
	r, _ := regexp.Compile("[a-z]")
	return r.MatchString(cave.name)
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
