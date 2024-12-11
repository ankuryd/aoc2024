package day15

import (
	"fmt"
	"time"

	"aoc2024/util"
)

func solve1() string {
	return fmt.Sprintf("%d", 0)
}

func solve2() string {
	return fmt.Sprintf("%d", 0)
}

func Run(day int, input []string) {
	util.Fatal("Error: Day '%d' not implemented.", day)

	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}
	}

	startTime := time.Now()
	util.Output(1, solve1())
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2())
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
