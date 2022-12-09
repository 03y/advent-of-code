package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := readLines("test_data.txt")
	HPos := [2]int{0, 0}
	TPos := [2]int{0, 0}
	// knots := [][2]int{}
	visitedCoords := [][2]int{}

	for i := 0; i < len(input); i++ {
		direction := strings.Split(input[i], " ")[0]
		distance, _ := strconv.Atoi(strings.Split(input[i], " ")[1])

		// fmt.Println("HPos:", HPos, "\tTPos:", TPos)

		switch direction {
		case "R":
			for j := 0; j < distance; j++ {
				HPos[1]++
				if HPos[0] > TPos[0] && HPos[1] - TPos[1] > 1 {
					// case for when the Head is above the Tail and the Tail is too far to the right
					TPos[0]++
					TPos[1]++
				} else if HPos[0] < TPos[0] && HPos[1] - TPos[1] > 1 {
					// case for when the Head is below the Tail and the Tail is too far to the right
					TPos[0]--
					TPos[1]++
				} else if HPos[1] - TPos[1] > 1 {
					// case for when the Tail is too far to the right
					TPos[1]++
				}
				if !contains(visitedCoords, TPos) {
					visitedCoords = append(visitedCoords, TPos)
				}
			}
		case "L":
			for j := 0; j < distance; j++ {
				HPos[1]--
				if HPos[0] > TPos[0] && HPos[1] - TPos[1] < -1 {
					// case for when the Head is above the Tail and the Tail is too far to the left
					TPos[0]++
					TPos[1]--
				} else if HPos[0] < TPos[0] && HPos[1] - TPos[1] < -1 {
					// case for when the Head is below the Tail and the Tail is too far to the left
					TPos[0]--
					TPos[1]--
				} else if HPos[1] - TPos[1] < -1 {
					// case for when the Tail is too far to the left
					TPos[1]--
				}
				if !contains(visitedCoords, TPos) {
					visitedCoords = append(visitedCoords, TPos)
				}
			}
		case "U":
			for j := 0; j < distance; j++ {
				HPos[0]++
				if HPos[1] > TPos[1] && HPos[0] - TPos[0] > 1 {
					// case for when the Head is above the Tail and the Tail is too far to the left
					TPos[0]++
					TPos[1]++
				} else if HPos[1] < TPos[1] && HPos[0] - TPos[0] > 1 {
					// case for when the Head is below the Tail and the Tail is too far to the left
					TPos[0]++
					TPos[1]--
				} else if HPos[0] - TPos[0] > 1 {
					// case for when the Tail is too far to the left
					TPos[0]++
				}
				if !contains(visitedCoords, TPos) {
					visitedCoords = append(visitedCoords, TPos)
				}
			}
		case "D":
			for j := 0; j < distance; j++ {
				HPos[0]--
				if HPos[1] > TPos[1] && HPos[0] - TPos[0] < -1 {
					// case for when the Head is above the Tail and the Tail is too far to the right
					TPos[0]--
					TPos[1]++
				} else if HPos[1] < TPos[1] && HPos[0] - TPos[0] < -1 {
					// case for when the Head is below the Tail and the Tail is too far to the right
					TPos[0]--
					TPos[1]--
				} else if HPos[0] - TPos[0] < -1 {
					// case for when the Tail is too far to the right
					TPos[0]--
				}
				if !contains(visitedCoords, TPos) {
					visitedCoords = append(visitedCoords, TPos)
				}
			}
		}
	}
	fmt.Println()
	printHeatmap(visitedCoords)
	fmt.Println()

	fmt.Println("Part 1 Answer:", len(visitedCoords))
}

func printHeatmap(visitedCoords [][2]int) {
	maxX, maxY, minX, minY := 0, 0, 0, 0
	for _, coord := range visitedCoords {
		if coord[0] > maxX {
			maxX = coord[0]
		}
		if coord[0] < minX {
			minX = coord[0]
		}
		if coord[1] > maxY {
			maxY = coord[1]
		}
		if coord[1] < minY {
			minY = coord[1]
		}
	}

	for i := maxX; i >= minX; i-- {
		for j := minY; j <= maxY; j++ {
			if i == 0 && j == 0 {
				fmt.Print("S")
			} else if contains(visitedCoords, [2]int{i, j}) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func contains(s [][2]int, e [2]int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
