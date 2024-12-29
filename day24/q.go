package day24

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"
	"time"

	"aoc2024/util"
)

type Op string

const (
	AND Op = "AND"
	OR  Op = "OR"
	XOR Op = "XOR"
)

type Gate struct {
	input1, input2 string
	operation      Op
	output         string
}

func (g *Gate) Eval(wires map[string]int) (int, bool) {
	if _, ok := wires[g.input1]; !ok {
		return 0, false
	}

	if _, ok := wires[g.input2]; !ok {
		return 0, false
	}

	if _, ok := wires[g.output]; ok {
		return 0, false
	}

	output := 0
	switch g.operation {
	case AND:
		output = wires[g.input1] & wires[g.input2]
	case OR:
		output = wires[g.input1] | wires[g.input2]
	case XOR:
		output = wires[g.input1] ^ wires[g.input2]
	}

	return output, true
}

const (
	bits = 45
)

var (
	re = regexp.MustCompile(`(\w+) (\w+) (\w+) -> (\w+)`)
)

func solve1(wires map[string]int, gates []Gate) string {
	active := make(map[string]int)
	for wire, val := range wires {
		active[wire] = val
	}

	for {
		updates := make(map[string]int)
		for _, gate := range gates {
			if output, ok := gate.Eval(active); ok {
				updates[gate.output] = output
			}
		}

		for wire, val := range updates {
			active[wire] = val
		}

		if len(updates) == 0 {
			break
		}
	}

	result := 0
	for wire, val := range active {
		if strings.HasPrefix(wire, "z") {
			index := util.ConvertToInt(wire[1:])
			result += int(math.Pow(2, float64(index))) * val
		}
	}

	return fmt.Sprintf("%d", result)
}

func lookup(gates []Gate, input1 string, op Op, input2 string) string {
	for _, gate := range gates {
		if gate.input1 == input1 && gate.operation == op && gate.input2 == input2 {
			return gate.output
		}

		if gate.input1 == input2 && gate.operation == op && gate.input2 == input1 {
			return gate.output
		}
	}

	return ""
}

func solve2(gates []Gate) string {
	var (
		c0 = "" // C

		swapped = make([]string, 0)
	)

	for i := 0; i < bits; i++ {
		bit := fmt.Sprintf("%02d", i)

		var (
			s = "" // x xor y
			S = "" // s xor c0
			c = "" // x and y
			z = "" // s and c0
			C = "" // c or z
		)

		s = lookup(gates, "x"+bit, XOR, "y"+bit)
		c = lookup(gates, "x"+bit, AND, "y"+bit)

		if c0 != "" {
			z = lookup(gates, s, AND, c0)

			if z == "" {
				s, c = c, s
				swapped = append(swapped, s, c)

				z = lookup(gates, s, AND, c0)
			}

			S = lookup(gates, s, XOR, c0)

			if strings.HasPrefix(s, "z") {
				s, S = S, s
				swapped = append(swapped, s, S)
			}

			if strings.HasPrefix(c, "z") {
				c, S = S, c
				swapped = append(swapped, c, S)
			}

			if strings.HasPrefix(z, "z") {
				z, S = S, z
				swapped = append(swapped, z, S)
			}

			C = lookup(gates, c, OR, z)
		}

		if strings.HasPrefix(C, "z") {
			if C != fmt.Sprintf("z%02d", bits) {
				C, S = S, C
				swapped = append(swapped, C, S)
			}
		}

		if c0 == "" {
			c0 = c
		} else {
			c0 = C
		}
	}

	slices.Sort(swapped)

	return strings.Join(swapped, ",")
}

func Run(day int, input []string) {
	wires := make(map[string]int)
	gates := make([]Gate, 0)

	readGates := false
	for i, line := range input {
		if line == "" {
			readGates = true
			continue
		}

		if readGates {
			parts := re.FindStringSubmatch(line)
			if len(parts) != 5 {
				util.Fatal("Invalid format on line %d: %s", i, line)
			}

			gates = append(gates, Gate{input1: parts[1], input2: parts[3], operation: Op(parts[2]), output: parts[4]})
		} else {
			parts := strings.Split(line, ": ")
			if len(parts) != 2 {
				util.Fatal("Invalid format on line %d: %s", i, line)
			}

			wires[parts[0]] = util.ConvertToInt(parts[1])
		}
	}

	startTime := time.Now()
	util.Output(1, solve1(wires, gates))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(gates))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
