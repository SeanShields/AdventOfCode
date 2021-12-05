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
	numbers := strings.Split(lines[0], ",")
	boards := parseBoards(lines)
	for _, n := range numbers {
		number, _ := strconv.Atoi(n)
		for _, board := range boards {
			board.mark(number)
			if board.isBingo() {
				return number * board.sumUnmarkedSquares()
			}
		}
	}

	panic("winner not found")
}

func solvePart2(input string) string {
	return "Not Implemented"
}

type Board []Square
type Square struct {
	number int
	marked bool
}

func (board *Board) mark(number int) {
	for i, square := range *board {
		if square.number == number {
			(*board)[i].marked = true
		}
	}
}

func (board Board) isBingo() bool {
	return board.hasCompletedRow() || board.hasCompletedColumn()
}

func (board Board) hasCompletedRow() bool {
	for i := 0; i < len(board); i += 5 {
		completed := true
		for r := 0; r < i+5; r++ {
			if !board[r].marked {
				completed = false
			}
		}
		if completed {
			return true
		}
	}
	return false
}

func (board Board) hasCompletedColumn() bool {
	for i := 0; i < 5; i++ {
		completed := true
		for r := i; r <= i+20; r += 5 {
			if !board[r].marked {
				completed = false
			}
		}
		if completed {
			return true
		}
	}
	return false
}

func (board Board) sumUnmarkedSquares() int {
	sum := 0
	for _, square := range board {
		if !square.marked {
			sum += square.number
		}
	}
	return sum
}

func parseBoards(input []string) []Board {
	boards := []Board{}
	board := Board{}

	for i, line := range input {
		if i <= 1 || line == "" {
			continue
		}

		squares := []Square{}
		numbers := strings.Split(line, " ")
		for _, number := range numbers {
			if number != "" {
				n, _ := strconv.Atoi(number)
				squares = append(squares, Square{n, false})
			}
		}

		board = append(board, squares...)
		if i == len(input)-1 || input[i+1] == "" {
			boards = append(boards, board)
			board = []Square{}
		}
	}

	return boards
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
