package main

import (
    "strconv"
    "bufio"
    "fmt"
    "os"
)

func main() {
    input, _ := readLines("input.txt")
    input = append(input, "")
    list := make([]int, 0)
    limit := 25

    for i := range input {
        if input[i] != "" {
            if i < limit {
                continue
            } else {
                n, _ := strconv.Atoi(input[i])
                
                for j := i-1; j > i-1-limit; j-- {
                    value, _ := strconv.Atoi(input[j])
                    list = append(list, value)
                }
                
                if !part1(list, n) {
                    fmt.Println("Part 1:", input[i])

                    superset := make([]int, 0)
                    for j := range input {
                        n, _ := strconv.Atoi(input[j])
                        superset = append(superset, n)
                    }

                    temp := part2(superset, n)
                    if temp != nil {
                        fmt.Println("Part 2:", min(temp)+max(temp))
                    }
                }
                list = make([]int, 0)
            }
        }
    }
}

func part1(list []int, n int) bool {
    for i := range list {
        for j := range list {
            if list[i]+list[j] == n {
                return true
            }
        }
    }
    return false
}

func part2(list []int, n int) []int {
    for i := 0; i < len(list)-1; i++ {
        // for each position in list

        for j := 2; j < len(list)-1; j ++ {
            // for each possible length
            // which is length of list-1
            // min size 2

            // generate subset at position i, of length j
            subset := get_subset(list, i, j)

            // check if sum of subset is n
            if sum(subset) == n {
                return subset
            }
        }
    }
    return nil
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

func get_subset(list []int, start int, length int) []int {
    result := make([]int, 0)
    for i := start; i < length; i++ {
        result = append(result, list[i])
    }
    return result
}

func sum(list []int) int {
    res := 0
    for i := range list {
        res += list[i]
    }
    return res
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
