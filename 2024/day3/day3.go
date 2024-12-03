package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := readLines("input.txt")
	part1, part2 := 0, 0
	input2 := ""

	for i := range input {
		part1 += calculate(input[i])
		input2 += input[i]
	}
	part2 += conditionals(input2)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func conditionals(s string) int {
	result := 0
	input := strings.Split(s, "do()")
	for i := range input {
		do := strings.Split(input[i], "don't()")[0]
		result += calculate(do)
	}
	return result
}

func calculate(s string) int {
	re := regexp.MustCompile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
	result := 0
	instructions := re.FindAllStringSubmatch(s, -1)
	a, b := 0, 0

	if len(instructions) > 0 {
		for j := range len(instructions) {
			a, _ = strconv.Atoi(instructions[j][1])
			b, _ = strconv.Atoi(instructions[j][2])
			result += a * b
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
