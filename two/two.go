package two

import (
	"fmt"
	"strings"

	"github.com/emday4prez/advent_of_code/utils"
)

func SolvePart1() int {
	lines, err := utils.ReadInput("two/intput.txt")
	if err != nil {
		fmt.Println(err)
		return 1
	}
	total := 0
	for i := 0; i < len(lines); i++ {
		report := strings.Split(lines[i], " ")

		if isLevelSafe(report) {
			total++
		}

	}
	fmt.Println(total)
	return total
}

func isLevelSafe(levels []string) bool {

}
