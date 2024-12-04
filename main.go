package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ankuryd/aoc2024/day1"
	"github.com/ankuryd/aoc2024/day2"
	"github.com/ankuryd/aoc2024/day3"
	"github.com/ankuryd/aoc2024/day4"

	"github.com/ankuryd/aoc2024/util"

	"github.com/joho/godotenv"
)

var (
	dayFuncs = map[int]func(day int){
		1: day1.Run,
		2: day2.Run,
		3: day3.Run,
		4: day4.Run,
	}
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Error: No arguments provided. Use -h or --help for more information.")
	}

	switch args[0] {
	case "-h", "--help":
		fmt.Println(`
Usage: go run main.go [options]

Options:
  -d <day>    Specify the day to run (1-25)
  -a          Run all days
		`)
		return
	case "-d":
		if len(args) < 2 {
			log.Fatal("Error: -d flag requires a day number.")
		}

		day, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Error: Invalid day '%s': %v", args[1], err)
		}

		if day < 1 || day > 25 {
			log.Fatalf("Error: Day '%d' is out of range (1-25).", day)
		}

		runFunc, ok := dayFuncs[day]
		if !ok {
			log.Fatalf("Error: Day '%d' not implemented.", day)
		}

		util.ValidateFile(day)
		runFunc(day)
	case "-a":
		for day := 1; day <= 25; day++ {
			runFunc, ok := dayFuncs[day]
			if ok {
				fmt.Printf("Running day %d\n", day)
				util.ValidateFile(day)
				runFunc(day)
				fmt.Println("================================")
			} else {
				fmt.Printf("Day %d is not implemented. Skipping.\n", day)
			}
		}
	default:
		log.Fatalf("Error: Invalid argument '%s'. Use -h or --help for more information.", args[0])
	}
}
