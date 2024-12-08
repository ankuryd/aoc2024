package day08

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

// Delta : delta vector from p to other
func (p Pos) Delta(other Pos) Pos {
	return Pos{x: other.x - p.x, y: other.y - p.y}
}

func (p Pos) Add(other Pos) Pos {
	return Pos{x: p.x + other.x, y: p.y + other.y}
}

func (p Pos) Sub(other Pos) Pos {
	return Pos{x: p.x - other.x, y: p.y - other.y}
}

var (
	rows, cols int
)

func solve1(antennas map[rune][]Pos) {
	antiNodes := make(map[Pos]struct{})
	var antiNode Pos
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				delta := positions[i].Delta(positions[j])

				antiNode = positions[i].Sub(delta)
				if antiNode.InBounds() {
					antiNodes[antiNode] = struct{}{}
				}

				antiNode = positions[j].Add(delta)
				if antiNode.InBounds() {
					antiNodes[antiNode] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(antiNodes))
}

func solve2(antennas map[rune][]Pos) {
	antiNodes := make(map[Pos]struct{})
	var antiNode Pos
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				delta := positions[i].Delta(positions[j])

				antiNode = positions[i].Sub(delta)
				for {
					if !antiNode.InBounds() {
						break
					}

					antiNodes[antiNode] = struct{}{}
					antiNode = antiNode.Sub(delta)
				}

				antiNode = positions[j].Add(delta)
				for {
					if !antiNode.InBounds() {
						break
					}

					antiNodes[antiNode] = struct{}{}
					antiNode = antiNode.Add(delta)
				}
			}
		}
	}

	for _, positions := range antennas {
		for _, position := range positions {
			antiNodes[position] = struct{}{}
		}
	}

	fmt.Println(len(antiNodes))
}

func Run(day int, input []string) {
	antennas := make(map[rune][]Pos)

	rows = len(input)
	cols = len(input[0])

	for i, line := range input {
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", i)
		}

		for j, char := range line {
			if char == '.' {
				continue
			}

			if _, ok := antennas[char]; !ok {
				antennas[char] = make([]Pos, 0)
			}
			antennas[char] = append(antennas[char], Pos{x: i, y: j})
		}
	}

	fmt.Println("Question 1 output:")
	solve1(antennas)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(antennas)
}
