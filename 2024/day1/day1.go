package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := readLines("input.txt")

	part1, part2 := 0, 0
	var left []int
	var right []int

	for i := range input {
		c := strings.Split(input[i], " ")
		n, _ := strconv.Atoi(c[0])
		left = append(left, n)
		n, _ = strconv.Atoi(c[len(c)-1])
		right = append(right, n)
	}

	for i := range len(left) {
		part2 += left[i] * occurrences(right, left[i])
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := range len(left) {
		part1 += max(left[i], right[i]) - min(left[i], right[i])
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func occurrences(list []int, target int) int {
	result := 0
	for i := range list {
		if list[i] == target {
			result++
		}
	}
	return result
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
