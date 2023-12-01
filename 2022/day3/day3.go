package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    part1()
    part2()
}

func part1() {
    input, _ := readLines("input.txt")
    var total int
    
    for i := range input {
        var compartment1 string

        for j := 0; j < len(input[i]) / 2; j++ {
            compartment1 += string(input[i][j])
        }

        for j := len(input[i]) / 2; j < len(input[i]); j++ {
            if stringContains(compartment1, string(input[i][j])) {
                total += charValue(string(input[i][j]))
                break
            }
        }
    }

    fmt.Println("Part 1 Answer:", total)
}

func part2() {
    input, _ := readLines("input.txt")
    var total int

    for i := 0; i < len(input); i += 3 {
        var pack1 string = input[i]
        var pack2 string = input[i+1]
        var pack3 string = input[i+2]

        for j := 0; j < len(pack1); j++ {
            if stringContains(pack2, string(pack1[j])) && stringContains(pack3, string(pack1[j])) {
                total += charValue(string(pack1[j]))
                break
            }
        }
    }

    fmt.Println("Part 2 Answer:", total)
}

func charValue(c string) int {
    if c >= "a" && c <= "z" {
        return int(c[0]) - 96
    } else {
        return int(c[0]) - 38
    }
}

func stringContains(s string, e string) bool {
    for _, a := range s {
        if string(a) == e {
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
