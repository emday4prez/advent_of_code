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

	positions := buildPositionMap(pagesToUpdate)

	matches := make([]int, 0)
	seen := make(map[int]bool)
	for _, pair := range orderRules {

		first, second := pair[0], pair[1]
		// Get slices containing first number
		firstSlices := positions[first]
		if firstSlices == nil {
			continue
		}

		// Get slices containing second number
		secondSlices := positions[second]
		if secondSlices == nil {
			continue
		}

		for sliceIndex, firstPositions := range firstSlices {
			secondPositions, exists := secondSlices[sliceIndex]
			if !exists {
				continue
			}

			for _, firstPos := range firstPositions {
				for _, secondPos := range secondPositions {
					if firstPos < secondPos && !seen[sliceIndex] {
						matches = append(matches, sliceIndex)
						seen[sliceIndex] = true
						break
					}
				}
				if seen[sliceIndex] {
					break
				}
			}
		}

	}

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
