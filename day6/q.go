package day6

import (
	"fmt"
	"log"
)

type Pos struct {
	i, j int
}

type Cell struct {
	Pos
}

type Dir struct {
	di, dj int
}

func (d Dir) RotateCW() Dir {
	return Dir{di: d.dj, dj: -d.di}
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
	dirs = map[rune]Dir{
		north: {di: -1, dj: 0},
		east:  {di: 0, dj: 1},
		south: {di: 1, dj: 0},
		west:  {di: 0, dj: -1},
	}
)

func solve1(grid [][]int, startPos Pos, startDir Dir) {
	visited := make(map[Cell]struct{})
	currentPos := startPos
	currentDir := startDir
	visited[Cell{currentPos}] = struct{}{}

	for {
		ni, nj := currentPos.i+currentDir.di, currentPos.j+currentDir.dj
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[ni]) {
			break
		}

		if grid[ni][nj] == 1 {
			currentDir = currentDir.RotateCW()
			continue
		}

		currentPos = Pos{i: ni, j: nj}
		currentCell := Cell{currentPos}
		visited[currentCell] = struct{}{}
	}

	fmt.Println(len(visited))
}

func isLoop(grid [][]int, startPos Pos, startDir Dir) bool {
	seenStates := make(map[State]struct{})
	currentPos := startPos
	currentDir := startDir
	seenStates[State{Pos: currentPos, Dir: currentDir}] = struct{}{}

	for {
		ni, nj := currentPos.i+currentDir.di, currentPos.j+currentDir.dj
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[ni]) {
			return false
		}

		if grid[ni][nj] == 1 {
			currentDir = currentDir.RotateCW()
			currentState := State{Pos: currentPos, Dir: currentDir}
			if _, ok := seenStates[currentState]; ok {
				return true
			}
			seenStates[currentState] = struct{}{}
			continue
		}

		currentPos = Pos{i: ni, j: nj}
		currentState := State{Pos: currentPos, Dir: currentDir}
		if _, ok := seenStates[currentState]; ok {
			return true
		}
		seenStates[currentState] = struct{}{}
	}
}

func solve2(grid [][]int, startPos Pos, startDir Dir) {
	visited := make(map[Cell]struct{})
	currentPos := startPos
	currentDir := startDir
	visited[Cell{currentPos}] = struct{}{}

	for {
		ni, nj := currentPos.i+currentDir.di, currentPos.j+currentDir.dj
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[ni]) {
			break
		}

		if grid[ni][nj] == 1 {
			currentDir = currentDir.RotateCW()
			continue
		}

		currentPos = Pos{i: ni, j: nj}
		currentCell := Cell{currentPos}
		visited[currentCell] = struct{}{}
	}

	result := 0
	for cell := range visited {
		if cell.Pos == startPos {
			continue
		}

		grid[cell.i][cell.j] = 1
		if isLoop(grid, startPos, startDir) {
			result++
		}
		grid[cell.i][cell.j] = 0
	}

	fmt.Println(result)
}

func Run(day int, input []string) {
	grid := make([][]int, 0)
	startPos := Pos{-1, -1}
	startDir := Dir{0, 0}

	for lineNumber, line := range input {
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", lineNumber)
		}

		row := make([]int, 0)

		for i, char := range line {
			switch char {
			case '.':
				row = append(row, 0)
			case '#':
				row = append(row, 1)
			case north, south, west, east:
				row = append(row, 0)
				startPos = Pos{lineNumber - 1, i}
				startDir = dirs[char]
			default:
				log.Fatalf("Invalid character '%c' on line %d, column %d", char, lineNumber, i)
			}
		}

		grid = append(grid, row)
	}

	fmt.Println("Question 1 output:")
	solve1(grid, startPos, startDir)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(grid, startPos, startDir)
}
