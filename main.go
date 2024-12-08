package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"aoc2024/day1"
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
	"aoc2024/day2"
	"aoc2024/day20"
	"aoc2024/day21"
	"aoc2024/day22"
	"aoc2024/day23"
	"aoc2024/day24"
	"aoc2024/day25"
	"aoc2024/day3"
	"aoc2024/day4"
	"aoc2024/day5"
	"aoc2024/day6"
	"aoc2024/day7"
	"aoc2024/day8"
	"aoc2024/day9"

	"aoc2024/util"

	"github.com/joho/godotenv"
)

var (
	dayFuncs = map[int]func(day int, input []string){
		1:  day1.Run,
		2:  day2.Run,
		3:  day3.Run,
		4:  day4.Run,
		5:  day5.Run,
		6:  day6.Run,
		7:  day7.Run,
		8:  day8.Run,
		9:  day9.Run,
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
		log.Fatalf("Error: Day '%d' not implemented.", day)
	}

	fmt.Printf("Running day %d\n", day)
	input := util.ProcessInput(day, isTest)
	runFunc(day, input)
	fmt.Println("================================")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Error: No arguments provided. Use -h or --help for more information.")
	}

	day := 0
	isTest := false
	runAll := false

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-h", "--help":
			fmt.Println(`
Usage: go run main.go [options]

Options:
  -d <day>    Specify the day to run (1-25)
  -a          Run all days
  -t          Run the test input for the given day
			`)
			return
		case "-d":
			if i+1 >= len(args) {
				log.Fatal("Error: -d flag requires a day number.")
			}

			day, err = strconv.Atoi(args[i+1])
			if err != nil {
				log.Fatalf("Error: Invalid day '%s': %v", args[i+1], err)
			}
			i++

			if day < 1 || day > 25 {
				log.Fatalf("Error: Day '%d' is out of range (1-25).", day)
			}
		case "-t":
			isTest = true
		case "-a":
			runAll = true
		default:
			log.Fatalf("Error: Invalid argument '%s'. Use -h or --help for more information.", args[i])
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
