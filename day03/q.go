package day03

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	DONT = "don't"
)

var (
	re  = regexp.MustCompile(`mul\((\d+),(\d+)\)`)                    // mul(x,y)
	nre = regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\).*?do\(\)`) // mul(x,y)|don't()...do()
)

func solve1(input string) {
	matches := re.FindAllStringSubmatch(input, -1)
	if len(matches) == 0 {
		fmt.Println("no matches for mul(x,y)")
		return
	}

	result := 0
	for _, match := range matches {
		first, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatalf("Error converting '%s' to integer: %v", match[1], err)
		}

		second, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatalf("Error converting '%s' to integer: %v", match[2], err)
		}

		result += first * second
	}

	fmt.Println(result)
}

func solve2(input string) {
	matches := nre.FindAllStringSubmatch(input, -1)
	if len(matches) == 0 {
		fmt.Println("no matches for mul(x,y)|don't()...do()")
		return
	}

	result := 0
	for _, match := range matches {
		if strings.HasPrefix(match[0], DONT) {
			continue
		}

		first, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatalf("Error converting '%s' to integer: %v", match[1], err)
		}

		second, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatalf("Error converting '%s' to integer: %v", match[2], err)
		}

		result += first * second
	}

	fmt.Println(result)
}

func Run(day int, input []string) {
	for lineNumber, line := range input {
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", lineNumber)
		}
	}

	combined := strings.Join(input, "")

	fmt.Println("Question 1 output:")
	solve1(combined)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(combined)
}
