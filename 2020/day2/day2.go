package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := readLines("input.txt")
	part1, part2 := 0, 0

	for i := range input {
		if input[i] != "" {
			rangeStr := strings.Split(input[i][:strings.Index(input[i], " ")], "-")
			min, _ := strconv.Atoi(rangeStr[0])
			max, _ := strconv.Atoi(rangeStr[1])
			letter := input[i][strings.Index(input[i], " ")+1]
			password := input[i][strings.Index(input[i], ":")+2:]
			cnt, pass := 0, false

			for j := range password {
				if password[j] == letter {
					cnt++
					if j == min-1 || j == max-1 {
						if !pass {
							pass = true
						} else {
							pass = false
						}
					}
				}
			}
			if cnt >= min && cnt <= max {
				part1++
			}
			if pass {
				part2++
			}
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
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
