package helpers

import (
	"bufio"
	"os"
)

// ReadLines reads a file and returns a slice of strings, one for each line of the file.
func ReadLines(path string) ([]string, error) {
	if path == "" {
		path = "./input.txt"
	}
	
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}