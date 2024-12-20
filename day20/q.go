package day20

import (
	"fmt"
	"time"

	"aoc2024/util"
)

type Pos struct {
	x, y int
}

func (p *Pos) Add(other Pos) Pos {
	return Pos{x: p.x + other.x, y: p.y + other.y}
}

type Dir struct {
	dx, dy int
}

func (p *Pos) Move(d Dir) Pos {
	return Pos{x: p.x + d.dx, y: p.y + d.dy}
}

type State struct {
	Pos
	Dir
}

type Grid struct {
	start, end Pos
	walls      map[Pos]struct{}
}

var (
	north = Dir{dx: -1, dy: 0}
	east  = Dir{dx: 0, dy: 1}
	south = Dir{dx: 1, dy: 0}
	west  = Dir{dx: 0, dy: -1}

	dirs = []Dir{north, east, south, west}
)

func solve1(grid *Grid) string {
	distance := map[Pos]int{grid.start: 0}
	queue := make([]Pos, 0, 1e6)
	queue = append(queue, grid.start)

	for index := 0; index < len(queue); index++ {
		curr := queue[index]

		if curr == grid.end {
			break
		}

		for _, dir := range dirs {
			next := curr.Move(dir)

			if _, ok := grid.walls[next]; ok {
				continue
			}

			if _, ok := distance[next]; ok {
				continue
			}

			distance[next] = distance[curr] + 1
			queue = append(queue, next)
		}
	}

	relPos := []Pos{}
	for dx := -2; dx <= 2; dx++ {
		for dy := -2; dy <= 2; dy++ {
			if util.Abs(dx)+util.Abs(dy) <= 2 && !(dx == 0 && dy == 0) {
				relPos = append(relPos, Pos{dx, dy})
			}
		}
	}

	count := make(map[int]int)
	for p1, d1 := range distance {
		for _, rel := range relPos {
			p2 := p1.Add(rel)
			d2, ok := distance[p2]
			if !ok {
				continue
			}

			delta := util.Abs(d1-d2) - (util.Abs(rel.x) + util.Abs(rel.y))
			if delta >= 100 {
				count[delta]++
			}
		}
	}

	result := 0
	for delta, count := range count {
		if delta >= 100 {
			result += count / 2
		}
	}

	return fmt.Sprintf("%d", result)
}

func solve2(grid *Grid) string {
	distance := map[Pos]int{grid.start: 0}
	queue := make([]Pos, 0, 1e6)
	queue = append(queue, grid.start)

	for index := 0; index < len(queue); index++ {
		curr := queue[index]

		if curr == grid.end {
			break
		}

		for _, dir := range dirs {
			next := curr.Move(dir)

			if _, ok := grid.walls[next]; ok {
				continue
			}

			if _, ok := distance[next]; ok {
				continue
			}

			distance[next] = distance[curr] + 1
			queue = append(queue, next)
		}
	}

	relPos := []Pos{}
	for dx := -20; dx <= 20; dx++ {
		for dy := -20; dy <= 20; dy++ {
			if util.Abs(dx)+util.Abs(dy) <= 20 && !(dx == 0 && dy == 0) {
				relPos = append(relPos, Pos{dx, dy})
			}
		}
	}

	count := make(map[int]int)
	for p1, d1 := range distance {
		for _, rel := range relPos {
			p2 := p1.Add(rel)
			d2, ok := distance[p2]
			if !ok {
				continue
			}

			delta := util.Abs(d1-d2) - (util.Abs(rel.x) + util.Abs(rel.y))
			if delta >= 100 {
				count[delta]++
			}
		}
	}

	result := 0
	for delta, count := range count {
		if delta >= 100 {
			result += count / 2
		}
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	grid := &Grid{
		walls: make(map[Pos]struct{}),
	}

	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}

		for j, char := range line {
			switch char {
			case 'S':
				grid.start = Pos{x: i, y: j}
			case 'E':
				grid.end = Pos{x: i, y: j}
			case '#':
				grid.walls[Pos{x: i, y: j}] = struct{}{}
			}
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
