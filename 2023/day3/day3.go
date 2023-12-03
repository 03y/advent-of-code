package main

import (
    "strconv"
    "bufio"
    "fmt"
    "os"
)

type part struct {
    value int
    line  int
    start int
    end   int
}

type gear struct {
    line  int
    pos   int
}

func main() {
    input, _ := readLines("input.txt")
    part1, part2 := 0, 0
    component, start, end := "", -1, -1
    component1, component2 := -1, -1
    var parts []part
    var gears []gear

    buffer := ""
    for i := 0; i < len(input); i++ {
        buffer += "."
    }
    input = append(input, buffer)
    for i := range input {
        for j := 0; j < len(input[i]); j++ {
            if string(input[i][j]) == "*" {
                gears = append(gears, gear{i, j})
            }
            if isDigit(string(input[i][j])) {
                if component == "" {
                    start = j
                }
                component += string(input[i][j])
            }
            if j == len(input[i])-1 || !isDigit(string(input[i][j+1])) {
                if component != "" && start != -1 {
                    if end == -1 {
                        end = j
                    }
                    pass := false
                    
                    // east
                    if end < len(input[i])-1 {
                        if isSymbol(string(input[i][end+1])) {
                            // todo: optimisation: if we pass 1 check we dont need to check the rest, make func done
                            pass = true
                        }
                    }

                    // west
                    if start > 0 {
                        if isSymbol(string(input[i][start-1])) {
                            pass = true
                        }
                    }

                    // north
                    if i > 0 {
                        for k := start; k <= end; k++ {
                            if isSymbol(string(input[i-1][k])) {
                                pass = true
                                break
                            }
                        }
                    }

                    // north west
                    if i > 0 && start > 0 {
                        if isSymbol(string(input[i-1][start-1])) {
                            pass = true
                        }
                    }

                    // north east
                    if i > 0 && end < len(input[i])-1 {
                        if isSymbol(string(input[i-1][end+1])) {
                            pass = true
                        }
                    }

                    // south 
                    if i < len(input)-1 {
                        for k := start; k <= end; k++ {
                            if isSymbol(string(input[i+1][k])) {
                                pass = true
                                break
                            }
                        }
                    }

                    // south west
                    if i < len(input)-1 && start > 0 {
                        if isSymbol(string(input[i+1][start-1])) {
                            pass = true
                        }
                    }

                    // south east
                    if i < len(input)-1 && end < len(input[i])-1 {
                        if isSymbol(string(input[i+1][end+1])) {
                            pass = true
                        }
                    }

                    n, _ := strconv.Atoi(component)
                    if pass {
                        part1 += n
                    }
                    parts = append(parts, part{n, i, start, end})
                    component, start, end = "", -1, -1
                }
            }
        }
    }

    // todo: optimise and run in one pass
    for _, gear := range gears {
        touchCount := 0
        for _, part := range parts {
            if (gear.line == part.line || gear.line - 1 == part.line || gear.line + 1 == part.line) && 
                gear.pos >= part.start - 1 && gear.pos <= part.end + 1 {
                touchCount++
                if touchCount == 1 {
                    component2 = part.value
                } else {
                    component1 = part.value
                }

                if touchCount == 2 {
                    part2 += component1*component2
                }
            }
        }
    }

    fmt.Println("Part 1:", part1)
    fmt.Println("Part 2:", part2)
}

func isDigit(s string) bool {
    switch s {
        case "0":
            return true
        case "1":
            return true
        case "2":
            return true
        case "3":
            return true
        case "4":
            return true
        case "5":
            return true
        case "6":
            return true
        case "7":
            return true
        case "8":
            return true
        case "9":
            return true
        default:
            return false
    }
}

func isSymbol(s string) bool {
    illegal := []string{".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
    for _, v := range illegal {
        if s == v {
            return false
        }
    }
    return true
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
