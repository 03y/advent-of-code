package main

import (
    "strconv"
    "bufio"
    "fmt"
    "os"
)

func main() {
    input, _ := readLines("test_input.txt")
    input = append(input, "")
    adapters := make([]int, 0)
    device_joltage := 0
    order := make([]int, 0)
    order = append(order, 0)

    for i := range input {
        if input[i] != "" {
            n, _ := strconv.Atoi(input[i])
            adapters = append(adapters, n)
        }
    }

    device_joltage = max(adapters)+3

    for i := range adapters {
        if adapters[i] - order[len(order)-1] <= 3 && adapters[i] - order[len(order)-1] > 0 {
            order = append(order, adapters[i])
        }
    }
}

func min(list []int) int {
    result := list[0]
    for i := range list {
        if list[i] < result {
            result = list[i]
        }
    }
    return result
}

func max(list []int) int {
    result := list[0]
    for i := range list {
        if list[i] > result {
            result = list[i]
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
