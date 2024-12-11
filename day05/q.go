package day05

import (
	"fmt"
	"strings"
	"time"

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

func solve1(orders map[int]map[int]struct{}, updates [][]int) string {
	result := 0
	for _, update := range updates {
		_, _, isInvalid := isInvalid(orders, update)
		if !isInvalid {
			result += update[len(update)/2]
		}
	}

	return fmt.Sprintf("%d", result)
}

func solve2(orders map[int]map[int]struct{}, updates [][]int) string {
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

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	orders := make(map[int]map[int]struct{})
	updates := make([][]int, 0)

	isOrder := true
	for i, line := range input {
		if line == "" {
			isOrder = false
			continue
		}

		if isOrder {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				util.Fatal("Invalid format on line %d: %s", i, line)
			}

			intParts, err := util.ConvertToIntSlice(parts)
			if err != nil {
				util.Fatal("Error converting '%s' to integer on line %d: %v", line, i, err)
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
				util.Fatal("Error converting '%s' to integer on line %d: %v", line, i, err)
			}

			updates = append(updates, intParts)
		}
	}

	startTime := time.Now()
	util.Output(1, solve1(orders, updates))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(orders, updates))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
