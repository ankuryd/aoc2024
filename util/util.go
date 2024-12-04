package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// ValidateFile checks if the input file for the given day exists, and if not, it downloads it
func ValidateFile(day int) {
	filename := fmt.Sprintf("day%d/input.txt", day)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("Downloading input for day %d\n", day)
		DownloadInput(day)
	}
}

// DownloadInput downloads the input for the given day and saves it to the input.txt file
func DownloadInput(day int) {
	filename := fmt.Sprintf("day%d/input.txt", day)

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day), nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Accept", "text/html")
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", os.Getenv("SESSION_COOKIE")))

	//////////
	fmt.Println(req.Header)
	//////////

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to download input: status code %d", resp.StatusCode)
	}

	out, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}
