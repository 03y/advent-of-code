package main

import (
    "strings"
    "bufio"
    "fmt"
    "os"
)

type scratchcard struct {
    id        int
    children  []int
}

func main() {
    var tree   []scratchcard
    var part1  []int
    input, _ := readLines("input.txt")
    cards    := make(map[int]int)
    part2    := 0

    for i := range input {
        var numbers []string
        var winning []string
        card, read := strings.Fields(input[i]), false
        won        := 0
        part1       = append(part1, 0)
        tree        = append(tree, scratchcard{i, []int{}})

        for j := 2; j < len(card); j++ {
            if card[j] == "|" {
                read = true
                continue
            }

            if !read {
                numbers = append(numbers, card[j])
            } else {
                winning = append(winning, card[j])
            }
        }

        for _, n := range numbers {
            for _, w := range winning {
                if n == w {
                    if part1[i] == 0 {
                        part1[i]++
                    } else {
                        part1[i] *= 2
                    }
                    won++
                    tree[i].children = append(tree[i].children, i+won)
                }
            }
        }
        cards[i] = won
    }

    for i := range tree {
        traverse(tree[i], tree, &part2)
    }

    fmt.Println("Part 1:", sum(part1))
    fmt.Println("Part 1:", part2)
}

func traverse(card scratchcard, tree []scratchcard, cardCount *int) {
    *cardCount = *cardCount+1
    if len(card.children) == 0 {
        return
    }

    for _, child := range card.children {
        traverse(tree[child], tree, cardCount)
    }
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
