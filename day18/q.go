package day18

import (
	"fmt"
	"strings"
	"time"

	"aoc2024/util"
)

type Pos struct {
	x, y int
}

func (p *Pos) InBounds() bool {
	return p.x >= 0 && p.x < rows && p.y >= 0 && p.y < cols
}

type Dir struct {
	dx, dy int
}

type State struct {
	Pos
	steps int
}

func (p *Pos) Move(d Dir) Pos {
	return Pos{x: p.x + d.dx, y: p.y + d.dy}
}

type Grid struct {
	start, end Pos
	walls      map[Pos]struct{}
}

func (g *Grid) AddWall(pos Pos) {
	g.walls[pos] = struct{}{}
}

func (g *Grid) CreateWalls(bytes []Pos, k int) {
	for i := 0; i < k; i++ {
		g.walls[bytes[i]] = struct{}{}
	}
}

var (
	north = Dir{dx: 0, dy: -1}
	east  = Dir{dx: 1, dy: 0}
	south = Dir{dx: 0, dy: 1}
	west  = Dir{dx: -1, dy: 0}

	dirs = []Dir{north, east, south, west}
)

const (
	rows, cols = 71, 71
	k          = 1024
)

func solve1(bytes []Pos) string {
	grid := &Grid{
		start: Pos{x: 0, y: 0},
		end:   Pos{x: cols - 1, y: rows - 1},
		walls: make(map[Pos]struct{}),
	}
	grid.CreateWalls(bytes, k)

	visited := map[Pos]struct{}{grid.start: {}}
	queue := make([]State, 0, 1e6)
	queue = append(queue, State{Pos: grid.start, steps: 1})

	result := 0
	for index := 0; index < len(queue); index++ {
		curr := queue[index]

		for _, dir := range dirs {
			next := curr.Move(dir)

			if !next.InBounds() {
				continue
			}

			if _, ok := grid.walls[next]; ok {
				continue
			}

			if _, ok := visited[next]; ok {
				continue
			}

			visited[next] = struct{}{}

			if next == grid.end {
				result = curr.steps
				break
			}

			queue = append(queue, State{Pos: next, steps: curr.steps + 1})
		}
	}

	return fmt.Sprintf("%d", result)
}

func (g *Grid) FindPath() []Pos {
	visited := map[Pos]struct{}{g.start: {}}
	queue := make([][]Pos, 0, 1e6)
	queue = append(queue, []Pos{g.start})

	for index := 0; index < len(queue); index++ {
		curr := queue[index]

		for _, dir := range dirs {
			next := curr[len(curr)-1].Move(dir)

			if !next.InBounds() {
				continue
			}

			if _, ok := g.walls[next]; ok {
				continue
			}

			if _, ok := visited[next]; ok {
				continue
			}

			visited[next] = struct{}{}

			if next == g.end {
				return append(curr, next)
			}

			temp := make([]Pos, len(curr))
			copy(temp, curr)
			queue = append(queue, append(temp, next))
		}
	}

	return nil
}

func solve2(bytes []Pos) string {
	grid := &Grid{
		start: Pos{x: 0, y: 0},
		end:   Pos{x: cols - 1, y: rows - 1},
		walls: make(map[Pos]struct{}),
	}

	result := Pos{}
	for i := 0; i < len(bytes); i++ {
		path := grid.FindPath()

		if path == nil {
			result = bytes[i-1]
			break
		}

		for {
			grid.AddWall(bytes[i])
			if util.Contains(path, bytes[i]) {
				break
			}

			i++
		}
	}

	return fmt.Sprintf("%d,%d", result.x, result.y)
}

func Run(day int, input []string) {
	bytes := make([]Pos, 0)

	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}

		parts := strings.Split(line, ",")
		intParts, err := util.ConvertToInts(parts)
		if err != nil {
			util.Fatal("Invalid format on line %d: %s", i, err)
		}

		bytes = append(bytes, Pos{x: intParts[0], y: intParts[1]})
	}

	startTime := time.Now()
	util.Output(1, solve1(bytes))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(bytes))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
