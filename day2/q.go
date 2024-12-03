package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

			if i > 0 && isDiffValid(report[i+1]-report[i-1]) && ((report[i+1]-report[i-1] > 0) == (prevDiff > 0)) {
				prevDiff = report[i+1] - report[i-1]
			} else {
				prevDiff = currDiff
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

func Run() {
	const filename = "day2/input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file '%s': %v", filename, err)
	}
	defer file.Close()

	reports := make([][]int, 0)
	lineNumber := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", lineNumber)
		}

		fields := strings.Fields(line)
		report := make([]int, 0, len(fields))

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				log.Fatalf("Error converting '%s' to integer on line %d: %v", field, lineNumber, err)
			}

			report = append(report, num)
		}

		reports = append(reports, report)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file '%s': %v", filename, err)
	}

	fmt.Println("Question 1 output:")
	solve1(reports)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(reports)
}
