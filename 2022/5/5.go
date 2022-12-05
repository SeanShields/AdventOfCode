package main

import (
	"fmt"
	"os"
	"regexp"
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

func solvePart1(input string) string {
	crates, instructions := parse(input)
	for _, instruction := range instructions {
		crates.movePart1(instruction.amount, instruction.from, instruction.to)
	}
	top := getTop(crates)
	return top
}

func solvePart2(input string) string {
	crates, instructions := parse(input)
	for _, instruction := range instructions {
		crates.movePart2(instruction.amount, instruction.from, instruction.to)
	}
	top := getTop(crates)
	return top
}

func parse(input string) (crates, []instruction) {
	crates := crates{
		1: Stack(),
		2: Stack(),
		3: Stack(),
		4: Stack(),
		5: Stack(),
		6: Stack(),
		7: Stack(),
		8: Stack(),
		9: Stack(),
	}

	crates[1].Push("T")
	crates[1].Push("D")
	crates[1].Push("W")
	crates[1].Push("Z")
	crates[1].Push("V")
	crates[1].Push("P")

	crates[2].Push("L")
	crates[2].Push("S")
	crates[2].Push("W")
	crates[2].Push("V")
	crates[2].Push("F")
	crates[2].Push("J")
	crates[2].Push("D")

	crates[3].Push("Z")
	crates[3].Push("M")
	crates[3].Push("L")
	crates[3].Push("S")
	crates[3].Push("V")
	crates[3].Push("T")
	crates[3].Push("B")
	crates[3].Push("H")

	crates[4].Push("R")
	crates[4].Push("S")
	crates[4].Push("J")

	crates[5].Push("C")
	crates[5].Push("Z")
	crates[5].Push("B")
	crates[5].Push("G")
	crates[5].Push("F")
	crates[5].Push("M")
	crates[5].Push("L")
	crates[5].Push("W")

	crates[6].Push("Q")
	crates[6].Push("W")
	crates[6].Push("V")
	crates[6].Push("H")
	crates[6].Push("Z")
	crates[6].Push("R")
	crates[6].Push("G")
	crates[6].Push("B")

	crates[7].Push("V")
	crates[7].Push("J")
	crates[7].Push("P")
	crates[7].Push("C")
	crates[7].Push("B")
	crates[7].Push("D")
	crates[7].Push("N")

	crates[8].Push("P")
	crates[8].Push("T")
	crates[8].Push("B")
	crates[8].Push("Q")

	crates[9].Push("H")
	crates[9].Push("G")
	crates[9].Push("Z")
	crates[9].Push("R")
	crates[9].Push("C")

	instructions := []instruction{}
	reInstructions := regexp.MustCompile(`move \d+ from \d+ to \d+`)
	for _, line := range strings.Split(input, "\r\n") {
		if !reInstructions.MatchString(line) {
			continue
		}
		matches := regexp.MustCompile(`\d+`).FindAllString(line, -1)
		amount, _ := strconv.Atoi(matches[0])
		from, _ := strconv.Atoi(matches[1])
		to, _ := strconv.Atoi(matches[2])
		instructions = append(instructions, instruction{amount, from, to})
	}

	return crates, instructions
}

func (c *crates) movePart1(amount int, from int, to int) {
	for amount > 0 {
		(*c)[to].Push((*c)[from].Pop())
		amount--
	}
}

func (c *crates) movePart2(amount int, from int, to int) {
	moving := []string{}
	for amount > 0 {
		moving = append(moving, (*c)[from].Pop())
		amount--
	}
	for _, move := range reverse(moving) {
		(*c)[to].Push(move)
	}
}

func getTop(crates crates) string {
	top := ""
	i := 1
	for i <= 3 {
		top += crates[i].Peek()
		i++
	}
	return top
}

func reverse(arr []string) []string {
	result := []string{}
	i := len(arr)
	for i > 0 {
		result = append(result, arr[i-1])
		i--
	}
	return result
}

type crates map[int]stack
type stack struct {
	Push   func(string)
	Pop    func() string
	Peek   func() string
	Length func() int
}
type instruction struct {
	amount int
	from   int
	to     int
}

func Stack() stack {
	slice := make([]string, 0)
	return stack{
		Push: func(i string) {
			slice = append(slice, i)
		},
		Pop: func() string {
			if len(slice) == 0 {
				panic("popping from empty stack")
			}
			result := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			return result
		},
		Peek: func() string {
			if len(slice) == 0 {
				panic("peeking an empty stack")
			}
			return slice[len(slice)-1]
		},
		Length: func() int {
			return len(slice)
		},
	}
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
