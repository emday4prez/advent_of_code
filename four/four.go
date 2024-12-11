package four

import (
	"fmt"

	"github.com/emday4prez/advent_of_code/utils"
)

type Grid [][]rune

func Solve1() {

	const target = "XMAS"
	const gridSize = 140

	grid, err := utils.CreateGridFromFile("four/input.txt")
	if err != nil {
		fmt.Println("error reading file")
	}

	fmt.Println(grid)

}
