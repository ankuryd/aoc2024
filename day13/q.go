package day13

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"aoc2024/util"
)

type Pos struct {
	x, y int
}

type Machine struct {
	ButtonA Pos
	ButtonB Pos
	Prize   Pos
}

const (
	ButtonA = "Button A: "
	ButtonB = "Button B: "
	Prize   = "Prize: "

	CostA = 3
	CostB = 1

	MarginError = 10000000000000
)

var (
	re = regexp.MustCompile(`X[+=](\d+),\s*Y[+=](\d+)`)
)

func solve1(machines []*Machine) string {
	result := 0
	for _, machine := range machines {
		denominator := machine.ButtonA.x*machine.ButtonB.y - machine.ButtonA.y*machine.ButtonB.x
		numerator_a := machine.ButtonA.x*machine.Prize.y - machine.ButtonA.y*machine.Prize.x
		numerator_b := machine.ButtonB.y*machine.Prize.x - machine.ButtonB.x*machine.Prize.y

		if numerator_a%denominator != 0 || numerator_b%denominator != 0 {
			continue
		}

		a := numerator_b / denominator
		b := numerator_a / denominator
		result += a*CostA + b*CostB
	}

	return fmt.Sprintf("%d", result)
}

func solve2(machines []*Machine) string {
	result := 0
	for _, machine := range machines {
		denominator := machine.ButtonA.x*machine.ButtonB.y - machine.ButtonA.y*machine.ButtonB.x
		numerator_a := machine.ButtonA.x*(machine.Prize.y+MarginError) - machine.ButtonA.y*(machine.Prize.x+MarginError)
		numerator_b := machine.ButtonB.y*(machine.Prize.x+MarginError) - machine.ButtonB.x*(machine.Prize.y+MarginError)

		if numerator_a%denominator != 0 || numerator_b%denominator != 0 {
			continue
		}

		a := numerator_b / denominator
		b := numerator_a / denominator
		result += a*CostA + b*CostB
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	machines := make([]*Machine, 0)

	var buttonA, buttonB, prize Pos
	for i, line := range input {
		if line == "" {
			machine := &Machine{ButtonA: buttonA, ButtonB: buttonB, Prize: prize}
			machines = append(machines, machine)
			continue
		}

		switch {
		case strings.HasPrefix(line, ButtonA):
			line = strings.TrimPrefix(line, ButtonA)
			matches := re.FindAllStringSubmatch(line, -1)
			if len(matches[0]) != 3 {
				util.Fatal("Invalid format on line %d: %s", i, line)
			}
			buttonA = Pos{x: util.ConvertToInt(matches[0][1]), y: util.ConvertToInt(matches[0][2])}
		case strings.HasPrefix(line, ButtonB):
			line = strings.TrimPrefix(line, ButtonB)
			matches := re.FindAllStringSubmatch(line, -1)
			if len(matches[0]) != 3 {
				util.Fatal("Invalid format on line %d: %s", i, line)
			}
			buttonB = Pos{x: util.ConvertToInt(matches[0][1]), y: util.ConvertToInt(matches[0][2])}
		case strings.HasPrefix(line, Prize):
			line = strings.TrimPrefix(line, Prize)
			matches := re.FindAllStringSubmatch(line, -1)
			if len(matches[0]) != 3 {
				util.Fatal("Invalid format on line %d: %s", i, line)
			}
			prize = Pos{x: util.ConvertToInt(matches[0][1]), y: util.ConvertToInt(matches[0][2])}
		}
	}

	// Add the last machine
	machine := &Machine{ButtonA: buttonA, ButtonB: buttonB, Prize: prize}
	machines = append(machines, machine)

	startTime := time.Now()
	util.Output(1, solve1(machines))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(machines))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
