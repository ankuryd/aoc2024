package day2

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

	removalUsed := false
	prevDiff := 0

	for i := 0; i < len(report)-1; i++ {
		currDiff := report[i+1] - report[i]
		if !isDiffValid(currDiff) || (i > 0 && (currDiff > 0) != (prevDiff > 0)) {
			if removalUsed {
				return false
			}

			removalUsed = true

			// Try removing report[i]
			if i > 0 && isDiffValid(report[i+1]-report[i-1]) && ((report[i+1]-report[i-1]) > 0) == (prevDiff > 0) {
				prevDiff = report[i+1] - report[i-1]
			} else {
				// Try removing report[i+1]
				if i+2 < len(report) {
					if isDiffValid(report[i+2]-report[i]) && ((report[i+2]-report[i] > 0) == (prevDiff > 0)) {
						prevDiff = report[i+2] - report[i]
						i++ // Skip the next element as it's considered removed
						continue
					}
				}

				// If neither removal fixes the issue
				return false
			}
		} else {
			prevDiff = currDiff
		}
	}

	return true
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

	for lineNumber, line := range input {
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", lineNumber)
		}

		fields := strings.Fields(line)
		intFields, err := util.ConvertToIntSlice(fields)
		if err != nil {
			log.Fatalf("Error converting '%s' to integer on line %d: %v", line, lineNumber, err)
		}

		reports = append(reports, intFields)
	}

	fmt.Println("Question 1 output:")
	solve1(reports)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(reports)
}
