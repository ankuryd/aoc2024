package day15

import (
	"fmt"
	"time"

	"aoc2024/util"
)

type Pos struct {
	x, y int
}

type Dir struct {
	dx, dy int
}

func (p *Pos) Move(d Dir) Pos {
	return Pos{x: p.x + d.dx, y: p.y + d.dy}
}

type Cell rune

func (c Cell) String() string {
	return string(c)
}

type Warehouse struct {
}

const (
	north = rune('^')
	east  = rune('>')
	south = rune('v')
	west  = rune('<')

	wall     Cell = '#'
	open     Cell = '.'
	box      Cell = 'O'
	robot    Cell = '@'
	leftBox  Cell = '['
	rightBox Cell = ']'
)

var (
	North = Dir{dx: -1, dy: 0}
	East  = Dir{dx: 0, dy: 1}
	South = Dir{dx: 1, dy: 0}
	West  = Dir{dx: 0, dy: -1}

	dirs = map[rune]Dir{
		north: North,
		east:  East,
		south: South,
		west:  West,
	}

	replace = map[Cell][]Cell{
		open:  {open, open},
		wall:  {wall, wall},
		robot: {robot, open},
		box:   {leftBox, rightBox},
	}

	adjacentDir = map[Cell]Dir{
		leftBox:  East,
		rightBox: West,
	}
)

func push(grid [][]Cell, pos Pos, nextPos Pos) {
	cell := grid[nextPos.x][nextPos.y]
	grid[nextPos.x][nextPos.y] = grid[pos.x][pos.y]
	grid[pos.x][pos.y] = cell
}

func canPushBox(grid [][]Cell, pos Pos, move Dir) bool {
	nextPos := pos.Move(move)
	switch grid[nextPos.x][nextPos.y] {
	case wall:
		return false
	case box:
		return canPushBox(grid, nextPos, move)
	}

	return true
}

func pushBox(grid [][]Cell, pos Pos, move Dir) {
	nextPos := pos.Move(move)
	switch grid[nextPos.x][nextPos.y] {
	case open:
		push(grid, pos, nextPos)
	case box:
		pushBox(grid, nextPos, move)
		push(grid, pos, nextPos)
	}
}

func solve1(grid [][]Cell, movements []Dir) string {
	newGrid := make([][]Cell, 0)
	for i, row := range grid {
		newGrid = append(newGrid, make([]Cell, 0))
		newGrid[i] = append(newGrid[i], row...)
	}

	robotPos := Pos{x: -1, y: -1}
	for i, row := range newGrid {
		for j, cell := range row {
			if cell == robot {
				robotPos = Pos{x: i, y: j}
				break
			}
		}

		if robotPos.x != -1 {
			break
		}
	}

	for _, move := range movements {
		nextPos := robotPos.Move(move)
		switch newGrid[nextPos.x][nextPos.y] {
		case box:
			if canPushBox(newGrid, nextPos, move) {
				pushBox(newGrid, nextPos, move)
				push(newGrid, robotPos, nextPos)
				robotPos = nextPos
			}
		case open:
			push(newGrid, robotPos, nextPos)
			robotPos = nextPos
		}
	}

	result := 0
	for i, row := range newGrid {
		for j, cell := range row {
			if cell == box {
				result += 100*i + j
			}
		}
	}

	return fmt.Sprintf("%d", result)
}

func canPushBigBox(grid [][]Cell, pos Pos, move Dir) bool {
	nextPos := pos.Move(move)
	switch grid[nextPos.x][nextPos.y] {
	case wall:
		return false
	case leftBox, rightBox:
		switch move {
		case North, South:
			adjacentPos := nextPos.Move(adjacentDir[grid[nextPos.x][nextPos.y]])
			return canPushBigBox(grid, nextPos, move) && canPushBigBox(grid, adjacentPos, move)
		case East, West:
			return canPushBigBox(grid, nextPos, move)
		}
	}

	return true
}

func pushBigBox(grid [][]Cell, pos Pos, move Dir) {
	nextPos := pos.Move(move)
	switch grid[nextPos.x][nextPos.y] {
	case open:
		push(grid, pos, nextPos)
	case leftBox, rightBox:
		switch move {
		case North, South:
			adjacentPos := nextPos.Move(adjacentDir[grid[nextPos.x][nextPos.y]])
			pushBigBox(grid, nextPos, move)
			pushBigBox(grid, adjacentPos, move)
			push(grid, pos, nextPos)
		case East, West:
			pushBigBox(grid, nextPos, move)
			push(grid, pos, nextPos)
		}
	}
}

func solve2(grid [][]Cell, movements []Dir) string {
	newGrid := make([][]Cell, 0)
	for i, row := range grid {
		newGrid = append(newGrid, make([]Cell, 0))
		for _, cell := range row {
			newGrid[i] = append(newGrid[i], replace[cell]...)
		}
	}

	robotPos := Pos{x: -1, y: -1}
	for i, row := range newGrid {
		for j, cell := range row {
			if cell == robot {
				robotPos = Pos{x: i, y: j}
				break
			}
		}

		if robotPos.x != -1 {
			break
		}
	}

	for _, move := range movements {
		nextPos := robotPos.Move(move)
		switch newGrid[nextPos.x][nextPos.y] {
		case leftBox, rightBox:
			switch move {
			case North, South:
				adjacentPos := nextPos.Move(adjacentDir[newGrid[nextPos.x][nextPos.y]])
				if canPushBigBox(newGrid, nextPos, move) && canPushBigBox(newGrid, adjacentPos, move) {
					pushBigBox(newGrid, nextPos, move)
					pushBigBox(newGrid, adjacentPos, move)
					push(newGrid, robotPos, nextPos)
					robotPos = nextPos
				}
			case East, West:
				if canPushBigBox(newGrid, nextPos, move) {
					pushBigBox(newGrid, nextPos, move)
					push(newGrid, robotPos, nextPos)
					robotPos = nextPos
				}
			}
		case open:
			push(newGrid, robotPos, nextPos)
			robotPos = nextPos
		}
	}

	result := 0
	for i, row := range newGrid {
		for j, cell := range row {
			if cell == leftBox {
				result += 100*i + j
			}
		}
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	grid := make([][]Cell, 0)
	movements := make([]Dir, 0)

	readMovements := false
	for i, line := range input {
		if line == "" {
			readMovements = true
			continue
		}

		if readMovements {
			for _, move := range line {
				movements = append(movements, dirs[move])
			}
		} else {
			grid = append(grid, make([]Cell, len(line)))
			for j, char := range line {
				grid[i][j] = Cell(char)
			}
		}
	}

	startTime := time.Now()
	util.Output(1, solve1(grid, movements))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(grid, movements))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
