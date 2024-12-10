package three

import (
	"fmt"
	"regexp"
)

func Solve1() {

	longString := "This is a long string with some numbers 123 and another number 456."

	// Define the regular expression to match strings
	re := regexp.MustCompile(`[a-zA-Z]+`)

	// Find all matches
	matches := re.FindAllString(longString, -1)

	// Print the extracted strings
	for _, match := range matches {
		fmt.Println(match)
	}

}
