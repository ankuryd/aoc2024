package day07

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"aoc2024/util"
)

type Test struct {
	output int
	inputs []int
}

type Op int

const (
	Add Op = iota
	Mul
	Concat
)

func (t Test) isValid(validOps []Op, index int, result int) bool {
	if result > t.output {
		return false
	}

	if index == len(t.inputs)-1 {
		return result == t.output
	}

	for _, op := range validOps {
		index++
		switch op {
		case Add:
			newResult := result + t.inputs[index]
			if t.isValid(validOps, index, newResult) {
				return true
			}
		case Mul:
			newResult := result * t.inputs[index]
			if t.isValid(validOps, index, newResult) {
				return true
			}
		case Concat:
			digits := math.Floor(math.Log10(float64(t.inputs[index]))) + 1
			newResult := result*int(math.Pow(10, float64(digits))) + t.inputs[index]
			if t.isValid(validOps, index, newResult) {
				return true
			}
		}
		index--
	}

	return false
}

func solve1(tests []Test) {
	result := 0

	for _, test := range tests {
		if test.isValid([]Op{Add, Mul}, 0, test.inputs[0]) {
			result += test.output
		}
	}

	fmt.Println(result)
}

func solve2(tests []Test) {
	result := 0

	for _, test := range tests {
		if test.isValid([]Op{Add, Mul, Concat}, 0, test.inputs[0]) {
			result += test.output
		}
	}

	fmt.Println(result)
}

func Run(day int, input []string) {
	tests := make([]Test, 0)

	for lineNumber, line := range input {
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", lineNumber)
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Fatalf("Invalid format on line %d: expected 2 parts, got %d", lineNumber, len(parts))
		}

		output, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Invalid format on line %d: expected int, got %v", lineNumber, err)
		}

		inputs := strings.Fields(strings.TrimSpace(parts[1]))
		intInputs, err := util.ConvertToIntSlice(inputs)
		if err != nil {
			log.Fatalf("Invalid format on line %d: expected ints, got %v", lineNumber, err)
		}

		tests = append(tests, Test{output, intInputs})
	}

	fmt.Println("Question 1 output:")
	solve1(tests)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(tests)
}
