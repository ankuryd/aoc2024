package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"aoc2024/util"
)

const (
	DONT = "don't"
)

var (
	re  = regexp.MustCompile(`mul\((\d+),(\d+)\)`)                    // mul(x,y)
	nre = regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\).*?do\(\)`) // mul(x,y)|don't()...do()
)

func solve1(input string) string {
	matches := re.FindAllStringSubmatch(input, -1)
	if len(matches) == 0 {
		util.Fatal("no matches for mul(x,y)")
	}

	result := 0
	for _, match := range matches {
		first, err := strconv.Atoi(match[1])
		if err != nil {
			util.Fatal("Error converting '%s' to integer: %v", match[1], err)
		}

		second, err := strconv.Atoi(match[2])
		if err != nil {
			util.Fatal("Error converting '%s' to integer: %v", match[2], err)
		}

		result += first * second
	}

	return fmt.Sprintf("%d", result)
}

func solve2(input string) string {
	matches := nre.FindAllStringSubmatch(input, -1)
	if len(matches) == 0 {
		util.Fatal("no matches for mul(x,y)|don't()...do()")
	}

	result := 0
	for _, match := range matches {
		if strings.HasPrefix(match[0], DONT) {
			continue
		}

		first, err := strconv.Atoi(match[1])
		if err != nil {
			util.Fatal("Error converting '%s' to integer: %v", match[1], err)
		}

		second, err := strconv.Atoi(match[2])
		if err != nil {
			util.Fatal("Error converting '%s' to integer: %v", match[2], err)
		}

		result += first * second
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}
	}

	combined := strings.Join(input, "")

	startTime := time.Now()
	util.Output(1, solve1(combined))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(combined))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
