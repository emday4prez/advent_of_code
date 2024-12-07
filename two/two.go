package two

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/emday4prez/advent_of_code/utils"
)

func SolvePart1() int {
	lines, err := utils.ReadInput("two/input.txt")
	if err != nil {
		fmt.Println(err)

	}
	total := countSafeReports(lines)
	fmt.Println(total)
	return total
}

func isValidSequence(report []int) bool {
	if len(report) < 2 {
		return true
	}
	isIncreasing := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		dif := report[i] - report[i-1]

		if isIncreasing && (dif <= 0 || dif > 3) {
			return false
		}
		if !isIncreasing && (dif >= 0 || dif < -3) {
			return false
		}
	}
	return true
}

func countSafeReports(reports []string) int {
	count := 0
	for _, report := range reports {
		// parse numbers
		var levels []int

		for _, strNum := range strings.Fields(report) {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				panic(err)
			}
			levels = append(levels, num)
		}
		if isValidSequence(levels) {
			count++

		}
	}
	return count
}
