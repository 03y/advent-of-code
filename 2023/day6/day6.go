package main

import (
    "strconv"
    "strings"
    "bufio"
    "fmt"
    "os"
)

func main() {
    input, _ := readLines("input.txt")
    part1, part2 := 0, 0
    var races   []int
    var records []int

    for i := 1; i < len(strings.Fields(input[0])); i++ {
        time, _ := strconv.Atoi(strings.Fields(input[0])[i])
        distance, _ := strconv.Atoi(strings.Fields(input[1])[i])
        races = append(races, time)
        records = append(records, distance)
    }

    for i := range races {
        if i == 0 {
            part1 = getWins(races[i], records[i])
        } else {
            part1 *= getWins(races[i], records[i])
        }
    }

    bigRace, _ := strconv.Atoi(getDigits(input[0]))
    bigRecord, _ := strconv.Atoi(getDigits(input[1]))
    part2 = getWins(bigRace, bigRecord)

    fmt.Println("Part 1:", part1)
    fmt.Println("Part 2:", part2)
}

func getDigits(s string) string {
    result := ""
    for _, c := range s {
        switch string(c) {
        case "0":
            result += string(c)
            break
        case "1":
            result += string(c)
            break
        case "2":
            result += string(c)
            break
        case "3":
            result += string(c)
            break
        case "4":
            result += string(c)
            break
        case "5":
            result += string(c)
            break
        case "6":
            result += string(c)
            break
        case "7":
            result += string(c)
            break
        case "8":
            result += string(c)
            break
        case "9":
            result += string(c)
            break
        }
    }
    return result
}

func getWins(race int, record int) int {
    wins := 0
    for j := 0; j <= race; j++ {
        if getDist(j, race) > record {
            wins++
        }
    }
    return wins
}

func getDist(charge int, time int) int {
    speed := 0
    time-=charge
    speed+=charge
    return speed*time
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
