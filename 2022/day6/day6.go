package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	data, _ := readLines("input.txt")
	queues := make([][]string, 2)
	answers := make([]int, 2)

	for i := 0; i < len(data[0]); i++ {
		if i < 4 {
			queues[0] = append(queues[0], string(data[0][i]))
		}
		if i < 14 {
			queues[1] = append(queues[1], string(data[0][i]))
		}
		
		if i >= 4 {
			queues[0] = append(queues[0], string(data[0][i])) // push
			queues[0] = queues[0][1:] // pop

			if uniqueChars(sliceToStr(queues[0])) && answers[0] == 0 {
				answers[0] = i + 1
			}
		}
		
		if i >= 14 {
			queues[1] = append(queues[1], string(data[0][i])) // push
			queues[1] = queues[1][1:] // pop

			if uniqueChars(sliceToStr(queues[1])) && answers[1] == 0 {
				answers[1] = i + 1
			}
		}
	}

	fmt.Println("Part 1 answer: ", answers[0])
	fmt.Println("Part 2 answer: ", answers[1])
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

func sliceToStr(q []string) string {
	var s string
	for _, v := range q {
		s += v
	}
	return s
}

func uniqueChars(s string) bool {
	m := make(map[rune]bool)
	for _, c := range s {
		if m[c] {
			return false
		}
		m[c] = true
	}
	return true
}
