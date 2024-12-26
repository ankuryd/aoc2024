package day22

import (
	"fmt"
	"time"

	"aoc2024/util"
)

const (
	CYCLES = 2000

	MOD = 16777216
)

func transform(val int) int {
	// Perform the bit-shift transformations inline
	val = ((val << 6) ^ val) % MOD
	val = ((val >> 5) ^ val) % MOD
	val = ((val << 11) ^ val) % MOD
	return val
}

func solve1(secrets []int) string {
	result := 0
	for _, secret := range secrets {
		for i := 0; i < CYCLES; i++ {
			secret = transform(secret)
		}

		result += secret
	}

	return fmt.Sprintf("%d", result)
}

func solve2(secrets []int) string {
	sequences := make(map[string]map[int]struct{})
	sums := make(map[string]int)

	result := 0
	for _, secret := range secrets {
		changes := make([]int, 0, CYCLES)

		curr := secret
		for i := 0; i < CYCLES; i++ {
			prev := curr
			curr = transform(curr)

			changes = append(changes, curr%10-prev%10)

			if i >= 3 {
				bytes := make([]byte, 0, 8)
				for j := i - 3; j <= i; j++ {
					bytes = append(bytes, byte(changes[j]))
				}
				hash := string(bytes)

				if _, ok := sequences[hash]; !ok {
					sequences[hash] = make(map[int]struct{})
					sums[hash] = 0
				}

				if _, ok := sequences[hash][secret]; !ok {
					sequences[hash][secret] = struct{}{}
					sums[hash] += curr % 10

					if sums[hash] > result {
						result = sums[hash]
					}
				}
			}
		}
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	secrets := make([]int, 0)

	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}

		secrets = append(secrets, util.ConvertToInt(line))
	}

	startTime := time.Now()
	util.Output(1, solve1(secrets))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(secrets))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
