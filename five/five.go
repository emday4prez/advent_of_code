package five

import (
	"fmt"

	"github.com/emday4prez/advent_of_code/utils"
)

func Solve1() ([]utils.OrderRule, [][]int, error) {
	orderRules, pagesToUpdate, err := utils.ReadManual("five/input.txt")
	if err != nil {
		fmt.Println("errrrrr reading input")
	}

	buildPresenceMap(pagesToUpdate)

	return orderRules, pagesToUpdate, nil
}

// practicing go maps
func buildPresenceMap(slices [][]int) map[int]map[int]bool {
	presenceOfNumbers := make(map[int]map[int]bool)
	for sliceIndex, slice := range slices {
		for _, num := range slice {
			if presenceOfNumbers[num] == nil {
				presenceOfNumbers[num] = make(map[int]bool)
			}
			presenceOfNumbers[num][sliceIndex] = true
		}
	}

	return presenceOfNumbers
}

func buildPositionMap(slices [][]int) map[int]map[int][]int {
	positionOfNumbers := make(map[int]map[int][]int)
	for sliceIndex, slice := range slices {
		for position, number := range slice {
			if positionOfNumbers[number] == nil {
				positionOfNumbers[number] = make(map[int][]int)
			}
			positionOfNumbers[number][sliceIndex] = append(positionOfNumbers[number][sliceIndex], position)
		}

	}
	return positionOfNumbers
}
