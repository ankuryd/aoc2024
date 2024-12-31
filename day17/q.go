package day17

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"aoc2024/util"
)

type Instruction struct {
	opcode  int
	operand int
}

type Opcode int

const (
	adv Opcode = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

type Operand string

const (
	Literal Operand = "literal"
	Combo   Operand = "combo"
)

type Computer struct {
	registers    map[byte]int
	pc           int
	instructions []Instruction
	queue        []byte
}

func (c *Computer) Copy() *Computer {
	copy := &Computer{
		registers:    make(map[byte]int),
		pc:           c.pc,
		instructions: make([]Instruction, len(c.instructions)),
		queue:        make([]byte, len(c.queue)),
	}

	for k, v := range c.registers {
		copy.registers[k] = v
	}

	copy.instructions = append(copy.instructions, c.instructions...)
	copy.queue = append(copy.queue, c.queue...)

	return copy
}

func (c *Computer) GetOperand(operand int, optype Operand) int {
	if optype == Literal {
		return operand
	}

	value := -1
	switch operand {
	case 0, 1, 2, 3:
		value = operand
	case 4:
		value = c.registers['A']
	case 5:
		value = c.registers['B']
	case 6:
		value = c.registers['C']
	default:
		util.Fatal("Invalid operand: %d", operand)
	}

	return value
}

func (c *Computer) Eval(instruction Instruction) {
	switch Opcode(instruction.opcode) {
	case adv:
		value := c.GetOperand(instruction.operand, Combo)
		c.registers['A'] = int(c.registers['A'] / int(math.Pow(2, float64(value))))
	case bxl:
		value := c.GetOperand(instruction.operand, Literal)
		c.registers['B'] = c.registers['B'] ^ value
	case bst:
		value := c.GetOperand(instruction.operand, Combo)
		c.registers['B'] = value % 8
	case jnz:
		value := c.GetOperand(instruction.operand, Literal)
		if c.registers['A'] == 0 {
			c.pc += 2
		} else {
			c.pc = value
		}
		return
	case bxc:
		c.registers['B'] = c.registers['B'] ^ c.registers['C']
	case out:
		value := c.GetOperand(instruction.operand, Combo)
		value = value % 8
		c.queue = append(c.queue, strconv.Itoa(value)...)
		c.queue = append(c.queue, ',')
	case bdv:
		value := c.GetOperand(instruction.operand, Combo)
		c.registers['B'] = int(c.registers['A'] / int(math.Pow(2, float64(value))))
	case cdv:
		value := c.GetOperand(instruction.operand, Combo)
		c.registers['C'] = int(c.registers['A'] / int(math.Pow(2, float64(value))))
	}

	c.pc += 2
}

func (c *Computer) Run() {
	for c.pc/2 < len(c.instructions) {
		instruction := c.instructions[c.pc/2]
		c.Eval(instruction)
	}
}

func (c *Computer) Reset() {
	c.pc = 0
	c.queue = make([]byte, 0)
}

func (c *Computer) IsQuine(i int) bool {
	input := ""
	for _, instruction := range c.instructions {
		input += fmt.Sprintf("%d,%d,", instruction.opcode, instruction.operand)
	}

	input = input[2*i:]

	c.Reset()
	c.Run()

	output := string(c.queue)

	return input == output
}

const (
	RegisterA = "Register A: "
	RegisterB = "Register B: "
	RegisterC = "Register C: "
	Program   = "Program: "
)

func solve1(computer *Computer) string {
	computer.Run()

	result := string(computer.queue)
	result = strings.Trim(result, ",")

	return result
}

func solve2(computer *Computer) string {
	result := 0
	for i := 2*len(computer.instructions) - 1; i >= 0; i-- {
		result <<= 3
		for {
			computer.registers['A'] = result
			if computer.IsQuine(i) {
				break
			}

			result++
		}
	}

	return fmt.Sprintf("%d", result)
}

func Run(day int, input []string) {
	computer := &Computer{
		registers:    make(map[byte]int),
		pc:           0,
		instructions: make([]Instruction, 0),
		queue:        make([]byte, 0),
	}

	for i, line := range input {
		if line == "" {
			continue
		}

		switch {
		case strings.HasPrefix(line, RegisterA):
			computer.registers['A'] = util.ConvertToInt(strings.TrimPrefix(line, RegisterA))
		case strings.HasPrefix(line, RegisterB):
			computer.registers['B'] = util.ConvertToInt(strings.TrimPrefix(line, RegisterB))
		case strings.HasPrefix(line, RegisterC):
			computer.registers['C'] = util.ConvertToInt(strings.TrimPrefix(line, RegisterC))
		case strings.HasPrefix(line, Program):
			program := strings.TrimPrefix(line, Program)
			instructions, err := util.ConvertToInts(strings.Split(program, ","))
			if err != nil {
				util.Fatal("Invalid format on line %d: %s", i, line)
			}

			for i := 0; i < len(instructions); i += 2 {
				computer.instructions = append(computer.instructions, Instruction{opcode: instructions[i], operand: instructions[i+1]})
			}
		}
	}

	startTime := time.Now()
	util.Output(1, solve1(computer))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(computer))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
