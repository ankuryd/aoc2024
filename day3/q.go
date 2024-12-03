package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	re = regexp.MustCompile(`mul\((\d+),(\d+)\)`) // mul(x,y)

	nre = regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\).*?do\(\)`) // mul(x,y)|don't()...do()
)

const (
	DONT = "don't"
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

func Run() {
	const filename = "day3/input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file '%s': %v", filename, err)
	}
	defer file.Close()

	inputs := make([]string, 0)
	lineNumber := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			log.Fatalf("Invalid format on line %d: empty line", lineNumber)
		}

		inputs = append(inputs, line)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file '%s': %v", filename, err)
	}

	input := strings.Join(inputs, "")

	fmt.Println("Question 1 output:")
	solve1(input)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(input)
}
