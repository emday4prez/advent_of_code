package five

import (
	"fmt"

	"github.com/emday4prez/advent_of_code/utils"
)

// extract ordering rules
// -scan until empty string

// split ordering rules by | and convert to int and [][]int
//
// continue with scanning, but now for pages-to-produce
// convert to [][]int

func Solve1() ([]utils.OrderRule, [][]int, error) {
	orderRules, pagesToUpdate, err := utils.ReadManual("five/input.txt")
	if err != nil {
		fmt.Println("errrrrr reading input")
	}
	return orderRules, pagesToUpdate, nil
}
