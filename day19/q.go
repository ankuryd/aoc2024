package day19

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"aoc2024/util"
)

func solve1(patterns, towels []string) string {
	re := regexp.MustCompile("^(" + strings.Join(patterns, "|") + ")*$")

	result := 0
	for _, towel := range towels {
		if re.MatchString(towel) {
			result++
		}
	}

	return fmt.Sprintf("%d", result)
}

func solve2(patterns, towels []string) string {
	possibilities := make(map[string]int)
	var calc func(towel string) int
	calc = func(towel string) int {
		if len(towel) == 0 {
			return 1
		}

		if val, ok := possibilities[towel]; ok {
			return val
		}

		count := 0
		for _, pattern := range patterns {
			if strings.HasPrefix(towel, pattern) {
				count += calc(towel[len(pattern):])
			}
		}

		possibilities[towel] = count
		return count
	}

	result := 0
	for _, towel := range towels {
		result += calc(towel)
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	patterns, towels := make([]string, 0), make([]string, 0)

	readTowels := false
	for _, line := range input {
		if line == "" {
			readTowels = true
			continue
		}

		if readTowels {
			towels = append(towels, line)
		} else {
			patterns = append(patterns, strings.Split(line, ", ")...)
		}
	}

	startTime := time.Now()
	util.Output(1, solve1(patterns, towels))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(patterns, towels))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
