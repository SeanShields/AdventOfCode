package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"regexp"
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
	number := 0
	r, _ := regexp.Compile("/00000{5,}./g")
	for {
		data := []rune(input)
		data = append(data, rune(number))
		hash := md5.Sum(toBytes(data))
		str := fmt.Sprintf("%s", hash)
		if r.MatchString(str) {
			return str
		}
		number++
	}
}

func solvePart2(input string) string {
	return "Not Implemented"
}

func toBytes(runes []rune) []byte {
	arr := []byte{}
	for _, r := range runes {
		arr = append(arr, byte(r))
	}
	return arr
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
