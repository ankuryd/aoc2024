package day01

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"aoc2024/util"
)

func solve1(list1, list2 []int) {
	if len(list1) != len(list2) {
		log.Fatal("Error: Lists are of different lengths")
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

	fmt.Println(result)
}

func solve2(list1, list2 []int) {
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

	fmt.Println(result)
}

func Run(day int, input []string) {
	var list1, list2 []int

	for lineNumber, line := range input {
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", lineNumber)
		}

		fields := strings.Fields(line)
		if len(fields) != 2 {
			log.Fatalf("Invalid format on line %d: %q", lineNumber, line)
		}

		intFields, err := util.ConvertToIntSlice(fields)
		if err != nil {
			log.Fatalf("Error converting '%s' to integer on line %d: %v", line, lineNumber, err)
		}

		list1 = append(list1, intFields[0])
		list2 = append(list2, intFields[1])
	}

	fmt.Println("Question 1 output:")
	solve1(list1, list2)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(list1, list2)
}
