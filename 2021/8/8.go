package main

import (
	"fmt"
	"os"
	"sort"
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
	entries := []string{}
	for _, line := range lines {
		entries = append(entries, line)
	}

	count := 0
	for _, e := range entries {
		notes := strings.Split(e, " | ")[1]
		for _, note := range strings.Split(notes, " ") {
			number := calcSimpleNumber(note)
			if number != "" {
				count++
			}
		}
	}

	return count
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\r\n")
	entries := []string{}
	for _, line := range lines {
		entries = append(entries, line)
	}

	output := 0
	for _, e := range entries {
		s := strings.Split(e, " | ")
		signals := decodeSignals(strings.Split(s[0], " "))
		o := calcOutput(signals, strings.Split(s[1], " "))
		output += o
	}
	return output
}

func decodeSignals(signals []string) map[int]string {
	ret := make(map[int]string)

	// unique length: 1, 4, 7, 8
	// 5 chars: 2, 3, 5
	// 6 chars: 0, 6, 9

	// 1
	for _, signal := range signals {
		if len(signal) == 2 {
			ret[1] = signal
			signals = remove(signals, signal)
			break
		}
	}

	// 4
	for _, signal := range signals {
		if len(signal) == 4 {
			ret[4] = signal
			signals = remove(signals, signal)
			break
		}
	}

	// 7
	for _, signal := range signals {
		if len(signal) == 3 {
			ret[7] = signal
			signals = remove(signals, signal)
			break
		}
	}

	// 8
	for _, signal := range signals {
		if len(signal) == 7 {
			ret[8] = signal
			signals = remove(signals, signal)
			break
		}
	}

	// 3
	for _, signal := range signals {
		if len(signal) == 5 && substring(signal, ret[1]) {
			ret[3] = signal
			signals = remove(signals, signal)
			break
		}
	}

	// 6
	for _, signal := range signals {
		if len(signal) == 6 && !substring(signal, ret[1]) {
			ret[6] = signal
			signals = remove(signals, signal)
			break
		}
	}

	// 5
	for _, signal := range signals {
		if len(signal) == 5 && substring(ret[6], signal) {
			ret[5] = signal
			signals = remove(signals, signal)
			break
		}
	}

	// 2
	for _, signal := range signals {
		if len(signal) == 5 {
			ret[2] = signal
			signals = remove(signals, signal)
			break
		}
	}

	// 9
	for _, signal := range signals {
		if len(signal) == 6 && substring(signal, ret[5]) {
			ret[9] = signal
			signals = remove(signals, signal)
			break
		}
	}

	// 0
	ret[0] = signals[0]

	return ret
}

func calcOutput(signals map[int]string, notes []string) int {
	chars := ""
	for _, note := range notes {
		chars += calcNumberFromSignal(signals, note)
	}
	output, _ := strconv.Atoi(chars)
	return output
}

func calcNumberFromSignal(signals map[int]string, note string) string {
	for number := 0; number < 10; number++ {
		if sortString(note) == sortString(signals[number]) {
			return strconv.Itoa(number)
		}
	}
	panic("number could not be calculated")
}

func calcSimpleNumber(note string) string {
	if len(note) == 2 {
		return "1"
	} else if len(note) == 3 {
		return "7"
	} else if len(note) == 4 {
		return "4"
	} else if len(note) == 7 {
		return "8"
	} else {
		return ""
	}
}

func intExists(slice []int, i int) bool {
	for _, s := range slice {
		if s == i {
			return true
		}
	}
	return false
}

func charExists(source string, i rune) bool {
	for _, s := range source {
		if s == i {
			return true
		}
	}
	return false
}

func substring(source string, substring string) bool {
	for _, c := range substring {
		if !charExists(source, c) {
			return false
		}
	}
	return true
}

func remove(slice []string, s string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[i] == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
