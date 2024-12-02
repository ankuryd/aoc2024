package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve1(list1, list2 []int) {
	if len(list1) != len(list2) {
		log.Fatal("Error: Lists are of different lengths")
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var result int
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
		if count, exists := counter[num]; exists {
			result += num * count
		}
	}

	fmt.Println(result)
}

func main() {
	const filename = "input.txt"

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

		firstInt, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalf("Error converting first integer '%s' on line %d: %v", fields[0], lineNumber, err)
		}
		list1 = append(list1, firstInt)

		secondInt, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("Error converting second integer '%s' on line %d: %v", fields[1], lineNumber, err)
		}
		list2 = append(list2, secondInt)

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file '%s': %v", filename, err)
	}

	solve1(list1, list2)
	solve2(list1, list2)
}
