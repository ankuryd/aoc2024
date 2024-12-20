package util

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// ProcessInput validates the input for the given day and returns the input as a slice of strings
func ProcessInput(day int, isTest bool) []string {
	ValidateInput(day)

	filename := fmt.Sprintf("day%02d/input.txt", day)
	if isTest {
		filename = fmt.Sprintf("day%02d/test.txt", day)
	}

	file, err := os.Open(filename)
	if err != nil {
		Fatal("Error opening file '%s': %v", filename, err)
	}
	defer file.Close()

	input := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		Fatal("Error reading file '%s': %v", filename, err)
	}

	return input
}

// ValidateInput checks if the input file for the given day exists, and if not, it downloads it
func ValidateInput(day int) {
	filename := fmt.Sprintf("day%02d/input.txt", day)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		Print("Downloading input for day %d\n", day)
		DownloadInput(day)
	}
}

// DownloadInput downloads the input for the given day and saves it to the input.txt file
func DownloadInput(day int) {
	filename := fmt.Sprintf("day%02d/input.txt", day)

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day), nil)
	if err != nil {
		Fatal("Error creating request: %v", err)
	}

	req.Header.Set("Accept", "text/html")
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", os.Getenv("SESSION_COOKIE")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		Fatal("Error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		Fatal("Failed to download input: status code %d", resp.StatusCode)
	}

	out, err := os.Create(filename)
	if err != nil {
		Fatal("Error creating file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		Fatal("Error writing to file: %v", err)
	}
}

// ConvertToInts converts a slice of strings to a slice of integers
func ConvertToInts(input []string) ([]int, error) {
	result := make([]int, len(input))
	for i, v := range input {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		result[i] = num
	}

	return result, nil
}

// ConvertToInt converts a string to an integer
func ConvertToInt(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		Fatal("Error converting to int: %v", err)
	}

	return num
}

// Contains checks if a slice contains an item
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}

	return false
}

type Number interface {
	constraints.Integer | constraints.Float
}

// Abs returns the absolute value of a number
func Abs[T Number](a T) T {
	if a < 0 {
		return -a
	}

	return a
}
