package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

type State struct {
	Pos
	Dir
}

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
			currentDir = Dir{di: currentDir.dj, dj: -currentDir.di}
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
			currentDir = Dir{di: currentDir.dj, dj: -currentDir.di}
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
			currentDir = Dir{di: currentDir.dj, dj: -currentDir.di}
			continue
		}

		currentPos = Pos{i: ni, j: nj}
		currentCell := Cell{currentPos}
		visited[currentCell] = struct{}{}
	}

	result := 0
	for cell := range visited {
		if grid[cell.i][cell.j] == 1 || cell.Pos == startPos {
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

func Run(day int) {
	filename := fmt.Sprintf("day%d/input.txt", day)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file '%s': %v", filename, err)
	}
	defer file.Close()

	grid := make([][]int, 0)
	pos := Pos{-1, -1}
	dir := Dir{0, 0}
	lineNumber := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
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
			case '^', 'v', '<', '>':
				pos = Pos{lineNumber - 1, i}
				switch char {
				case '^': // North
					dir.di = -1
				case 'v': // South
					dir.di = 1
				case '<': // West
					dir.dj = -1
				case '>': // East
					dir.dj = 1
				}
			default:
				log.Fatalf("Invalid character '%c' on line %d, column %d", char, lineNumber, i)
			}
		}

		grid = append(grid, row)

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file '%s': %v", filename, err)
	}

	fmt.Println("Question 1 output:")
	solve1(grid, pos, dir)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(grid, pos, dir)
}
