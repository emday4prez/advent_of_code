package three

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/emday4prez/advent_of_code/utils"
)

func Solve1() {

	input, err := utils.ReadInputToString("three/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	// // Define the regular expression to match strings
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	total := 0
	// // Find all matches
	matches := re.FindAllString(input, -1)

	for _, match := range matches {
		digits := re.FindStringSubmatch(match)
		if len(digits) > 0 {
			// Convert string matches to integers
			num1, err1 := strconv.Atoi(digits[1])
			num2, err2 := strconv.Atoi(digits[2])

			if err1 == nil && err2 == nil {
				total += num1 * num2

			}
		} else {
			fmt.Printf("Input: %s\nNo match\n\n", match)
		}
	}
	fmt.Println(total)
}
