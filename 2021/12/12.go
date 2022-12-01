package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Caves []Cave
type Cave struct {
	name    string
	edges   []string
	visited bool
}

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

	caves.reset()

	for _, cave := range caves {
		fmt.Printf("%v: %v\n", cave.name, cave)
	}

	return 0
}

func solvePart2(input string) string {
	return "Not Implemented"
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

func getCave(name string, caves Caves) Cave {
	for _, cave := range caves {
		if cave.name == name {
			return cave
		}
	}
	panic("getCave failed to find cave")
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

func (caves *Caves) reset() {
	for _, c := range *caves {
		c.visited = true
	}
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
