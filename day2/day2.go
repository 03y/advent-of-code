package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := readLines("input.txt")
	part1_score := 0
	part2_score := 0

	for _, line := range input {
		part1_score += outcome(translate(strings.Split(line, " ")[0]), translate(strings.Split(line, " ")[1]))
		part2_score += outcome(translate(strings.Split(line, " ")[0]), part2_calculation(translate(strings.Split(line, " ")[0]), strings.Split(line, " ")[1]))
	}
	fmt.Println("Part 1 Answer: ", part1_score, "\nPart 2 Answer: ", part2_score)
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

func translate(input string) string {
	switch input {
		case "X":
			return "rock"
		case "Y":
			return "paper"
		case "Z":
			return "scissors"
		case "A":
			return "rock"
		case "B":
			return "paper"
		case "C":
			return "scissors"
	}
	return ""
}

func outcome(opponent_move string, my_move string) int {
	score := 0
	rps := map[string]string {
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}

	switch my_move {
		case "rock":
			score += 1
		case "paper":
			score += 2
		case "scissors":
			score += 3
	}

	if rps[my_move] == opponent_move {
		score += 6
	} else if rps[opponent_move] == my_move {
		score += 0
	} else {
		score += 3
	}

	return score
}

func part2_calculation(opponent_move string, outcome string) string {
	rps := map[string]string {
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}

	switch outcome {
		case "X":
			return rps[opponent_move]
		case "Y":
			return opponent_move
		case "Z":
			return rps[rps[opponent_move]]
	}
	return ""
}
