package parser

import (
	"bufio"
	"os"
)

// ReadFileLines reads a file and returns a string slice of all the lines
func ReadFileLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// ReadFile reads a file and returns a single string of the contents
func ReadFile(filename string) (string, error) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}
