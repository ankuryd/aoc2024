package day4

import (
	"fmt"
	"log"
)

type Dir struct {
	di, dj int
}

var (
	north = Dir{0, -1}
	east  = Dir{1, 0}
	south = Dir{0, 1}
	west  = Dir{-1, 0}

	northEast = Dir{1, -1}
	southEast = Dir{1, 1}
	southWest = Dir{-1, 1}
	northWest = Dir{-1, -1}

	allDirs = []Dir{north, east, south, west, northEast, southEast, southWest, northWest}

	xDirs = []Dir{northEast, southEast}
)

var (
	X = byte('X')
	M = byte('M')
	A = byte('A')
	S = byte('S')
)

func solve1(inputs []string) {
	result := 0
	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs[i]); j++ {
			if inputs[i][j] != X {
				continue
			}

			pattern := []byte{M, A, S}
			for _, dir := range allDirs {
				matched := true
				for k, char := range pattern {
					ni, nj := i+(k+1)*dir.di, j+(k+1)*dir.dj
					if ni < 0 || ni >= len(inputs) || nj < 0 || nj >= len(inputs[ni]) {
						matched = false
						break
					}

					if inputs[ni][nj] != char {
						matched = false
						break
					}
				}

				if matched {
					result++
				}
			}
		}
	}

	fmt.Println(result)
}

func solve2(inputs []string) {
	result := 0
	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs[i]); j++ {
			if inputs[i][j] != A {
				continue
			}

			pattern := []byte{M, S}
			matched := true
			for _, dir := range xDirs {
				ni1, nj1 := i+dir.di, j+dir.dj
				ni2, nj2 := i-dir.di, j-dir.dj

				if ni1 < 0 || ni1 >= len(inputs) || nj1 < 0 || nj1 >= len(inputs[ni1]) ||
					ni2 < 0 || ni2 >= len(inputs) || nj2 < 0 || nj2 >= len(inputs[ni2]) {
					matched = false
					break
				}

				if !((inputs[ni1][nj1] == pattern[0] && inputs[ni2][nj2] == pattern[1]) ||
					(inputs[ni1][nj1] == pattern[1] && inputs[ni2][nj2] == pattern[0])) {
					matched = false
					break
				}
			}

			if matched {
				result++
			}
		}
	}

	fmt.Println(result)
}

func Run(day int, input []string) {
	for lineNumber, line := range input {
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", lineNumber)
		}
	}

	fmt.Println("Question 1 output:")
	solve1(input)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(input)
}
