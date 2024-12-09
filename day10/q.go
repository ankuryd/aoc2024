package day10

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

func (p Pos) Move(d Dir) Pos {
	return Pos{x: p.x + d.dx, y: p.y + d.dy}
}

type Dir struct {
	dx, dy int
}

var (
	rows, cols int

	north = Dir{0, -1}
	east  = Dir{1, 0}
	south = Dir{0, 1}
	west  = Dir{-1, 0}

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

func solve1(grid [][]int) {
	result := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				score := walk(grid, Pos{x: i, y: j}, false)
				result += score
			}
		}
	}

	fmt.Println(result)
}

func solve2(grid [][]int) {
	result := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				rating := walk(grid, Pos{x: i, y: j}, true)
				result += rating
			}
		}
	}

	fmt.Println(result)
}

func Run(day int, input []string) {
	grid := make([][]int, len(input))

	rows = len(input)
	cols = len(input[0])

	for i, line := range input {
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", i)
		}

		grid[i] = make([]int, len(line))
		for j, char := range line {
			grid[i][j] = int(char - '0')
		}
	}

	fmt.Println("Question 1 output:")
	solve1(grid)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(grid)
}
