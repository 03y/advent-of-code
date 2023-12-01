package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    input, _ := readLines("input.txt")
    var heightmap [][]int
    visible, bestScenicScore := 0, 0

    for i := 0; i < len(input); i++ {
        var row []int
        for j := 0; j < len(input[i]); j++ {
            treeHeight := string(input[i][j])
            treeHeightCasted, _ := strconv.Atoi(treeHeight)
            row = append(row, treeHeightCasted)
        }
        heightmap = append(heightmap, row)
    }

    for i := 0; i < len(heightmap); i++ {
        for j := 0; j < len(heightmap[i]); j++ {
            visiblity, scenicScore := visibilityCheck(heightmap, i, j)
            if visiblity {
                visible++
            }
            if scenicScore > bestScenicScore {
                bestScenicScore = scenicScore
            }
        }
    }

    fmt.Println("Part 1 Answer:", visible)
    fmt.Println("Part 2 Answer:", bestScenicScore)
}

func visibilityCheck(heightmap [][]int, row int, col int) (bool, int) {
    var blocked [4]bool = [4]bool{false, false, false, false}
    var viewDistance [4]int = [4]int{0, 0, 0, 0}

    if row > 0 {
        for i := row - 1; i >= 0; i-- {
            if heightmap[i][col] >= heightmap[row][col] {
                blocked[0] = true
                viewDistance[0]++
                break
            } else {
                viewDistance[0]++
            }
        }
    }

    if col > 0 {
        for i := col - 1; i >= 0; i-- {
            if heightmap[row][i] >= heightmap[row][col] {
                blocked[1] = true
                viewDistance[1]++
                break
            } else {
                viewDistance[1]++
            }
        }
    }

    if row < len(heightmap[0])-1 {
        for i := row + 1; i < len(heightmap[0]); i++ {
            if heightmap[i][col] >= heightmap[row][col] {
                blocked[2] = true
                viewDistance[2]++
                break
            } else {
                viewDistance[2]++
            }
        }
    }

    if col < len(heightmap)-1 {
        for i := col + 1; i < len(heightmap); i++ {
            if heightmap[row][i] >= heightmap[row][col] {
                blocked[3] = true
                viewDistance[3]++
                break
            } else {
                viewDistance[3]++
            }
        }
    }

    c := 4
    for i := range blocked {
        if blocked[i] {
            c--
        }
    }
        return c > 0, viewDistance[0] * viewDistance[1] * viewDistance[2] * viewDistance[3]
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
