package main

import (
    "strconv"
    "regexp"
    "bufio"
    "fmt"
    "os"
)

func main() {
    input, _ := readLines("input.txt")
    input = append(input, "")
    re := regexp.MustCompile(`^([a-z]{3})\s([+-]\d+)$`)
    seen := make([]int, 0)
    acc := 0

    for i := 0; i < len(input); {
        if input[i] != "" {
            if i > 0 && contains(seen, i) {
                fmt.Println("Part 1:", acc)
                return
            }
            seen = append(seen, i)
            match := re.FindAllStringSubmatch(input[i], -1)
            if len(match) > 0 {
                n, _ := strconv.Atoi(match[0][2])
                switch (match[0][1]) {
                case "nop":
                    i++
                case "acc":
                    acc += n
                    i++
                case "jmp":
                    i += n
                }
            }
        }
    }
}

func contains(slice []int, n int) bool {
    for i := range slice {
        if slice[i] == n {
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
