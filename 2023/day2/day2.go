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
    re := regexp.MustCompile(`\d{1,3}\ \w{3,5}`)
    maxRed, maxGreen, maxBlue := 12, 13, 14
    part1, part2 := 0, 0

    for i := range input {
        matches := re.FindAllStringSubmatch(input[i], -1)
        pass := true
        red, green, blue := 0, 0, 0
        if len(matches) > 1 {
            for j := range matches {
                if matches[j][0][len(matches[j][0])-3:] == "red" {
                    val, _ := strconv.Atoi(matches[j][0][:len(matches[j][0])-4])
                    if val > red {
                        red = val
                    }
                    if val > maxRed {
                        pass = false
                    }
                } else if matches[j][0][len(matches[j][0])-5:] == "green" {
                    val, _ := strconv.Atoi(matches[j][0][:len(matches[j][0])-6])
                    if val > green {
                        green = val
                    }
                    if val > maxGreen {
                        pass = false
                    }
                } else if matches[j][0][len(matches[j][0])-4:] == "blue" {
                    val, _ := strconv.Atoi(matches[j][0][:len(matches[j][0])-5])
                    if val > blue {
                        blue = val
                    }
                    if val > maxBlue {
                        pass = false
                    }
                }
            }
            if pass {
                part1 += i+1
            }
            part2 += red*green*blue
        }
    }

    fmt.Println("Part 1:", part1)
    fmt.Println("Part 2:", part2)
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
