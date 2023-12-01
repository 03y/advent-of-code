package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    input, _ := readLines("test_input.txt")
    input = append(input, "")
    answers, part1_sum, part2_sum := make(map[string]bool), 0, 0
    everyone_agrees, first := make(map[string]bool), true

    for i := range input {
        if input[i] != "" {
            for j := range input[i] {
                answers[string(input[i][j])] = true
            }
            fmt.Println("person", i, "answered", answers)
            if first {
                first = false
                for k, v := range answers {
                    everyone_agrees[k] = v
                }
            } else {
                everyone_agrees = intersection(everyone_agrees, answers)
            }
        } else {
            part1_sum += len(answers)
            part2_sum += len(everyone_agrees)
            fmt.Println("everyone agreed on", everyone_agrees, "\n")

            first = true
            for k := range answers {
                delete(answers, k)
            }
            for k := range everyone_agrees {
                delete(everyone_agrees, k)
            }
            fmt.Println("new group")
        }
    }

    fmt.Println("Part 1:", part1_sum)
    fmt.Println("Part 2:", part2_sum)
}

func intersection(a, b map[string]bool) map[string]bool {
    result := make(map[string]bool)
    for k := range a {
        if b[k] {
            result[k] = true
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
