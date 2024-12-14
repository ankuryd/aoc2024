package day14

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"aoc2024/util"
)

type Pos struct {
	x, y int
}

type Vel struct {
	dx, dy int
}

type Robot struct {
	Pos
	Vel
	quadrant int
}

func (r *Robot) Move(steps int) {
	r.Pos.x += r.Vel.dx * steps
	r.Pos.y += r.Vel.dy * steps

	r.Pos.x = (r.Pos.x%cols + cols) % cols
	r.Pos.y = (r.Pos.y%rows + rows) % rows
}

func (r *Robot) Bound() {
	midX, midY := cols/2, rows/2
	switch {
	case r.Pos.x < midX && r.Pos.y < midY:
		r.quadrant = 0
	case r.Pos.x > midX && r.Pos.y < midY:
		r.quadrant = 1
	case r.Pos.x < midX && r.Pos.y > midY:
		r.quadrant = 2
	case r.Pos.x > midX && r.Pos.y > midY:
		r.quadrant = 3
	default:
		r.quadrant = -1
	}
}

type State struct {
	robots      []Robot
	time        int
	safetyScore int
}

func (s *State) SafetyScore() int {
	quadrants := make([]int, 4)
	for _, robot := range s.robots {
		if robot.quadrant == -1 {
			continue
		}
		quadrants[robot.quadrant]++
	}

	s.safetyScore = 1
	for _, count := range quadrants {
		s.safetyScore *= count
	}

	return s.safetyScore
}

func (s *State) Render() {
	grid := make([][]string, rows)
	for i := range grid {
		grid[i] = make([]string, cols)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	for _, robot := range s.robots {
		grid[robot.Pos.y][robot.Pos.x] = "#"
	}

	for _, row := range grid {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
}

const (
	cols, rows = 101, 103

	seconds = 100
)

var (
	re = regexp.MustCompile(`p=(-?\d+),(-?\d+)\s*v=(-?\d+),(-?\d+)`)
)

func Copy(robots []*Robot) []Robot {
	copy := make([]Robot, 0, len(robots))
	for _, robot := range robots {
		copy = append(copy, *robot)
	}
	return copy
}

func solve1(robots []*Robot) string {
	for _, robot := range robots {
		robot.Move(seconds)
		robot.Bound()
	}

	state := State{robots: Copy(robots), time: seconds}
	result := state.SafetyScore()

	return fmt.Sprintf("%d", result)
}

func solve2(robots []*Robot) string {
	states := make([]State, 0)

	t := seconds
	states = append(states, State{robots: Copy(robots), time: t})
	for i := 0; i < cols*rows; i++ {
		for _, robot := range robots {
			robot.Move(1)
			robot.Bound()
		}
		t++
		states = append(states, State{robots: Copy(robots), time: t})
	}

	sort.Slice(states, func(i, j int) bool {
		return states[i].SafetyScore() < states[j].SafetyScore()
	})

	states[0].Render()

	return fmt.Sprintf("%d", states[0].time)
}

func Run(day int, input []string) {
	robots := make([]*Robot, 0)

	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}

		matches := re.FindAllStringSubmatch(line, -1)
		if len(matches[0]) != 5 {
			util.Fatal("Invalid format on line %d: %s", i, line)
		}

		robots = append(robots, &Robot{
			Pos: Pos{x: util.ConvertToInt(matches[0][1]), y: util.ConvertToInt(matches[0][2])},
			Vel: Vel{dx: util.ConvertToInt(matches[0][3]), dy: util.ConvertToInt(matches[0][4])},
		})
	}

	startTime := time.Now()
	util.Output(1, solve1(robots))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(robots))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
