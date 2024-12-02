package main

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

func isValid(report []int) bool {
	if len(report) < 2 {
		return true
	}

	prevDiff := report[1] - report[0]

	if prevDiff < -maxDiff || prevDiff > maxDiff || (prevDiff > -minDiff && prevDiff < minDiff) {
		return false
	}

	for i := 1; i < len(report)-1; i++ {
		currDiff := report[i+1] - report[i]
		if currDiff < -maxDiff || currDiff > maxDiff || 
		(currDiff > -minDiff && currDiff < minDiff) || 
		(currDiff > 0) != (prevDiff > 0) {
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

func main() {
	const filename = "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file '%s': %v", filename, err)
	}
	defer file.Close()

	var reports [][]int
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

	solve1(reports)
	solve2(reports)
}
