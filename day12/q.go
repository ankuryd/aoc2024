package day12

import (
	"fmt"
	"time"

	"aoc2024/util"
)

type Pos struct {
	x, y int
}

func (k *Pos) String() string {
	return fmt.Sprintf("(%d, %d)", k.x, k.y)
}

func (k *Pos) InBounds() bool {
	return k.x >= 0 && k.x < rows && k.y >= 0 && k.y < cols
}

type Cell struct {
	Pos
	val             rune
	area, perimeter int
	corners         int
}

func (c *Cell) String() string {
	return fmt.Sprintf("(Pos: %v, Val: %c, Area: %d, Perimeter: %d, Corners: %d)", c.Pos, c.val, c.area, c.perimeter, c.corners)
}

func NewCell(x, y int, val rune) *Cell {
	return &Cell{Pos: Pos{x: x, y: y}, val: val, area: 1, perimeter: 4, corners: 0}
}

type Dir struct {
	dx, dy int
}

func (c *Cell) Move(dir Dir) Pos {
	return Pos{x: c.x + dir.dx, y: c.y + dir.dy}
}

func (c *Cell) Adjacent(dir Dir) *Cell {
	next := c.Move(dir)
	if !next.InBounds() {
		return nil
	}

	nextCell := lookup[next]
	if nextCell.val == c.val {
		return nextCell
	}

	return nil
}

func (c *Cell) Corners() int {
	neighbors := map[string]*Cell{
		"N":  c.Adjacent(north),
		"E":  c.Adjacent(east),
		"S":  c.Adjacent(south),
		"W":  c.Adjacent(west),
		"NE": c.Adjacent(northEast),
		"SE": c.Adjacent(southEast),
		"SW": c.Adjacent(southWest),
		"NW": c.Adjacent(northWest),
	}

	corners := 0
	switch c.perimeter {
	case 0, 1:
		if isConcave(neighbors["N"], neighbors["E"], neighbors["NE"]) {
			corners++
		}
		if isConcave(neighbors["S"], neighbors["E"], neighbors["SE"]) {
			corners++
		}
		if isConcave(neighbors["S"], neighbors["W"], neighbors["SW"]) {
			corners++
		}
		if isConcave(neighbors["N"], neighbors["W"], neighbors["NW"]) {
			corners++
		}
	case 2:
		if (neighbors["N"] != nil && neighbors["E"] != nil) ||
			(neighbors["S"] != nil && neighbors["E"] != nil) ||
			(neighbors["S"] != nil && neighbors["W"] != nil) ||
			(neighbors["N"] != nil && neighbors["W"] != nil) {
			corners = 1
		}
		if isConcave(neighbors["N"], neighbors["E"], neighbors["NE"]) {
			corners++
		}
		if isConcave(neighbors["S"], neighbors["E"], neighbors["SE"]) {
			corners++
		}
		if isConcave(neighbors["S"], neighbors["W"], neighbors["SW"]) {
			corners++
		}
		if isConcave(neighbors["N"], neighbors["W"], neighbors["NW"]) {
			corners++
		}
	case 3:
		corners = 2
	case 4:
		corners = 4
	default:
		util.Fatal("Unexpected perimeter length: %d", c.perimeter)
	}

	return corners
}

type Component struct {
	cells           map[Pos]struct{}
	area, perimeter int
	corners         int
}

func (c *Component) String() string {
	cells := make([]string, 0, len(c.cells))
	for cell := range c.cells {
		cells = append(cells, cell.String())
	}

	return fmt.Sprintf("(Area: %d, Perimeter: %d, Corners: %d, Cells: %v)", c.area, c.perimeter, c.corners, cells)
}

func (c *Component) Add(cell *Cell) {
	c.cells[cell.Pos] = struct{}{}
	c.area += cell.area
	c.perimeter += cell.perimeter
	c.corners += cell.corners
}

func (c *Component) Has(cell *Cell) bool {
	_, ok := c.cells[cell.Pos]
	return ok
}

var (
	rows, cols int

	north = Dir{dx: -1, dy: 0}
	east  = Dir{dx: 0, dy: 1}
	south = Dir{dx: 1, dy: 0}
	west  = Dir{dx: 0, dy: -1}

	northEast = Dir{dx: -1, dy: 1}
	southEast = Dir{dx: 1, dy: 1}
	southWest = Dir{dx: 1, dy: -1}
	northWest = Dir{dx: -1, dy: -1}

	dirs = []Dir{north, east, south, west}

	lookup = make(map[Pos]*Cell)
)

func walk(cell *Cell, grid [][]*Cell) *Component {
	queue := make([]*Cell, 0)
	queue = append(queue, cell)

	component := &Component{cells: make(map[Pos]struct{}), area: 0, perimeter: 0, corners: 0}

	index := 0
	for index < len(queue) {
		curr := queue[index]
		index++

		if component.Has(curr) {
			continue
		}

		component.Add(curr)

		for _, dir := range dirs {
			nextCell := curr.Adjacent(dir)
			if nextCell != nil {
				queue = append(queue, nextCell)
			}
		}
	}

	return component
}

func solve1(grid [][]*Cell) string {
	components := make([]*Component, 0)
	for _, row := range grid {
		for _, cell := range row {
			found := false
			for _, component := range components {
				if component.Has(cell) {
					found = true
					break
				}
			}

			if found {
				continue
			}

			component := walk(cell, grid)
			components = append(components, component)
		}
	}

	result := 0
	for _, component := range components {
		result += component.area * component.perimeter
	}

	return fmt.Sprintf("%d", result)
}

func solve2(grid [][]*Cell) string {
	components := make([]*Component, 0)
	for _, row := range grid {
		for _, cell := range row {
			found := false
			for _, component := range components {
				if component.Has(cell) {
					found = true
					break
				}
			}

			if found {
				continue
			}

			component := walk(cell, grid)
			components = append(components, component)
		}
	}

	result := 0
	for _, component := range components {
		result += component.area * component.corners
	}

	return fmt.Sprintf("%d", result)
}

func isConcave(a, b, ab *Cell) bool {
	return a != nil && b != nil && ab == nil
}

func Run(day int, input []string) {
	rows, cols = len(input), len(input[0])

	grid := make([][]*Cell, len(input))
	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}

		grid[i] = make([]*Cell, len(line))
		for j, val := range line {
			grid[i][j] = NewCell(i, j, val)
			lookup[grid[i][j].Pos] = grid[i][j]
		}
	}

	for _, row := range grid {
		for _, cell := range row {
			curr := cell
			for _, dir := range dirs {
				if next := curr.Adjacent(dir); next != nil {
					curr.perimeter--
				}
			}

			curr.corners = curr.Corners()
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
