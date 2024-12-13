package day10

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

func (p *Pos) Move(d Dir) Pos {
	return Pos{x: p.x + d.dx, y: p.y + d.dy}
}

type Dir struct {
	dx, dy int
}

var (
	rows, cols int

	north = Dir{dx: -1, dy: 0}
	east  = Dir{dx: 0, dy: 1}
	south = Dir{dx: 1, dy: 0}
	west  = Dir{dx: 0, dy: -1}

	allDirs = []Dir{north, east, south, west}
)

func walk(grid [][]int, pos Pos, repeat bool) int {
	count := 0

	visited := map[Pos]struct{}{pos: {}}
	queue := []Pos{pos}

	head := 0
	for head < len(queue) {
		curr := queue[head]
		head++

		if grid[curr.x][curr.y] == 9 {
			count++
			continue
		}

		for _, dir := range allDirs {
			nextPos := curr.Move(dir)
			if !nextPos.InBounds() {
				continue
			}

			if _, ok := visited[nextPos]; ok && !repeat {
				continue
			}

			if grid[nextPos.x][nextPos.y]-grid[curr.x][curr.y] != 1 {
				continue
			}

			visited[nextPos] = struct{}{}
			queue = append(queue, nextPos)
		}
	}

	return count
}

func solve1(grid [][]int) string {
	result := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				score := walk(grid, Pos{x: i, y: j}, false)
				result += score
			}
		}
	}

	return fmt.Sprintf("%d", result)
}

func solve2(grid [][]int) string {
	result := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				rating := walk(grid, Pos{x: i, y: j}, true)
				result += rating
			}
		}
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	grid := make([][]int, len(input))

	rows = len(input)
	cols = len(input[0])

	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}

		grid[i] = make([]int, len(line))
		for j, char := range line {
			grid[i][j] = int(char - '0')
		}
	}

	startTime := time.Now()
	util.Output(1, solve1(grid))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(grid))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
