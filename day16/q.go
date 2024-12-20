package day16

import (
	"fmt"
	"math"
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

func (d *Dir) RotateCW() Dir {
	return Dir{dx: d.dy, dy: -d.dx}
}

func (d *Dir) RotateCCW() Dir {
	return Dir{dx: -d.dy, dy: d.dx}
}

type State struct {
	Pos
	Dir
}

type Grid struct {
	start, end Pos
	walls      map[Pos]struct{}
}

const (
	STRAIGHT_COST = 1
	ROTATE_COST   = 1000
)

var (
	north = Dir{dx: -1, dy: 0}
	east  = Dir{dx: 0, dy: 1}
	south = Dir{dx: 1, dy: 0}
	west  = Dir{dx: 0, dy: -1}

	dirs = []Dir{north, east, south, west}
)

func solve1(grid *Grid, startDir Dir) string {
	distance := map[State]int{{Pos: grid.start, Dir: startDir}: 0}
	queue := make([]State, 0, 1e6)
	queue = append(queue, State{Pos: grid.start, Dir: startDir})

	for index := 0; index < len(queue); index++ {
		curr := queue[index]
		currCost := distance[curr]

		if curr.Pos == grid.end {
			continue
		}

		for _, dir := range []Dir{curr.Dir, curr.Dir.RotateCW(), curr.Dir.RotateCCW()} {
			nextPos := curr.Move(dir)
			if _, ok := grid.walls[nextPos]; ok {
				continue
			}

			nextState := State{Pos: nextPos, Dir: dir}
			newCost := currCost + STRAIGHT_COST

			if dir != curr.Dir {
				nextState.Pos = curr.Pos
				newCost += ROTATE_COST - STRAIGHT_COST
			}

			if cost, ok := distance[nextState]; !ok || newCost < cost {
				distance[nextState] = newCost
				queue = append(queue, nextState)
			}
		}
	}

	minCost := math.MaxInt
	for _, dir := range dirs {
		state := State{Pos: grid.end, Dir: dir}
		if cost, ok := distance[state]; ok && cost < minCost {
			minCost = cost
		}
	}

	return fmt.Sprintf("%d", minCost)
}

func solve2(grid *Grid, startDir Dir) string {
	distance := map[State]int{{Pos: grid.start, Dir: startDir}: 0}
	queue := make([]State, 0, 1e6)
	queue = append(queue, State{Pos: grid.start, Dir: startDir})

	parents := make(map[State][]State)
	paths := make(map[int][]State)

	for index := 0; index < len(queue); index++ {
		curr := queue[index]
		currCost := distance[curr]

		if curr.Pos == grid.end {
			paths[currCost] = append(paths[currCost], curr)
			continue
		}

		for _, dir := range []Dir{curr.Dir, curr.Dir.RotateCW(), curr.Dir.RotateCCW()} {
			nextPos := curr.Move(dir)
			if _, ok := grid.walls[nextPos]; ok {
				continue
			}

			nextState := State{Pos: nextPos, Dir: dir}
			newCost := currCost + STRAIGHT_COST

			if dir != curr.Dir {
				nextState.Pos = curr.Pos
				newCost += ROTATE_COST - STRAIGHT_COST
			}

			if cost, ok := distance[nextState]; !ok || newCost < cost {
				distance[nextState] = newCost
				parents[nextState] = []State{curr}
				queue = append(queue, nextState)
			} else if newCost == cost {
				parents[nextState] = append(parents[nextState], curr)
			}
		}
	}

	minCost := math.MaxInt
	endStates := make([]State, 0)
	for _, dir := range dirs {
		state := State{Pos: grid.end, Dir: dir}
		if cost, ok := distance[state]; ok && cost < minCost {
			minCost = cost
			endStates = []State{state}
		} else if cost == minCost {
			endStates = append(endStates, state)
		}
	}

	points := make(map[Pos]struct{})

	queue = make([]State, 0, len(parents))
	queue = append(queue, endStates...)

	visited := make(map[State]struct{})

	for index := 0; index < len(queue); index++ {
		curr := queue[index]

		if _, ok := points[curr.Pos]; !ok {
			points[curr.Pos] = struct{}{}
		}

		for _, parent := range parents[curr] {
			if _, ok := visited[parent]; ok {
				continue
			}

			queue = append(queue, parent)
			visited[parent] = struct{}{}
		}
	}

	return fmt.Sprintf("%d", len(points))
}

func Run(day int, input []string) {
	grid := &Grid{walls: make(map[Pos]struct{})}
	startDir := east

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
	util.Output(1, solve1(grid, startDir))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(grid, startDir))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
