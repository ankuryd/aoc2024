package day21

import (
	"fmt"
	"strings"
	"time"

	"aoc2024/util"
)

type Pos struct {
	x, y int
}

func (p *Pos) Delta(other Pos) Pos {
	return Pos{x: other.x - p.x, y: other.y - p.y}
}

func (p *Pos) IsInvalid(others []Pos) bool {
	for _, other := range others {
		if p.x == other.x && p.y == other.y {
			return true
		}
	}
	return false
}

type Keypad struct {
	keys    []string
	keypad  map[string]Pos
	invalid []Pos
}

func (k *Keypad) CreateGraph() map[string]map[string]string {
	graph := make(map[string]map[string]string)
	for _, a := range k.keys {
		for _, b := range k.keys {
			posA := k.keypad[a]
			posB := k.keypad[b]

			path := ""
			delta := posA.Delta(posB)

			if delta.y < 0 {
				path += strings.Repeat("<", -delta.y)
			}
			if delta.x > 0 {
				path += strings.Repeat("v", delta.x)
			}
			if delta.x < 0 {
				path += strings.Repeat("^", -delta.x)
			}
			if delta.y > 0 {
				path += strings.Repeat(">", delta.y)
			}

			corner1, corner2 := Pos{x: posA.x, y: posB.y}, Pos{x: posB.x, y: posA.y}
			if corner1.IsInvalid(k.invalid) || corner2.IsInvalid(k.invalid) {
				path = util.ReverseString(path)
			}

			if _, ok := graph[a]; !ok {
				graph[a] = make(map[string]string)
			}

			if _, ok := graph[a][b]; !ok {
				graph[a][b] = path + "A"
			}
		}
	}
	return graph
}

var (
	/*
		+---+---+---+
		| 7 | 8 | 9 |
		+---+---+---+
		| 4 | 5 | 6 |
		+---+---+---+
		| 1 | 2 | 3 |
		+---+---+---+
			| 0 | A |
			+---+---+
	*/
	numpad = Keypad{
		keys: []string{"7", "8", "9", "4", "5", "6", "1", "2", "3", "0", "A"},
		keypad: map[string]Pos{
			"7": {x: 0, y: 0},
			"8": {x: 0, y: 1},
			"9": {x: 0, y: 2},
			"4": {x: 1, y: 0},
			"5": {x: 1, y: 1},
			"6": {x: 1, y: 2},
			"1": {x: 2, y: 0},
			"2": {x: 2, y: 1},
			"3": {x: 2, y: 2},
			"0": {x: 3, y: 1},
			"A": {x: 3, y: 2},
		},
		invalid: []Pos{{x: 3, y: 0}},
	}

	/*
		    +---+---+
			| ^ | A |
		+---+---+---+
		| < | v | > |
		+---+---+---+
	*/
	dirpad = Keypad{
		keys: []string{"^", "A", "<", "v", ">"},
		keypad: map[string]Pos{
			"^": {x: 0, y: 1},
			"A": {x: 0, y: 2},
			"<": {x: 1, y: 0},
			"v": {x: 1, y: 1},
			">": {x: 1, y: 2},
		},
		invalid: []Pos{{x: 0, y: 0}},
	}
)

var (
	graphNumpad = numpad.CreateGraph()
	graphDirpad = dirpad.CreateGraph()
)

func expand(input string, graph map[string]map[string]string) string {
	prev := "A"
	output := ""
	for _, char := range input {
		output += graph[prev][string(char)]
		prev = string(char)
	}
	return output
}

func solve1(inputs []string) string {
	result := 0
	for _, input := range inputs {
		output := expand(input, graphNumpad)
		output = expand(output, graphDirpad)
		output = expand(output, graphDirpad)

		complexity := len(output) * util.ConvertToInt(input[:len(input)-1])
		result += complexity
	}

	return fmt.Sprintf("%d", result)
}

var (
	cache = make(map[string]map[int]int)
)

func compute(input string, iterations int) int {
	if iterations == 0 {
		return len(input)
	}

	if _, ok := cache[input]; !ok {
		cache[input] = make(map[int]int)
	}

	if _, ok := cache[input][iterations]; ok {
		return cache[input][iterations]
	}

	prev := "A"
	length := 0
	for _, char := range input {
		length += compute(graphDirpad[prev][string(char)], iterations-1)
		prev = string(char)
	}

	cache[input][iterations] = length
	return length
}

func solve2(inputs []string) string {
	result := 0
	for _, input := range inputs {
		output := expand(input, graphNumpad)
		length := compute(output, 25)

		complexity := length * util.ConvertToInt(input[:len(input)-1])
		result += complexity
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	inputs := make([]string, 0)

	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}

		inputs = append(inputs, line)
	}

	startTime := time.Now()
	util.Output(1, solve1(inputs))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(inputs))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
