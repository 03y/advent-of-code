package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Knot struct {
    parent *Knot
    pos   [2]int
}

func main() {
    input, _ := readLines("input.txt")
    head := Knot{nil, [2]int{0, 0}}
    tail := Knot{&head, [2]int{0, 0}}
    
    ropeA := []*Knot{&head, &tail}
    shortTailVisitedCoords := [][2]int{}

    for i := 0; i < len(input); i++ {
        direction := strings.Split(input[i], " ")[0]
        distance, _ := strconv.Atoi(strings.Split(input[i], " ")[1])
        
        ropeA[1].parent.pos, ropeA[1].pos, shortTailVisitedCoords = moveKnot(direction, distance, ropeA[1].parent.pos, ropeA[1].pos, shortTailVisitedCoords)
    }

    // printHeatmap(shortTailVisitedCoords)
    fmt.Println("Part 1 Answer:", len(shortTailVisitedCoords))
}

func moveKnot(direction string, distance int, parentPos [2]int, pos [2]int, visitedCoords [][2]int) ([2]int, [2]int, [][2]int) {
    switch direction {
    case "R":
        for j := 0; j < distance; j++ {
            parentPos[1]++
            if parentPos[0] > pos[0] && parentPos[1] - pos[1] > 1 {
                // case for when the Head is above the Tail and the Tail is too far to the right
                pos[0]++
                pos[1]++
            } else if parentPos[0] < pos[0] && parentPos[1] - pos[1] > 1 {
                // case for when the Head is below the Tail and the Tail is too far to the right
                pos[0]--
                pos[1]++
            } else if parentPos[1] - pos[1] > 1 {
                // case for when the Tail is too far to the right
                pos[1]++
            }
            if !contains(visitedCoords, pos) {
                visitedCoords = append(visitedCoords, pos)
            }
        }
    case "L":
        for j := 0; j < distance; j++ {
            parentPos[1]--
            if parentPos[0] > pos[0] && parentPos[1] - pos[1] < -1 {
                // case for when the Head is above the Tail and the Tail is too far to the left
                pos[0]++
                pos[1]--
            } else if parentPos[0] < pos[0] && parentPos[1] - pos[1] < -1 {
                // case for when the Head is below the Tail and the Tail is too far to the left
                pos[0]--
                pos[1]--
            } else if parentPos[1] - pos[1] < -1 {
                // case for when the Tail is too far to the left
                pos[1]--
            }
            if !contains(visitedCoords, pos) {
                visitedCoords = append(visitedCoords, pos)
            }
        }
    case "U":
        for j := 0; j < distance; j++ {
            parentPos[0]++
            if parentPos[1] > pos[1] && parentPos[0] - pos[0] > 1 {
                // case for when the Head is above the Tail and the Tail is too far to the left
                pos[0]++
                pos[1]++
            } else if parentPos[1] < pos[1] && parentPos[0] - pos[0] > 1 {
                // case for when the Head is below the Tail and the Tail is too far to the left
                pos[0]++
                pos[1]--
            } else if parentPos[0] - pos[0] > 1 {
                // case for when the Tail is too far to the left
                pos[0]++
            }
            if !contains(visitedCoords, pos) {
                visitedCoords = append(visitedCoords, pos)
            }
        }
    case "D":
        for j := 0; j < distance; j++ {
            parentPos[0]--
            if parentPos[1] > pos[1] && parentPos[0] - pos[0] < -1 {
                // case for when the Head is above the Tail and the Tail is too far to the right
                pos[0]--
                pos[1]++
            } else if parentPos[1] < pos[1] && parentPos[0] - pos[0] < -1 {
                // case for when the Head is below the Tail and the Tail is too far to the right
                pos[0]--
                pos[1]--
            } else if parentPos[0] - pos[0] < -1 {
                // case for when the Tail is too far to the right
                pos[0]--
            }
            if !contains(visitedCoords, pos) {
                visitedCoords = append(visitedCoords, pos)
            }
        }
    }
    return parentPos, pos, visitedCoords
}

func printHeatmap(visitedCoords [][2]int) {
    maxX, maxY, minX, minY := 0, 0, 0, 0
    for _, coord := range visitedCoords {
        if coord[0] > maxX {
            maxX = coord[0]
        }
        if coord[0] < minX {
            minX = coord[0]
        }
        if coord[1] > maxY {
            maxY = coord[1]
        }
        if coord[1] < minY {
            minY = coord[1]
        }
    }

    for i := maxX; i >= minX; i-- {
        for j := minY; j <= maxY; j++ {
            if i == 0 && j == 0 {
                fmt.Print("S")
            } else if contains(visitedCoords, [2]int{i, j}) {
                fmt.Print("#")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}

func contains(s [][2]int, e [2]int) bool {
    for _, a := range s {
        if a == e {
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
