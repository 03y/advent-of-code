package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := readLines("input.txt")
	part1, part2 := 0, 0

	for i := range input {
		var data []int

		for _, c := range strings.Split(input[i], " ") {
			n, _ := strconv.Atoi(c)
			data = append(data, n)
		}

		if valid(data) {
			part1++
			part2++
		} else {
			for j := range data {
				t := append([]int(nil), data[:j]...)
				t = append(t, data[j+1:]...)
				if valid(t) {
					part2++
					break
				}
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func valid(list []int) bool {
	if sort.IntsAreSorted(list) {
		return checkDeltas(list)
	} else {
		slices.Reverse(list)
		if sort.IntsAreSorted(list) {
			return checkDeltas(list)
		}
	}
	return false
}

func checkDeltas(list []int) bool {
	for i := range list {
		if i > 0 {
			if list[i]-list[i-1] < 1 || list[i]-list[i-1] > 3 {
				return false
			}
		}
	}
	return true
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
