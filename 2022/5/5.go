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
	return getTop(crates)
}

func solvePart2(input string) string {
	crates, instructions := parse(input)
	for _, instruction := range instructions {
		crates.movePart2(instruction.amount, instruction.from, instruction.to)
	}
	return getTop(crates)
}

func parse(input string) (crates, []instruction) {
	return initCrates(input), initInstructions(input)
}

func initCrates(input string) crates {
	crates := crates{}
	lines := strings.Split(input, "\r\n")
	num := 1
	i := 1
	for i < len(lines[0]) {
		crates[num] = Stack()
		i += 4
		num++
	}

	reCrates := regexp.MustCompile(`\[[A-Z]\]`)
	for _, line := range lines {
		if !reCrates.MatchString(line) {
			break
		}
		num = 1
		i = 1
		for i < len(line) {
			char := string(line[i])
			if regexp.MustCompile(`[A-Z]`).MatchString(char) {
				crates[num].Push(char)
			}
			i += 4
			num++
		}
	}

	for _, crate := range crates {
		reversed := []string{}
		for crate.Length() > 0 {
			reversed = append(reversed, crate.Pop())
		}
		for _, reverse := range reversed {
			crate.Push(reverse)
		}
	}

	return crates
}

func initInstructions(input string) []instruction {
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
	return instructions
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
	for i <= len(crates) {
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
