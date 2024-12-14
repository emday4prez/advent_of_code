package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"strings"
)

type Grid [][]rune
type OrderRule [2]int

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

func CreateGridFromFile(path string) (Grid, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rows Grid
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, []rune(scanner.Text()))
	}

	return rows, scanner.Err()
}

func ReadManual(path string) ([]OrderRule, [][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var part1 []OrderRule
	var part2 [][]int
	finishedPartOne := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			finishedPartOne = true
			continue
		}
		if finishedPartOne {
			pageList := strings.Split(line, ",")

			var pages []int

			for _, page := range pageList {

				p, err := strconv.Atoi(page)
				if err != nil {
					continue
				}
				pages = append(pages, p)

			}

			part2 = append(part2, pages)
		}
	}
	fmt.Println(part2)
	return part1, part2, nil
}
