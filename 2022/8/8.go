package main

import (
	"fmt"
	"math"
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
	trees := parseGrid(input)
	total := 0
	for _, tree := range trees {
		if tree.visible(trees) {
			total++
		}
	}
	return total
}

func solvePart2(input string) int {
	trees := parseGrid(input)
	highest := 0
	for _, tree := range trees {
		score := tree.score(trees)
		if score > highest {
			highest = score
		}
	}
	return highest
}

type tree struct {
	height int
	x      int
	y      int
}

func (tree tree) visible(trees []tree) bool {
	length := int(math.Sqrt(float64(len(trees))))
	if tree.x == 0 || tree.x == int(length)-1 {
		return true
	}

	row := tree.getRow(trees)
	col := tree.getCol(trees)

	left := true
	l := tree.y - 1
	for l >= 0 {
		if tree.height <= row[l].height {
			left = false
		}
		l--
	}

	right := true
	r := tree.y + 1
	for r < length {
		if tree.height <= row[r].height {
			right = false
		}
		r++
	}

	up := true
	u := tree.x - 1
	for u >= 0 {
		if tree.height <= col[u].height {
			up = false
		}
		u--
	}

	down := true
	d := tree.x + 1
	for d < length {
		if tree.height <= col[d].height {
			down = false
		}
		d++
	}

	return left || right || up || down
}

func (tree tree) score(trees []tree) int {
	length := int(math.Sqrt(float64(len(trees))))
	if tree.x == 0 || tree.x == int(length)-1 {
		return 0
	}

	row := tree.getRow(trees)
	col := tree.getCol(trees)

	left := 0
	l := tree.y - 1
	for l >= 0 {
		left++
		if tree.height <= row[l].height {
			break
		}
		l--
	}

	right := 0
	r := tree.y + 1
	for r < length {
		right++
		if tree.height <= row[r].height {
			break
		}
		r++
	}

	up := 0
	u := tree.x - 1
	for u >= 0 {
		up++
		if tree.height <= col[u].height {
			break
		}
		u--
	}

	down := 0
	d := tree.x + 1
	for d < length {
		down++
		if tree.height <= col[d].height {
			break
		}
		d++
	}

	return left * right * up * down
}

func (this tree) getRow(tt []tree) []tree {
	trees := []tree{}
	for _, t := range tt {
		if this.x == t.x {
			trees = append(trees, t)
		}
	}
	return trees
}

func (this tree) getCol(tt []tree) []tree {
	trees := []tree{}
	for _, t := range tt {
		if this.y == t.y {
			trees = append(trees, t)
		}
	}
	return trees
}

func parseGrid(input string) []tree {
	trees := []tree{}
	x := 0
	y := 0
	rows := strings.Split(input, "\r\n")
	for _, row := range rows {
		for r, t := range row {
			height, _ := strconv.Atoi(string(t))
			trees = append(trees, tree{height, x, y})
			y++
			if r == len(row)-1 {
				y = 0
			}
		}
		x++
	}
	return trees
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
