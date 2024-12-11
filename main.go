package main

import (
	"os"
	"strconv"

	"aoc2024/day01"
	"aoc2024/day02"
	"aoc2024/day03"
	"aoc2024/day04"
	"aoc2024/day05"
	"aoc2024/day06"
	"aoc2024/day07"
	"aoc2024/day08"
	"aoc2024/day09"
	"aoc2024/day10"
	"aoc2024/day11"
	"aoc2024/day12"
	"aoc2024/day13"
	"aoc2024/day14"
	"aoc2024/day15"
	"aoc2024/day16"
	"aoc2024/day17"
	"aoc2024/day18"
	"aoc2024/day19"
	"aoc2024/day20"
	"aoc2024/day21"
	"aoc2024/day22"
	"aoc2024/day23"
	"aoc2024/day24"
	"aoc2024/day25"

	"aoc2024/util"

	"github.com/joho/godotenv"
)

var (
	dayFuncs = map[int]func(day int, input []string){
		1:  day01.Run,
		2:  day02.Run,
		3:  day03.Run,
		4:  day04.Run,
		5:  day05.Run,
		6:  day06.Run,
		7:  day07.Run,
		8:  day08.Run,
		9:  day09.Run,
		10: day10.Run,
		11: day11.Run,
		12: day12.Run,
		13: day13.Run,
		14: day14.Run,
		15: day15.Run,
		16: day16.Run,
		17: day17.Run,
		18: day18.Run,
		19: day19.Run,
		20: day20.Run,
		21: day21.Run,
		22: day22.Run,
		23: day23.Run,
		24: day24.Run,
		25: day25.Run,
	}
)

func run(day int, isTest bool) {
	runFunc, ok := dayFuncs[day]
	if !ok {
		util.Fatal("Error: Day '%d' not implemented.", day)
	}

	util.Print("Running day %d", day)
	input := util.ProcessInput(day, isTest)
	runFunc(day, input)
	util.MegaSeparator()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		util.Fatal("Error loading .env file: %v", err)
	}

	args := os.Args[1:]
	if len(args) == 0 {
		util.Fatal("Error: No arguments provided. Use -h or --help for more information.")
	}

	day := 0
	isTest := false
	runAll := false

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-h", "--help":
			util.Print(`
Usage: go run main.go [options]

Options:
  -d <day>    Specify the day to run (1-25)
  -a          Run all days
  -t          Run the test input for the given day
			`)
			return
		case "-d":
			if i+1 >= len(args) {
				util.Fatal("Error: -d flag requires a day number.")
			}

			day, err = strconv.Atoi(args[i+1])
			if err != nil {
				util.Fatal("Error: Invalid day '%s': %v", args[i+1], err)
			}
			i++

			if day < 1 || day > 25 {
				util.Fatal("Error: Day '%d' is out of range (1-25).", day)
			}
		case "-t":
			isTest = true
		case "-a":
			runAll = true
		default:
			util.Fatal("Error: Invalid argument '%s'. Use -h or --help for more information.", args[i])
		}
	}

	if runAll {
		for day := 1; day <= 25; day++ {
			run(day, isTest)
		}
	} else {
		run(day, isTest)
	}
}
