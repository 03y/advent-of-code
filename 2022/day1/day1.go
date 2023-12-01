package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    input, _ := readLines("input.txt")
    var all_elf_calories []int
    elf_calories := 0

    for i := range input {
        if input[i] != "" {
            food_item_calories, _ := strconv.Atoi(input[i])
            elf_calories += food_item_calories
        } else {
            all_elf_calories = append(all_elf_calories, elf_calories)
            elf_calories = 0
        }
    }
    all_elf_calories = append(all_elf_calories, elf_calories)
    fmt.Println("Part 1 Answer:", get_largest(all_elf_calories))

    total := 0
    for i := 0; i < 3; i++ {
        current_largest := get_largest(all_elf_calories)
        total += current_largest

        for j := range all_elf_calories {
            if all_elf_calories[j] == current_largest {
                all_elf_calories = append(all_elf_calories[:j], all_elf_calories[j+1:]...)
                break
            }
        }
    }
    fmt.Println("Part 2 Answer:", total)
}

func get_largest(arr []int) int {
    var largest int
    for _, v := range arr {
        if v > largest {
            largest = v
        }
    }
    return largest
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
