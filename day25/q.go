package day25

import (
	"fmt"
	"strings"
	"time"

	"aoc2024/util"
)

func fit(lock, key []int) bool {
	for i := 0; i < len(lock); i++ {
		if lock[i]+key[i] > 5 {
			return false
		}
	}

	return true
}

func solve1(locks, keys [][]int) string {
	result := 0
	for _, lock := range locks {
		for _, key := range keys {
			if fit(lock, key) {
				result++
			}
		}
	}

	return fmt.Sprintf("%d", result)
}

func solve2(locks, keys [][]int) string {
	fmt.Println("There is no part 2 for this problem! TADAAAA!!!")
	return fmt.Sprintf("%d", 0)
}

func Run(day int, input []string) {
	locks, keys := make([][]int, 0), make([][]int, 0)

	curr := make([]string, 0)
	for _, line := range input {
		if line == "" {
			if len(curr) > 0 {
				h, w := len(curr), len(curr[0])
				count := make([]int, w)
				for i := 1; i < h-1; i++ {
					for j := 0; j < w; j++ {
						if curr[i][j] == '#' {
							count[j]++
						}
					}
				}

				switch {
				case curr[0] == strings.Repeat("#", w):
					locks = append(locks, count)
				case curr[h-1] == strings.Repeat("#", w):
					keys = append(keys, count)
				}
			}

			curr = make([]string, 0)
			continue
		}

		curr = append(curr, line)
	}

	if len(curr) > 0 {
		h, w := len(curr), len(curr[0])
		count := make([]int, w)
		for i := 1; i < h-1; i++ {
			for j := 0; j < w; j++ {
				if curr[i][j] == '#' {
					count[j]++
				}
			}
		}

		switch {
		case curr[0] == strings.Repeat("#", w):
			locks = append(locks, count)
		case curr[h-1] == strings.Repeat("#", w):
			keys = append(keys, count)
		}
	}

	startTime := time.Now()
	util.Output(1, solve1(locks, keys))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(locks, keys))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
