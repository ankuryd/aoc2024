package day01

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"aoc2024/util"
)

func solve1(list1, list2 []int) string {
	if len(list1) != len(list2) {
		util.Fatal("Error: Lists are of different lengths")
	}

	sort.Ints(list1)
	sort.Ints(list2)

	result := 0
	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff
		}

		result += diff
	}

	return fmt.Sprintf("%d", result)
}

func solve2(list1, list2 []int) string {
	counter := make(map[int]int)
	for _, num := range list2 {
		counter[num]++
	}

	var result int
	for _, num := range list1 {
		if count, ok := counter[num]; ok {
			result += num * count
		}
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	var list1, list2 []int

	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}

		fields := strings.Fields(line)
		if len(fields) != 2 {
			util.Fatal("Invalid format on line %d: %q", i, line)
		}

		intFields, err := util.ConvertToIntSlice(fields)
		if err != nil {
			util.Fatal("Error converting '%s' to integer on line %d: %v", line, i, err)
		}

		list1 = append(list1, intFields[0])
		list2 = append(list2, intFields[1])
	}

	startTime := time.Now()
	util.Output(1, solve1(list1, list2))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(list1, list2))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
