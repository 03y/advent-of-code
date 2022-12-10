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
	cycle, register := 1, 1
	signalStrength := map[int]int{}

	for i := 0; i < len(input); i++ {
		instruction := strings.Split(input[i], " ")

		if instruction[0] == "addx" {
			val, _ := strconv.Atoi(instruction[1])
			// first cycle of addx operation
			cycle++
			if cycle == 20 || (cycle-20) % 40 == 0 {
				signalStrength[cycle] = register
			}
			// second cycle of addx operation
			cycle++
			register += val
			if cycle == 20 || (cycle-20) % 40 == 0 {
				signalStrength[cycle] = register
			}
		} else {
			cycle++
			if cycle == 20 || (cycle-20) % 40 == 0 {
				signalStrength[cycle] = register
			}
		}
	}

	total := 0
	for cycle, value := range signalStrength {
		total = total + (cycle*value)
	}

	fmt.Println("Part 1 Answer:", total)
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