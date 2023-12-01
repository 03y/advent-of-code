package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
    input, _ := readLines("input.txt")
    eclipse_count := 0
    overlap := 0

    for i := range input {
        words := strings.Split(input[i], ",")

        range1_lower, _ := strconv.Atoi(strings.Split(words[0], "-")[0])
        range1_upper, _ := strconv.Atoi(strings.Split(words[0], "-")[1])

        range2_lower, _ := strconv.Atoi(strings.Split(words[1], "-")[0])
        range2_upper, _ := strconv.Atoi(strings.Split(words[1], "-")[1])

        if (range1_lower <= range2_lower && range1_upper >= range2_upper) || (range2_lower <= range1_lower && range2_upper >= range1_upper) {
            eclipse_count++
        }

        if (range1_lower <= range2_lower && range1_upper >= range2_lower) || (range2_lower <= range1_lower && range2_upper >= range1_lower) {
            overlap++
        }
    }
    fmt.Println("Part 1 Answer:", eclipse_count)
    fmt.Println("Part 2 Answer:", overlap)
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

