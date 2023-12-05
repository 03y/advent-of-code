package main

import (
    "strconv"
    "strings"
    "regexp"
    "bufio"
    "math"
    "fmt"
    "os"
)

var seedToSoil [][]int
var soilToFert [][]int
var fertToWater [][]int
var waterToLight [][]int
var lightToTemp [][]int
var tempToHumid [][]int
var humidToLoc [][]int

func main() {
    input, _ := readLines("input.txt")
    re := regexp.MustCompile(`\d{2,9}`)
    part1 := math.MaxInt64
    part2 := math.MaxInt64
    var seeds []int

    for i := range input {
        if i == 0 {
            matches := re.FindAllStringSubmatch(input[i], -1)
            if len(matches) > 0 {
                for _, j := range matches {
                    n, _ := strconv.Atoi(j[0])
                    seeds = append(seeds, n)
                }
            }
        } else if strings.Contains(input[i], "seed-to-soil") {
            seedToSoil = parse(input, i)
        } else if strings.Contains(input[i], "soil-to-fert") {
            soilToFert = parse(input, i)
        } else if strings.Contains(input[i], "fertilizer-to-water") {
            fertToWater = parse(input, i)
        } else if strings.Contains(input[i], "water-to-light") {
            waterToLight = parse(input, i)
        } else if strings.Contains(input[i], "light-to-temp") {
            lightToTemp = parse(input, i)
        } else if strings.Contains(input[i], "temperature-to-humid") {
            tempToHumid = parse(input, i)
        } else if strings.Contains(input[i], "humidity-to-loc") {
            humidToLoc = parse(input, i)
        }
    }

    run(seeds, &part1)
    fmt.Println("Part 1:", part1)

    // todo: optimisation:
    // sort the ranges by size, then just keep track of the largest number we have calulcated so far
    // this way we dont have a map that is gigabytes in size

    // todo: another optimisation: rewrite to work backwards

    done := make(map[int]bool)
    for i := 0; i < len(seeds); i+=2 {
        for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
            if done[j] {
                continue
            }
            run([]int{j}, &part2)
            done[j] = true
        }
    }
    fmt.Println("Part 2:", part2)
}

func run(seeds []int, dest *int) {
    result := 0
    for i, seed := range seeds {
        if i == 0 {
            result = seed
        }
        seed = translate(seed, seedToSoil)
        seed = translate(seed, soilToFert)
        seed = translate(seed, fertToWater)
        seed = translate(seed, waterToLight)
        seed = translate(seed, lightToTemp)
        seed = translate(seed, tempToHumid)
        seed = translate(seed, humidToLoc)
        result = min(seed, result)
    }
    *dest = min(result, *dest)
}

func translate(input int, matrix [][]int) int {
    for i := range matrix {
        destStart := matrix[i][0]
        sourceStart := matrix[i][1]
        translateRange := matrix[i][2]
        offset := -(sourceStart-input)

        if input >= sourceStart && offset < translateRange {
            return destStart+offset
        }
    }
    return input
}

func parse(input []string, pos int) [][]int {
    var result [][]int
    loaded := false
    re := regexp.MustCompile(`(\d{1,10})\ (\d{1,10})\ (\d{1,10})`)
    for i := pos + 1; i < len(input); i++ {
        if input[i] != "" && !loaded {
            matches := re.FindAllStringSubmatch(input[i], -1)
            for _, match := range matches {
                var temp []int
                for j, group := range match {
                    if j > 0 {
                        n, _ := strconv.Atoi(group)
                        temp = append(temp, n)
                    }
                }
                result = append(result, temp)
            }
        } else {
            loaded = true
        }
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
