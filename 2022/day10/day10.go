package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    input, _ := readLines("input.txt")
    var display [240]string
    spritePos := 0
    for i := 0; i < 240; i++ {
        display[i] = "."
    }

    cycle, register := 1, 1
    signalStrength := map[int]int{}

    for i := 0; i < len(input); i++ {
        if cycle % 4 == 0 {
            spritePos += 3
        }

        if cycle >= spritePos && cycle <= spritePos+2 {
            display[i] = "#"
        }
        
        instruction := strings.Split(input[i], " ")

        if instruction[0] == "addx" {
            val, _ := strconv.Atoi(instruction[1])
            // first cycle of addx operation
            cycle++
            if cycle == 20 || (cycle-20) % 40 == 0 {
                signalStrength[cycle] = register
            }
            // second cycle of addx operation
            cycle++
            register += val
            if cycle == 20 || (cycle-20) % 40 == 0 {
                signalStrength[cycle] = register
            }
        } else {
            cycle++
            if cycle == 20 || (cycle-20) % 40 == 0 {
                signalStrength[cycle] = register
            }
        }
    }

    total := 0
    for cycle, value := range signalStrength {
        total = total + (cycle*value)
    }

    fmt.Println("Part 1 Answer:", total)
    fmt.Println("Part 2 Answer:")

    for i := 0; i < 240; i++ {
        fmt.Print(display[i])
        if i % 40 == 39 {
            fmt.Println()
        }
    }
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
