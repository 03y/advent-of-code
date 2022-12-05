package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stacks := [][]string {
		{"H", "C", "R"},
		{"B", "J", "H", "L", "S", "F"},
		{"R", "M", "D", "H", "J", "T", "Q"},
		{"S", "G", "R", "H", "Z", "B", "J"},
		{"R", "P", "F", "Z", "T", "D", "C", "B"},
		{"T", "H", "C", "G"},
		{"S", "N", "V", "Z", "B", "P", "W", "L"},
		{"R", "J", "Q", "G", "C"},
		{"L", "D", "T", "R", "H", "P", "F", "S"},
	}

	stacks2 := [][]string {
		{"H", "C", "R"},
		{"B", "J", "H", "L", "S", "F"},
		{"R", "M", "D", "H", "J", "T", "Q"},
		{"S", "G", "R", "H", "Z", "B", "J"},
		{"R", "P", "F", "Z", "T", "D", "C", "B"},
		{"T", "H", "C", "G"},
		{"S", "N", "V", "Z", "B", "P", "W", "L"},
		{"R", "J", "Q", "G", "C"},
		{"L", "D", "T", "R", "H", "P", "F", "S"},
	}

	// stacks := [][]string{
	// 	{"Z", "N"},
	// 	{"M", "C", "D"},
	// 	{"P"},
	// }

	// stacks2 := [][]string{
	// 	{"Z", "N"},
	// 	{"M", "C", "D"},
	// 	{"P"},
	// }

	// queue := []string{}

	input, _ := readLines("test_data.txt")
	var sectionMarker, stackCount = 0, 0
	for i := range input {
		if input[i] == "" {
			sectionMarker = i
			break
		} else if (input[i][1] == '1') {
			stackCount = parseInt(strings.Fields(input[i])[len(strings.Fields(input[i]))-1])
		}
	}

	yValues := []int{}
	for i := 1; i <= stackCount; i++ {
		yValues = append(yValues, charLocation(input[sectionMarker - 1], strconv.Itoa(i)))
	}

	data := make([][]string, len(input) - sectionMarker - 1)
	for i := 0; i < sectionMarker - 1; i++ {
		for j := range yValues {
			if input[i][yValues[j]] != ' ' {
				data[j] = append(data[j], string(input[i][yValues[j]]))
			}
		}
	}

	for i := sectionMarker; i < len(input); i++ {
		if input[i] != "" {
			var line string = strings.Replace(input[i], "move ", "", -1)
			line = strings.Replace(line, " from ", " ", -1)
			line = strings.Replace(line, " to ", " ", -1)

			var n int = parseInt(strings.Fields(line)[0])
			var from int = parseInt(strings.Fields(line)[1]) - 1
			var to int = parseInt(strings.Fields(line)[2]) - 1

			crates := ""
			for j := 0; j < n; j++ {
				crates = stacks[from][len(stacks[from]) - 1] + crates
				stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
				stacks[from] = stacks[from][:len(stacks[from])-1]
			}
			
			for j := 0; j < len(crates); j++ {
				stacks2[to] = append(stacks2[to], string(crates[j]))
				stacks2[from] = stacks2[from][:len(stacks2[from])-1]
			}
		}
	}

	fmt.Print("Part 1 Answer: ")
	for i := range stacks {
		fmt.Print(stacks[i][len(stacks[i])-1])
	}
	fmt.Print("\nPart 2 Answer: ")
	for i := range stacks2 {
		fmt.Print(stacks2[i][len(stacks2[i])-1])
	}
	fmt.Println()
}

func charLocation(s string, c string) int {
	for i := range s {
		if string(s[i]) == c {
			return i
		}
	}
	return -1
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func readLines(path string) ([]string, error) {
    file, _ := os.Open(path)
    defer file.Close()
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}
