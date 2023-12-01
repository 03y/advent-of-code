package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := readLines("input.txt")

	for i := range input {
		if input[i] != "" {
			value, _ := strconv.Atoi(input[i])
			for j := range input {
				if input[j] != "" {
					value2, _ := strconv.Atoi(input[j])
					if value+value2 == 2020 {
						fmt.Println("Part 1:", value, "*", value2, "=", value*value2)
					}
					for k := range input {
						if input[k] != "" {
							value3, _ := strconv.Atoi(input[k])
							if value+value2+value3 == 2020 {
								fmt.Println("Part 2:", value, "*", value2, "*", value3, "=", value*value2*value3)
								return
							}
						}
					}
				}
			}
		}
	}
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
