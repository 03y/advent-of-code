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
    var sectionMarker, stackCount = 0, 0
    for i := range input {
        if input[i] == "" {
            sectionMarker = i
            break
        } else if (input[i][1] == '1') {
            stackCount = parseInt(strings.Fields(input[i])[len(strings.Fields(input[i]))-1])
        }
    }

    yValues := []int{}
    for i := 1; i <= stackCount; i++ {
        yValues = append(yValues, charLocation(input[sectionMarker - 1], strconv.Itoa(i)))
    }

    data := make([][]string, len(input) - sectionMarker - 1)
    for i := 0; i < sectionMarker - 1; i++ {
        for j := range yValues {
            if input[i][yValues[j]] != ' ' {
                data[j] = append([]string{string(input[i][yValues[j]])}, data[j]...)
            }
        }
    }
    
    data2 := make([][]string, len(data))
    for i := range data {
        data2[i] = make([]string, len(data[i]))
        copy(data2[i], data[i])
    }

    for i := sectionMarker; i < len(input); i++ {
        if input[i] != "" {
            var line string = strings.Replace(input[i], "move ", "", -1)
            line = strings.Replace(line, " from ", " ", -1)
            line = strings.Replace(line, " to ", " ", -1)

            var n int = parseInt(strings.Fields(line)[0])
            var from int = parseInt(strings.Fields(line)[1]) - 1
            var to int = parseInt(strings.Fields(line)[2]) - 1

            s := ""
            for j := 0; j < n; j++ {
                data[to] = append(data[to], data[from][len(data[from])-1])
                data[from] = data[from][:len(data[from])-1]

                s += data2[from][len(data2[from])-1]
                data2[from] = data2[from][:len(data2[from])-1]
            }
            
            for j := 0; j < n; j++ {
                data2[to] = append(data2[to], string(s[len(s)-1]))
                s = s[:len(s)-1]
            }
        }
    }

    fmt.Print("Part 1 Answer: ")
    for i := range data {
        if len(data[i]) > 0 {
            fmt.Print(data[i][len(data[i])-1])
        }
    }
    fmt.Print("\nPart 2 Answer: ")
    for i := range data2 {
        if len(data2[i]) > 0 {
            fmt.Print(data2[i][len(data2[i])-1])
        }
    }
    fmt.Println()
}


func charLocation(s string, c string) int {
    for i := range s {
        if string(s[i]) == c {
            return i
        }
    }
    return -1
}

func parseInt(s string) int {
    n, _ := strconv.Atoi(s)
    return n
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
