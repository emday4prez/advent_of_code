package two

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/emday4prez/advent_of_code/utils"
)

func SolvePart1() int {
	reports, err := utils.ReadInput("two/intput.txt")
	if err != nil {
		fmt.Println(err)
		return 1
	}

}

func isValidSequence(report []int) bool {

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
