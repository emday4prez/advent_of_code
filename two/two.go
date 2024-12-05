package two

import (
	"fmt"

	"github.com/emday4prez/advent_of_code/utils"
)

func SolvePart1() int {
	lines, err := utils.ReadInput("two/intput.txt")
	if err != nil {
		fmt.Println(err)
		return 1
	}

	for _, levels := range lines {
		fmt.Println(levels)

	}

	return 0
}
