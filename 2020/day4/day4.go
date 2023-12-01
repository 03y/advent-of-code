package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, _ := readLines("input.txt")
	input = append(input, "")
	passport := make(map[string]string)
	part1, part2 := 0, 0

	for i := range input {
		if input[i] != "" {
			pattern := "([a-z]{3}):([a-zA-z0-9#]{1,99})"
			re := regexp.MustCompile(pattern)
			match := re.FindAllStringSubmatch(input[i], -1)

			if len(match) > 0 {
				for j := range match {
					passport[match[j][1]] = match[j][2]
				}
			}
		} else {
			if passport["byr"] != "" && passport["iyr"] != "" && passport["eyr"] != "" && passport["hgt"] != "" && passport["hcl"] != "" && passport["ecl"] != "" && passport["pid"] != "" {
				part1++
				byr, _ := strconv.Atoi(passport["byr"])
				if byr >= 1920 && byr <= 2002 {
					iyr, _ := strconv.Atoi(passport["iyr"])
					if iyr >= 2010 && iyr <= 2020 {
						eyr, _ := strconv.Atoi(passport["eyr"])
						if eyr >= 2020 && eyr <= 2030 {
							hcl := passport["hcl"]
							if hcl[0] == '#' {
								ecl := passport["ecl"]
								if ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth" {
									pid := passport["pid"]
									if len(pid) == 9 {
										if passport["hgt"][len(passport["hgt"])-2:] == "cm" {
											hgt, _ := strconv.Atoi(passport["hgt"][:len(passport["hgt"])-2])
											if hgt >= 150 && hgt <= 193 {
												part2++
											}
										} else if passport["hgt"][len(passport["hgt"])-2:] == "in" {
											hgt, _ := strconv.Atoi(passport["hgt"][:len(passport["hgt"])-2])
											if hgt >= 59 && hgt <= 76 {
												part2++
											}
										}
									}
								}
							}
						}
					}
				}
			}
			passport = make(map[string]string)
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
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
