package main

import (
    "strconv"
    "bufio"
    "fmt"
    "os"
)

func main() {
    input, _ := readLines("input.txt")
    pairs := parsePairs(input)
    correctOrder := 0
    for i, pair := range pairs {
        outcome := lessThan(pair[0], pair[1])
        if outcome >= 0 {
            correctOrder += i+1
        }
    }
    fmt.Println("Part 1 Answer:", correctOrder)
}

func largestItem(arr []int) int {
    max := 0
    for _, n := range arr {
        if n > max {
            max = n
        }
    }
    return max
}

func lessThan(a, b any) int {
    // returns -1 is a smaller than b, 1 if b is smaller than a, 0 if they are equal
    aInt, aValidInt := a.(int)
    bInt, bValidInt := b.(int)
    if aValidInt && bValidInt {
        if aInt > bInt {
            return -1
        } else if bInt > aInt {
            return 1
        }
        return 0
    }
    aArr, aValidArr := a.([]any)
    bArr, bValidArr := b.([]any)
    if !aValidArr {
        aArr = []any{aInt}
    }
    if !bValidArr {
        bArr = []any{bInt}
    }
    max := largestItem([]int{len(aArr), len(bArr)})
    for i := 0; i < max; i++ {
        if i >= len(aArr) {
            return 1
        }
        if i >= len(bArr) {
            return -1
        }
        if sub := lessThan  (aArr[i], bArr[i]); sub != 0 {
            return sub
        }
    }
    return 0
}

func parsePairs(input []string) [][]any {
    pairs := make([][]any, 0)
    pair := make([]any, 0, 2)
    for i := range input {
        if input[i] == "" {
            continue
        }

        expression, _ := parseExpression(input[i])
        pair = append(pair, expression)
        if len(pair) == 2 {
            pairs = append(pairs, pair)
            pair = make([]any, 0, 2)
        }
    }
    return pairs
}

func parseExpression(expression string) (any, int) {
    chars := []rune(expression)
    index := 0
    out := make([]any, 0)
    nChars := make([]rune, 0)
    for index < len(chars) {
        char := chars[index]
        switch char {
        case '[':
            x, i := parseExpression(string(chars[index+1:]))
            out = append(out, x)
            index += i + 1
        case ']':
            if len(nChars) > 0 {
                n, _ := strconv.Atoi(string(nChars))
                out = append(out, n)
                // nChars = make([]rune, 0)
            }
            index++
            return out, index
        case ',':
            if len(nChars) > 0 {
                n, _ := strconv.Atoi(string(nChars))
                out = append(out, n)
                nChars = make([]rune, 0)
            }
            index++
        default:
            nChars = append(nChars, char)
            index++
        }
    }
    return out, index
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
