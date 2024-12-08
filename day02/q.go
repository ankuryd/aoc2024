package day02

import (
	"fmt"
	"log"
	"strings"

	"aoc2024/util"
)

const (
	minDiff = 1
	maxDiff = 3
)

func isDiffValid(diff int) bool {
	return !(diff < -maxDiff || diff > maxDiff || (diff > -minDiff && diff < minDiff))
}

func isValid(report []int) bool {
	if len(report) < 2 {
		return true
	}

	prevDiff := report[1] - report[0]

	if !isDiffValid(prevDiff) {
		return false
	}

	for i := 1; i < len(report)-1; i++ {
		currDiff := report[i+1] - report[i]
		if !isDiffValid(currDiff) || (currDiff > 0) != (prevDiff > 0) {
			return false
		}

		prevDiff = currDiff
	}

	return true
}

func solve1(reports [][]int) {
	if len(reports) == 0 {
		log.Fatal("Error: No reports to process")
	}

	result := 0
	for _, report := range reports {
		if isValid(report) {
			result++
		}
	}

	fmt.Println(result)
}

func isValidWithOneRemoved(report []int) bool {
	if len(report) < 2 {
		return true
	}

	for i := 0; i < len(report); i++ {
		slice := make([]int, len(report))
		copy(slice, report)
		if isValid(append(slice[:i], slice[i+1:]...)) {
			return true
		}
	}

	return false
}

func solve2(reports [][]int) {
	if len(reports) == 0 {
		log.Fatal("Error: No reports to process")
	}

	result := 0
	for _, report := range reports {
		if isValid(report) || isValidWithOneRemoved(report) {
			result++
		}
	}

	fmt.Println(result)
}

func Run(day int, input []string) {
	reports := make([][]int, 0)

	for i, line := range input {
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", i)
		}

		fields := strings.Fields(line)
		intFields, err := util.ConvertToIntSlice(fields)
		if err != nil {
			log.Fatalf("Error converting '%s' to integer on line %d: %v", line, i, err)
		}

		reports = append(reports, intFields)
	}

	fmt.Println("Question 1 output:")
	solve1(reports)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(reports)
}
