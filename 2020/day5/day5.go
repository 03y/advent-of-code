package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
    input, _ := readLines("input.txt")
	sids := []int{}

    for i := range input {
        if input[i] != "" {
            var rowBin [7]int
            var colBin [3]int

            for j := range input[i] {
                if j < 7 {
                    if input[i][j] == 'B' {
                        rowBin[j] = 1
                    } else {
                        rowBin[j] = 0
                    }
                } else {
					if input[i][j] == 'R' {
						colBin [j-7] = 1
					} else {
						colBin [j-7] = 0
					}
				}
            }
            row := binToDec(rowBin[:])
            col := binToDec(colBin[:])
			sids = append(sids, row*8+col)
        }
    }

	sort.Ints(sids)
	fmt.Println("Part 1:", sids[len(sids)-1])

	for i := 1; i < len(sids); i++ {
		if sids[i]-1 != sids[i-1] {
			fmt.Println("Part 2:", sids[i]-1)
			return
		}
	}
}

func binToDec(bin []int) int {
    n := 0
    for i := range bin {
        n = (2 * n) + bin[i]
    }
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
