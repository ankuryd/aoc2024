package day20

import (
	"fmt"
	"log"
)

func solve1() {}

func solve2() {}

func Run(day int, input []string) {
	log.Fatalf("Error: Day '%d' not implemented.", day)

	for lineNumber, line := range input {
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", lineNumber)
		}
	}

	fmt.Println("Question 1 output:")
	solve1()

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2()
}
