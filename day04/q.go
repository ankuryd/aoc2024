package day04

import (
	"fmt"
	"time"

	"aoc2024/util"
)

type Pos struct {
	x, y int
}

func (p *Pos) InBounds() bool {
	return p.x >= 0 && p.x < rows && p.y >= 0 && p.y < cols
}

func (p *Pos) Move(dir Dir, steps int) Pos {
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

	north = Dir{dx: -1, dy: 0}
	east  = Dir{dx: 0, dy: 1}
	south = Dir{dx: 1, dy: 0}
	west  = Dir{dx: 0, dy: -1}

	northEast = Dir{dx: -1, dy: 1}
	southEast = Dir{dx: 1, dy: 1}
	southWest = Dir{dx: 1, dy: -1}
	northWest = Dir{dx: -1, dy: -1}

	allDirs = []Dir{north, east, south, west, northEast, southEast, southWest, northWest}

	xDirs = []Dir{northEast, southEast}
)

func solve1(inputs []string) string {
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

	return fmt.Sprintf("%d", result)
}

func solve2(inputs []string) string {
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

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	rows = len(input)
	cols = len(input[0])

	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}
	}

	startTime := time.Now()
	util.Output(1, solve1(input))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(input))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
