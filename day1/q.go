package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func Run(day int) {
	filename := fmt.Sprintf("day%d/input.txt", day)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file '%s': %v", filename, err)
	}
	defer file.Close()

	var list1, list2 []int
	lineNumber := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
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

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file '%s': %v", filename, err)
	}

	fmt.Println("Question 1 output:")
	solve1(list1, list2)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(list1, list2)
}
