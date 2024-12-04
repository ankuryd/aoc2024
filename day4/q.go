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
				// East
				if i+1 < len(inputs) && inputs[i+1][j] == 'M' &&
					i+2 < len(inputs) && inputs[i+2][j] == 'A' &&
					i+3 < len(inputs) && inputs[i+3][j] == 'S' {
					result++
				}

				// West
				if i-1 >= 0 && inputs[i-1][j] == 'M' &&
					i-2 >= 0 && inputs[i-2][j] == 'A' &&
					i-3 >= 0 && inputs[i-3][j] == 'S' {
					result++
				}

				// North
				if j-1 >= 0 && inputs[i][j-1] == 'M' &&
					j-2 >= 0 && inputs[i][j-2] == 'A' &&
					j-3 >= 0 && inputs[i][j-3] == 'S' {
					result++
				}

				// South
				if j+1 < len(inputs[i]) && inputs[i][j+1] == 'M' &&
					j+2 < len(inputs[i]) && inputs[i][j+2] == 'A' &&
					j+3 < len(inputs[i]) && inputs[i][j+3] == 'S' {
					result++
				}

				// South-East
				if i+1 < len(inputs) && j+1 < len(inputs[i+1]) && inputs[i+1][j+1] == 'M' &&
					i+2 < len(inputs) && j+2 < len(inputs[i+2]) && inputs[i+2][j+2] == 'A' &&
					i+3 < len(inputs) && j+3 < len(inputs[i+3]) && inputs[i+3][j+3] == 'S' {
					result++
				}

				// South-West
				if i+1 < len(inputs) && j-1 >= 0 && inputs[i+1][j-1] == 'M' &&
					i+2 < len(inputs) && j-2 >= 0 && inputs[i+2][j-2] == 'A' &&
					i+3 < len(inputs) && j-3 >= 0 && inputs[i+3][j-3] == 'S' {
					result++
				}

				// North-East
				if i-1 >= 0 && j+1 < len(inputs[i-1]) && inputs[i-1][j+1] == 'M' &&
					i-2 >= 0 && j+2 < len(inputs[i-2]) && inputs[i-2][j+2] == 'A' &&
					i-3 >= 0 && j+3 < len(inputs[i-3]) && inputs[i-3][j+3] == 'S' {
					result++
				}

				// North-West
				if i-1 >= 0 && j-1 >= 0 && inputs[i-1][j-1] == 'M' &&
					i-2 >= 0 && j-2 >= 0 && inputs[i-2][j-2] == 'A' &&
					i-3 >= 0 && j-3 >= 0 && inputs[i-3][j-3] == 'S' {
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
}
