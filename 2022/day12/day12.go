package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    dijkstra "github.com/RyanCarrier/dijkstra" // is this cheating? :p
)

func main() {
    input, _ := readLines("input.txt")

    var grid [][]int
    var val, S, G int
    var startingPoints []int
    graph := dijkstra.NewGraph()

    // parsing nodes
    for _, line := range input {
        var row []int
        for _, char := range strings.Split(line, "") {
            height := getHeight(char)
            if char == "S" { // start node
                height = 0
                S = val
            } else if char == "E" { // end node
                height = 25
                G = val
            }
            
            if height == 0 {
                startingPoints = append(startingPoints, val)
            }
            graph.AddVertex(val)
            val++
            row = append(row, int(height))
        }
        grid = append(grid, row)
    }

    // parsing edges (neighbours)
    for row := 0; row < len(grid); row++ {
        for col := 0; col < len(grid[row]); col++ {
            for i := row - 1; i <= row+1; i++ {
                if i < 0 || i >= len(grid) {
                    continue
                }
                for j := col - 1; j <= col+1; j++ {
                    if j < 0 || j >= len(grid[i]) {
                        continue
                    }
                    if i == row && j == col {
                        continue
                    }
                    if row != i && col != j {
                        continue
                    }
                    if grid[row][col]+1 >= grid[i][j] {
                        node1 := row*len(grid[0]) + col
                        node2 := i*len(grid[0]) + j
                        graph.AddArc(node1, node2, 1)
                    }
                }
            }
        }
    }

    cost, _ := graph.Shortest(S, G)
    fmt.Println("Part 1 Answer:", cost.Distance)

    min := 32767
    for _, node := range startingPoints {
        cost, _ := graph.Shortest(node, G)
        t := int(cost.Distance)
        if t< min && t != 0 {
            min = t
        }
    }
    fmt.Println("Part 2 Answer:", min)
}

func getHeight(s string) int {
    return int(s[0]) - 97
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