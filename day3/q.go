package day3

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	do   = regexp.MustCompile(`do\(\)`)
	dont = regexp.MustCompile(`don\'t\(\)`)
)

func solve1(input string) {
	matches := re.FindAllStringSubmatch(input, -1)
	if len(matches) == 0 {
		fmt.Println("no matches for mul(x,y)")
		return
	}

	result := 0
	for _, match := range matches {
		first, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatalf("Error converting '%s' to integer: %v", match[1], err)
		}

		second, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatalf("Error converting '%s' to integer: %v", match[2], err)
		}

		result += first * second
	}

	fmt.Println(result)
}

func findValidValues(pos, neg, vals []int) []int {
	validValues := make([]int, 0)
	i, j, k := 0, 0, 0

	for i < len(pos) {
		for j < len(neg) && neg[j] < pos[i] {
			j++
		}

		for k < len(vals) && vals[k] < pos[i] {
			k++
		}

		for k < len(vals) && vals[k] >= pos[i] && vals[k] < neg[j] {
			validValues = append(validValues, vals[k])
			k++
		}

		for i < len(pos) && pos[i] < neg[j] {
			i++
		}
	}

	return validValues
}

func solve2(input string) {
	doIndices := do.FindAllStringIndex(input, -1)
	dos := make([]int, len(doIndices))
	for j, index := range doIndices {
		dos[j] = index[0]
	}
	dos = append([]int{0}, dos...)

	dontIndices := dont.FindAllStringIndex(input, -1)
	donts := make([]int, len(dontIndices))
	for j, index := range dontIndices {
		donts[j] = index[0]
	}
	donts = append(donts, math.MaxInt)

	mulIndices := re.FindAllStringIndex(input, -1)
	muls := make([]int, len(mulIndices))
	for j, index := range mulIndices {
		muls[j] = index[0]
	}

	matches := re.FindAllStringSubmatch(input, -1)
	if len(matches) == 0 {
		fmt.Println("no matches for mul(x,y)")
		return
	}

	validMuls := make(map[int]struct{})
	if len(donts) != 0 {
		for _, v := range findValidValues(dos, donts, muls) {
			validMuls[v] = struct{}{}
		}
	}

	result := 0
	for i, match := range matches {
		mulIndex := muls[i]
		if _, ok := validMuls[mulIndex]; !ok {
			continue
		}

		first, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatalf("Error converting '%s' to integer: %v", match[1], err)
		}

		second, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatalf("Error converting '%s' to integer: %v", match[2], err)
		}

		result += first * second
	}

	fmt.Println(result)
}

func Run() {
	const filename = "day3/input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file '%s': %v", filename, err)
	}
	defer file.Close()

	inputs := make([]string, 0)
	lineNumber := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", lineNumber)
		}

		inputs = append(inputs, line)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file '%s': %v", filename, err)
	}

	input := strings.Join(inputs, "")

	fmt.Println("Question 1 output:")
	solve1(input)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(input)
}
