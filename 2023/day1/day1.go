package main

import (
    "strconv"
    "errors"
    "bufio"
    "fmt"
    "os"
)

func main() {
    input, _ := readLines("input.txt")
    fmt.Println("Part 1:", part1(input))
    fmt.Println("Part 2:", part2(input))
}

func part1(input []string) int {
    var data []int
    for i := range input {
        numbers := ""
        for j := 0; j < len(input[i]); j++ {
            _, err := strconv.Atoi(string(input[i][j]))
            if err == nil {
                numbers += string(input[i][j])
            }
        }
        if len(numbers) == 1 {
            a, _ := strconv.Atoi(string(numbers[0]))
            data = append(data, a*11)
        } else {
            b, _ := strconv.Atoi(string(numbers[0]) + string(numbers[len(numbers)-1]))
            data = append(data, b)
        }
    }
    return sum(data)
}

func part2(input []string) int {
    var data []int
    for i := range input {
        numbers := ""
        index, prev := -1, ""

        substrings := getSubstrings(input[i])
        for j := range substrings {
            val, err := matchInt(substrings[j])
            if err == nil {
                numbers += val
                index = j
                prev = val
                break
            }
        } 

        for j := len(substrings)-1; j >= 0; j-- {
            val, err := matchInt(substrings[j])
            if err == nil {
                if j != index {
                    numbers += string(val)
                    break
                } else {
                    numbers += prev
                    break
                }
            }
        }
        val, _ := strconv.Atoi(numbers)
        data = append(data, val)
    }
    return sum(data)
}

func matchInt(s string) (string, error) {
    if s == "1" || s == "one" {
        return "1", nil
    } else if s == "2" || s == "two" {
        return "2", nil
    } else if s == "3" || s == "three" {
        return "3", nil
    } else if s == "4" || s == "four" {
        return "4", nil
    } else if s == "5" || s == "five" {
        return "5", nil
    } else if s == "6" || s == "six" {
        return "6", nil
    } else if s == "7" || s == "seven" {
        return "7", nil
    } else if s == "8" || s == "eight" {
        return "8", nil
    } else if s == "9" || s == "nine" {
        return "9", nil
    }
    return "", errors.New("bad input")
}

func getSubstrings(input string) []string {
    var result []string
    for i := 0; i < len(input); i++ {
        for j := i + 1; j <= len(input); j++ {
            result = append(result, input[i:j])
        }
    }
    return result
}

func sum(slice []int) int {
    result := 0
    for i := range slice {
        result += slice[i]
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
