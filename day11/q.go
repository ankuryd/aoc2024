package day11

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"aoc2024/util"
)

func solve1(stones []int) string {
	for i := 0; i < 25; i++ {
		next := make([]int, 0)
		for _, stone := range stones {
			if stone == 0 { // # rule 1
				next = append(next, 1)
			} else if len(strconv.Itoa(stone))%2 == 0 { // # rule 2
				half := len(strconv.Itoa(stone)) / 2
				tens := int(math.Pow(10, float64(half)))
				next = append(next, stone/tens)
				next = append(next, stone%tens)
			} else { // # rule 3
				next = append(next, stone*2024)
			}
		}

		stones = next
	}

	return fmt.Sprintf("%d", len(stones))
}

func solve2(stones []int) string {
	stoneCount := make(map[int]int)
	for _, stone := range stones {
		stoneCount[stone]++
	}

	for i := 0; i < 75; i++ {
		nextCount := make(map[int]int)
		for stone, count := range stoneCount {
			if stone == 0 { // # rule 1
				nextCount[1] += count
			} else if len(strconv.Itoa(stone))%2 == 0 { // # rule 2
				half := len(strconv.Itoa(stone)) / 2
				tens := int(math.Pow(10, float64(half)))
				nextCount[stone/tens] += count
				nextCount[stone%tens] += count
			} else { // # rule 3
				nextCount[stone*2024] += count
			}
		}

		stoneCount = nextCount
	}

	result := 0
	for _, count := range stoneCount {
		result += count
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}
	}

	fields := strings.Fields(input[0])
	stones, err := util.ConvertToIntSlice(fields)
	if err != nil {
		util.Fatal("Invalid format on line %d: %s, expected ints, got %v", 0, fields, err)
	}

	startTime := time.Now()
	util.Output(1, solve1(stones))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(stones))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
