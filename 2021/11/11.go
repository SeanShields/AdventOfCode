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
	lines := strings.Split(input, "\r\n")
	grid := initGrid(lines)
	flashes := 0
	steps := 100
	width := 10
	for i := 0; i < steps; i++ {
		for o, _ := range grid {
			grid.increase(o, width)
		}
		flashes += grid.sumFlashed()
		grid.resetFlashed()
	}

	return flashes
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\r\n")
	grid := initGrid(lines)
	width := 10
	i := 0
	for {
		for o, _ := range grid {
			grid.increase(o, width)
		}
		flashes := grid.sumFlashed()
		if flashes == len(grid) {
			return i + 1
		}
		grid.resetFlashed()
		i++
	}
}

type Grid []Octopus
type Octopus struct {
	energy  int
	row     int
	col     int
	flashed bool
}

func (grid Grid) print(width int) {
	i := 0
	for i < width {
		fmt.Printf("%v\n", grid[i*width:i*width+width].selectDigits())
		i++
	}
	fmt.Printf("\n")
}

func (grid Grid) selectDigits() []int {
	digits := []int{}
	for _, o := range grid {
		digits = append(digits, o.energy)
	}
	return digits
}

func (grid *Grid) increase(i int, width int) {
	for o, _ := range *grid {
		if (*grid)[o].flashed || o != i {
			continue
		}

		// increase energy
		(*grid)[o].energy++

		// if energy is greater than 9, reset to 0
		if (*grid)[o].energy > 9 {
			(*grid)[o].energy = 0
			(*grid).flash(o, width)
		}
	}
}

func (grid *Grid) flash(i int, width int) {
	(*grid)[i].flashed = true
	adjacents := (*grid)[i].getAdjacents(*grid, width)
	for _, a := range adjacents {
		for g, oct := range *grid {
			if oct.row == a.row && oct.col == a.col {
				(*grid).increase(g, width)
			}
		}
	}
}

func (octopus Octopus) getAdjacents(grid Grid, width int) []Octopus {
	adjacents := []Octopus{}

	if octopus.row != 0 {
		top := grid.getOctopus(octopus.row-1, octopus.col)
		adjacents = append(adjacents, top)

		if octopus.col != 0 {
			topLeft := grid.getOctopus(octopus.row-1, octopus.col-1)
			adjacents = append(adjacents, topLeft)
		}

		if octopus.col < width-1 {
			topRight := grid.getOctopus(octopus.row-1, octopus.col+1)
			adjacents = append(adjacents, topRight)
		}
	}

	if octopus.row < width-1 {
		bottom := grid.getOctopus(octopus.row+1, octopus.col)
		adjacents = append(adjacents, bottom)

		if octopus.col != 0 {
			bottomLeft := grid.getOctopus(octopus.row+1, octopus.col-1)
			adjacents = append(adjacents, bottomLeft)
		}

		if octopus.col < width-1 {
			bottomRight := grid.getOctopus(octopus.row+1, octopus.col+1)
			adjacents = append(adjacents, bottomRight)
		}
	}

	if octopus.col < width-1 {
		right := grid.getOctopus(octopus.row, octopus.col+1)
		adjacents = append(adjacents, right)
	}

	if octopus.col > 0 {
		left := grid.getOctopus(octopus.row, octopus.col-1)
		adjacents = append(adjacents, left)
	}

	return adjacents
}

func (grid Grid) getOctopus(row int, col int) Octopus {
	for _, octopus := range grid {
		if octopus.row == row && octopus.col == col {
			return octopus
		}
	}
	panic("Failed to find Octopus")
}

func (grid *Grid) resetFlashed() {
	for i, _ := range *grid {
		(*grid)[i].flashed = false
	}
}

func (grid Grid) sumFlashed() int {
	sum := 0
	for _, octopus := range grid {
		if octopus.flashed {
			sum++
		}
	}
	return sum
}

func initGrid(lines []string) Grid {
	grid := Grid{}
	for l, line := range lines {
		for c, num := range line {
			energy, _ := strconv.Atoi(string(num))
			grid = append(grid, Octopus{energy, l, c, false})
		}
	}
	return grid
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
