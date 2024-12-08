package day04

import (
	"fmt"
	"log"
)

type Pos struct {
	x, y int
}

func (p Pos) InBounds() bool {
	return p.x >= 0 && p.x < rows && p.y >= 0 && p.y < cols
}

func (p Pos) Move(dir Dir, steps int) Pos {
	return Pos{x: p.x + dir.dx*steps, y: p.y + dir.dy*steps}
}

type Dir struct {
	dx, dy int
}

const (
	X = byte('X')
	M = byte('M')
	A = byte('A')
	S = byte('S')
)

var (
	rows, cols int

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

func solve1(inputs []string) {
	result := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if inputs[i][j] != X {
				continue
			}

			pos := Pos{x: i, y: j}

			pattern := []byte{M, A, S}
			for _, dir := range allDirs {
				matched := true
				for k, char := range pattern {
					nextPos := pos.Move(dir, k+1)
					if !nextPos.InBounds() {
						matched = false
						break
					}

					if inputs[nextPos.x][nextPos.y] != char {
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
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if inputs[i][j] != A {
				continue
			}

			pos := Pos{x: i, y: j}

			pattern := []byte{M, S}
			matched := true
			for _, dir := range xDirs {
				nextPos := pos.Move(dir, 1)
				prevPos := pos.Move(dir, -1)

				if !nextPos.InBounds() || !prevPos.InBounds() {
					matched = false
					break
				}

				if !((inputs[nextPos.x][nextPos.y] == pattern[0] && inputs[prevPos.x][prevPos.y] == pattern[1]) ||
					(inputs[nextPos.x][nextPos.y] == pattern[1] && inputs[prevPos.x][prevPos.y] == pattern[0])) {
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
	rows = len(input)
	cols = len(input[0])

	for i, line := range input {
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", i)
		}
	}

	fmt.Println("Question 1 output:")
	solve1(input)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(input)
}
