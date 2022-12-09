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

func solvePart1(input string) int {
	dirs := parseTree(input)
	totals := []int{}
	for _, dir := range dirs {
		total := getDirSize(dir, dirs)
		totals = append(totals, total)
	}
	return sum(totals, 100000)
}

func solvePart2(input string) int {
	dirs := parseTree(input)
	required := 30000000
	remove := required - getAvailableSpace(dirs)
	eligible := []int{}
	for _, dir := range dirs {
		size := getDirSize(dir, dirs)
		if size >= remove {
			eligible = append(eligible, size)
		}
	}
	return min(eligible)
}

type file struct {
	name string
	size int
}
type dir struct {
	id     int
	parent int
	name   string
	dirs   []string
	files  []file
}

func getAvailableSpace(dirs []dir) int {
	used := 0
	for _, dir := range dirs {
		if dir.name == "/" {
			used = getDirSize(dir, dirs)
		}
	}
	return 70000000 - used
}

func getDirSize(dir dir, dirs []dir) int {
	size := 0
	for _, file := range dir.files {
		size += file.size
	}
	for _, child := range dir.dirs {
		size += getDirSize(findChild(child, dir.id, dirs), dirs)
	}
	return size
}

func parseTree(input string) []dir {
	lines := strings.Split(input, "\r\n")
	id := 1
	root := dir{id: 0, parent: -1, name: "/"}
	dirs := []dir{root}
	current := root
	for i, line := range lines {
		if i == 0 {
			continue
		}
		directory := directory(line)
		cd := cd(line)
		ls := ls(line)
		if command(line) {
			if ls {
				continue
			}
			if cd == ".." {
				parent := findParent(current, dirs)
				current = parent
			} else {
				child := findChild(cd, current.id, dirs)
				current = child
			}
		} else {
			if directory != "" {
				new := dir{id: id, name: directory, parent: current.id}
				current.dirs = append(current.dirs, directory)
				dirs = append(dirs, new)
				id++
			} else {
				file := getFile(line)
				current.files = append(current.files, file)
			}
			dirs = remove(current.id, dirs)
			dirs = append(dirs, current)
		}
	}
	return dirs
}

func print(dirs []dir) {
	for _, dir := range dirs {
		fmt.Printf("%v\n", dir)
	}
	fmt.Printf("\n\n\n")
}

func sum(arr []int, max int) int {
	s := 0
	for _, a := range arr {
		if a <= max {
			s += a
		}
	}
	return s
}

func min(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func remove(id int, dirs []dir) []dir {
	ii := -1
	for i, dir := range dirs {
		if id == dir.id {
			ii = i
		}
	}
	if ii > -1 {
		return append(dirs[:ii], dirs[ii+1:]...)
	}
	return dirs
}

func find(name string, dirs []dir) dir {
	for _, dir := range dirs {
		if name == dir.name {
			return dir
		}
	}
	panic("failed to locate dir")
}

func findChild(name string, parent int, dirs []dir) dir {
	for _, dir := range dirs {
		if dir.name == name && dir.parent == parent {
			return dir
		}
	}
	panic("failed to locate child dir: " + name + " with parent: " + strconv.Itoa(parent))
}

func findParent(current dir, list []dir) dir {
	for _, dir := range list {
		if dir.id == current.parent {
			return dir
		}
	}
	panic("failed to locate parent dir")
}

func cd(line string) string {
	if regexp.MustCompile(`\$\scd\s.+`).MatchString(line) {
		return strings.Split(line, " ")[2]
	}
	return ""
}

func directory(line string) string {
	if regexp.MustCompile(`dir\s.+`).MatchString(line) {
		return strings.Split(line, " ")[1]
	}
	return ""
}

func ls(line string) bool {
	return regexp.MustCompile(`\$\sls`).MatchString(line)
}

func command(line string) bool {
	return regexp.MustCompile(`\$`).MatchString(line)
}

func getFile(line string) file {
	f := strings.Split(line, " ")
	size, _ := strconv.Atoi(f[0])
	return file{f[1], size}
}

func isCommand(line string) bool {
	return cd(line) != "" || ls(line)
}

func readFile(path string) string {
	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	text := string(bytes)
	return text
}
