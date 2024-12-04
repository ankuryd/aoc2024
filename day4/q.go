package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func solve1(inputs []string) {
	result := 0
	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs[i]); j++ {
			if inputs[i][j] == 'X' {
				directions := []struct{ di, dj int }{
					{1, 0},   // East
					{-1, 0},  // West
					{0, -1},  // North
					{0, 1},   // South
					{1, 1},   // South-East
					{1, -1},  // South-West
					{-1, 1},  // North-East
					{-1, -1}, // North-West
				}

				pattern := []byte{'M', 'A', 'S'}
				for _, dir := range directions {
					matched := true
					for k, char := range pattern {
						ni, nj := i+(k+1)*dir.di, j+(k+1)*dir.dj
						if ni < 0 || ni >= len(inputs) || nj < 0 || nj >= len(inputs[ni]) {
							matched = false
							break
						}

						if inputs[ni][nj] != char {
							matched = false
							break
						}
					}

					if matched {
						result++
					}
				}
			}
		}
	}

	fmt.Println(result)
}

func solve2(inputs []string) {
	result := 0
	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs[i]); j++ {
			if inputs[i][j] == 'A' {
				directions := []struct{ di, dj int }{
					{1, 1},  // South-East and North-West
					{1, -1}, // South-West and North-East
				}

				pattern := []byte{'M', 'S'}
				matched := true
				for _, dir := range directions {
					ni1, nj1 := i+dir.di, j+dir.dj
					ni2, nj2 := i-dir.di, j-dir.dj

					if ni1 < 0 || ni1 >= len(inputs) || nj1 < 0 || nj1 >= len(inputs[ni1]) ||
						ni2 < 0 || ni2 >= len(inputs) || nj2 < 0 || nj2 >= len(inputs[ni2]) {
						matched = false
						break
					}

					if !((inputs[ni1][nj1] == pattern[0] && inputs[ni2][nj2] == pattern[1]) ||
						(inputs[ni1][nj1] == pattern[1] && inputs[ni2][nj2] == pattern[0])) {
						matched = false
						break
					}
				}

				if matched {
					result++
				}
			}
		}
	}

	fmt.Println(result)
}

func Run(day int) {
	filename := fmt.Sprintf("day%d/input.txt", day)

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

	fmt.Println("Question 1 output:")
	solve1(inputs)

	fmt.Println("--------------------------------")

	fmt.Println("Question 2 output:")
	solve2(inputs)
}
