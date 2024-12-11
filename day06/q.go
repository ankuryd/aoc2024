package day06

import (
	"fmt"
	"time"

	"aoc2024/util"
)

type Pos struct {
	x, y int
}

func (p Pos) InBounds() bool {
	return p.x >= 0 && p.x < rows && p.y >= 0 && p.y < cols
}

func (p Pos) Move(d Dir) Pos {
	return Pos{x: p.x + d.dx, y: p.y + d.dy}
}

type Cell struct {
	Pos
}

type Dir struct {
	dx, dy int
}

func (d Dir) RotateCW() Dir {
	return Dir{dx: d.dy, dy: -d.dx}
}

type State struct {
	Pos
	Dir
}

const (
	north = rune('^')
	east  = rune('>')
	south = rune('v')
	west  = rune('<')
)

var (
	rows, cols int

	dirs = map[rune]Dir{
		north: {dx: -1, dy: 0},
		east:  {dx: 0, dy: 1},
		south: {dx: 1, dy: 0},
		west:  {dx: 0, dy: -1},
	}
)

func solve1(grid [][]int, startPos Pos, startDir Dir) string {
	visited := make(map[Cell]struct{})
	currentPos := startPos
	currentDir := startDir
	visited[Cell{currentPos}] = struct{}{}

	for {
		nextPos := currentPos.Move(currentDir)
		if !nextPos.InBounds() {
			break
		}

		if grid[nextPos.x][nextPos.y] == 1 {
			currentDir = currentDir.RotateCW()
			continue
		}

		currentPos = nextPos
		currentCell := Cell{currentPos}
		visited[currentCell] = struct{}{}
	}

	return fmt.Sprintf("%d", len(visited)+1)
}

func isLoop(grid [][]int, startPos Pos, startDir Dir) bool {
	seenStates := make(map[State]struct{})
	currentPos := startPos
	currentDir := startDir
	seenStates[State{Pos: currentPos, Dir: currentDir}] = struct{}{}

	for {
		nextPos := currentPos.Move(currentDir)
		if !nextPos.InBounds() {
			return false
		}

		if grid[nextPos.x][nextPos.y] == 1 {
			currentDir = currentDir.RotateCW()
			currentState := State{Pos: currentPos, Dir: currentDir}
			if _, ok := seenStates[currentState]; ok {
				return true
			}
			seenStates[currentState] = struct{}{}
			continue
		}

		currentPos = nextPos
		currentState := State{Pos: currentPos, Dir: currentDir}
		if _, ok := seenStates[currentState]; ok {
			return true
		}
		seenStates[currentState] = struct{}{}
	}
}

func solve2(grid [][]int, startPos Pos, startDir Dir) string {
	visited := make(map[Cell]struct{})
	currentPos := startPos
	currentDir := startDir
	visited[Cell{currentPos}] = struct{}{}

	for {
		nextPos := currentPos.Move(currentDir)
		if !nextPos.InBounds() {
			break
		}

		if grid[nextPos.x][nextPos.y] == 1 {
			currentDir = currentDir.RotateCW()
			continue
		}

		currentPos = nextPos
		currentCell := Cell{currentPos}
		visited[currentCell] = struct{}{}
	}

	result := 0
	for cell := range visited {
		if cell.Pos == startPos {
			continue
		}

		grid[cell.x][cell.y] = 1
		if isLoop(grid, startPos, startDir) {
			result++
		}
		grid[cell.x][cell.y] = 0
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	grid := make([][]int, 0)
	startPos := Pos{-1, -1}
	startDir := Dir{0, 0}

	rows = len(input)
	cols = len(input[0])

	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}

		row := make([]int, 0)

		for j, char := range line {
			switch char {
			case '.':
				row = append(row, 0)
			case '#':
				row = append(row, 1)
			case north, south, west, east:
				row = append(row, 0)
				startPos = Pos{i - 1, j}
				startDir = dirs[char]
			default:
				util.Fatal("Invalid character '%c' on line %d, column %d", char, i, j)
			}
		}

		grid = append(grid, row)
	}

	startTime := time.Now()
	util.Output(1, solve1(grid, startPos, startDir))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(grid, startPos, startDir))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
