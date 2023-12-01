package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := readLines("input.txt")
	matrix := make([][]string, len(input))
	part1, part2_collisions, part2 := 0, []int{0, 0, 0, 0, 0}, 0

	for i := range input {
		matrix[i] = make([]string, len(input[0]))
		if input[i] != "" {
			for j := range input[i] {
				matrix[i][j] = string(input[i][j])
			}
		}
	}

	move := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	for i := range move {
		row, col := 0, 0
		for row < len(matrix) {
			if matrix[row][col] == "#" {
				part2_collisions[i]++
				if move[i][0] == 3 && move[i][1] == 1 {
					part1++
				}
			}
			row += move[i][1]
			col += move[i][0]

			if col >= len(matrix[0]) {
				col = col - len(matrix[0])
			}
			if row >= len(matrix) {
				break
			}
		}
	}

	part2 = part2_collisions[0] * part2_collisions[1] * part2_collisions[2] * part2_collisions[3] * part2_collisions[4]

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
