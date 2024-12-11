package utils

import (
	"bufio"
	"os"
)

func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func ReadInputToString(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func CreateGridFromFile(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}

	return lines, scanner.Err()
}
