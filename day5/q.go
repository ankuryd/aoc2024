package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"aoc2024/util"
)

func isInvalid(orders map[int]map[int]struct{}, update []int) (int, int, bool) {
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if val, ok := orders[update[j]]; ok {
				if _, ok := val[update[i]]; ok {
					return i, j, true
				}
			}
		}
	}

	return -1, -1, false
}

func solve1(orders map[int]map[int]struct{}, updates [][]int) {
	result := 0
	for _, update := range updates {
		_, _, isInvalid := isInvalid(orders, update)
		if !isInvalid {
			result += update[len(update)/2]
		}
	}

	fmt.Println(result)
}

func solve2(orders map[int]map[int]struct{}, updates [][]int) {
	result := 0
	for _, update := range updates {
		swapped := false
		for {
			i, j, isInvalid := isInvalid(orders, update)
			if isInvalid {
				update[i], update[j] = update[j], update[i]
				swapped = true
			} else {
				break
			}
		}

		if swapped {
			result += update[len(update)/2]
		}
	}

	fmt.Println(result)
}

func Run(day int) {
	filename := fmt.Sprintf("day%d/input.txt", day)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file '%s': %v", filename, err)
	}
	defer file.Close()

	orders := make(map[int]map[int]struct{})
	updates := make([][]int, 0)
	lineNumber := 1

	scanner := bufio.NewScanner(file)
	isOrder := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			isOrder = false
			continue
		}

		if isOrder {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				log.Fatalf("Invalid format on line %d: %s", lineNumber, line)
			}

			intParts, err := util.ConvertToIntSlice(parts)
			if err != nil {
				log.Fatalf("Error converting '%s' to integer on line %d: %v", line, lineNumber, err)
			}

			u, v := intParts[0], intParts[1]
			if _, ok := orders[u]; !ok {
				orders[u] = make(map[int]struct{})
			}

			orders[u][v] = struct{}{}
		} else {
			parts := strings.Split(line, ",")
			intParts, err := util.ConvertToIntSlice(parts)
			if err != nil {
				log.Fatalf("Error converting '%s' to integer on line %d: %v", line, lineNumber, err)
			}

			updates = append(updates, intParts)
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file '%s': %v", filename, err)
	}

	fmt.Println("Question 1 output:")
	solve1(orders, updates)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(orders, updates)
}
