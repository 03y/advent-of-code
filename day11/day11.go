package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

type Monkey struct {
	id 				int
	items			[]int
	operation		string
	divideBy		int
	targetIfTrue 	int
	targetIfFalse	int
}

func main() {
	input, _ := readLines("input.txt")
	var monkeys []Monkey

	for i := 0; i < len(input); i += 7 {
		monkey := parseMonkey(input[i:i+7], i/7)
		monkeys = append(monkeys, monkey)
	}

	fmt.Println("note: it gives the right answer about 50% of the time, I dont think Go's int datatype is good with big numbers...")
	fmt.Println("Part 1 Answer:", simulateMonkeys(&monkeys, 20, 1))
	fmt.Println("Part 2 Answer:", simulateMonkeys(&monkeys, 20, 2))
}

func printMonkeyItems(monkeys *[]Monkey) {
	for i := 0; i < len(*monkeys); i++ {
		fmt.Print("Monkey ", (*monkeys)[i].id, ": ")
		for j := 0; j < len((*monkeys)[i].items); j++ {
			fmt.Print((*monkeys)[i].items[j], " ")
		}
		fmt.Println()
	}
}

func parseOperation(operation string) (string, int) {
	words := strings.Fields(operation)
	num, _ := strconv.Atoi(words[1])
	return words[0], num
}

func simulateMonkeys(monkeys *[]Monkey, rounds int, part int) int {
	inspectionCount := make(map[int]int)
	verbose := false

	for i := 0; i < rounds; i++ {
		for j := 0; j < len(*monkeys); j++ {
			if verbose {
				fmt.Printf("Monkey %d: \n", (*monkeys)[j].id)
			}
			
			if len((*monkeys)[j].items) > 0 {
				itemsToRemove := make([]int, 0)
				for k := 0; k < len((*monkeys)[j].items); k++ {
					if verbose {
						fmt.Println("\tMonkey inspects item with a worry level of", (*monkeys)[j].items[k])
					}		
					inspectionCount[(*monkeys)[j].id]++

					operation, num := parseOperation((*monkeys)[j].operation)
					if num == 0 {
						num = (*monkeys)[j].items[k]
					}
					switch operation {
					case "*":
						if verbose {
							fmt.Println("\t\tWorry level multiplied by", num, "to", (*monkeys)[j].items[k]*num)
						}
						(*monkeys)[j].items[k] *= num
					case "+":
						if verbose {
							fmt.Println("\t\tWorry level increases by", num, "to", (*monkeys)[j].items[k]+num)
						}
						(*monkeys)[j].items[k] += num
					}

					if part == 1{
						if verbose {
							fmt.Println("\t\tMonkey gets bored with item. Worry level is divided by", 3, "to", (*monkeys)[j].items[k]/3)
						}
						(*monkeys)[j].items[k] /= 3
					}

					if (*monkeys)[j].items[k] % (*monkeys)[j].divideBy == 0 {
						if verbose {
							fmt.Println("\t\tCurrent worry level is divisible by", (*monkeys)[j].divideBy)
							fmt.Println("\t\tItem with worry level", (*monkeys)[j].items[k], "is thrown to monkey", (*monkeys)[j].targetIfTrue)
						}
						(*monkeys)[(*monkeys)[j].targetIfTrue].items = append((*monkeys)[(*monkeys)[j].targetIfTrue].items, (*monkeys)[j].items[k])
						itemsToRemove = append(itemsToRemove, k)
					} else {
						if verbose {
							fmt.Println("\t\tCurrent worry level is not divisible by", (*monkeys)[j].divideBy)
							fmt.Println("\t\tItem with worry level", (*monkeys)[j].items[k], "is thrown to monkey", (*monkeys)[j].targetIfFalse)
						}
						(*monkeys)[(*monkeys)[j].targetIfFalse].items = append((*monkeys)[(*monkeys)[j].targetIfFalse].items, (*monkeys)[j].items[k])
						itemsToRemove = append(itemsToRemove, k)
					}
				}

				for k := len(itemsToRemove)-1; k >= 0; k-- {
					(*monkeys)[j].items = append((*monkeys)[j].items[:itemsToRemove[k]], (*monkeys)[j].items[itemsToRemove[k]+1:]...)
				}
			}
		}
		
		if verbose {
			fmt.Println("After round", i+1, ", the monkeys are holding items with these worry levels:")
			printMonkeyItems(monkeys)
			fmt.Println()
		}
	}
	
	var max1, max2 int
	for _, v := range inspectionCount {
		if v > max1 {
			max1 = v
		} else if v > max2 {
			max2 = v
		}
	}
	return max1*max2
}

func (m Monkey) String() string {
	return fmt.Sprintf("Monkey %d: %v %s %d %d %d", m.id, m.items, m.operation, m.divideBy, m.targetIfTrue, m.targetIfFalse)
}

func parseMonkey(input []string, id int) Monkey {
	var monkey Monkey

	monkey.id = id

	words := strings.Fields(input[1])
	for j := 2; j < len(words); j++ {
		words[j] = strings.TrimSuffix(words[j], ",")
		words[j] = strings.TrimSuffix(words[j], " ")
		num, _ := strconv.Atoi(words[j])
		monkey.items = append(monkey.items, num)
	}
	
	monkey.operation = input[2][23:]

	words = strings.Fields(input[3])
	monkey.divideBy, _ = strconv.Atoi(words[len(words)-1])

	words = strings.Fields(input[4])
	monkey.targetIfTrue, _ = strconv.Atoi(words[len(words)-1])

	words = strings.Fields(input[5])
	monkey.targetIfFalse, _ = strconv.Atoi(words[len(words)-1])

	return monkey
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
