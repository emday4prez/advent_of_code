package four

import (
	"fmt"

	"github.com/emday4prez/advent_of_code/utils"
)

  type Direction struct {
       dy, dx int
   }
   directions := []Direction{
    {-1, 0},  // UP (decrease y)
    {1, 0},   // DOWN (increase y)
    {0, 1},   // RIGHT (increase x)
    {0, -1},  // LEFT (decrease x)
    {-1, 1},  // UP-RIGHT
    {-1, -1}, // UP-LEFT
    {1, 1},   // DOWN-RIGHT
    {1, -1},  // DOWN-LEFT
   }

func Solve1() {

	const target = "XMAS"
	const gridSize = 140
 count := 0

	grid, err := utils.CreateGridFromFile("four/input.txt")
	if err != nil {
		fmt.Println("error reading file")
	}

	// 1. Iterate each position
	for y := range grid {
		for x := range grid[y]{
// 2. For each position, check all directions
			for _, dir := range directions{
					// 3. If valid pattern found, increment
				if checkPattern(grid, x,y, dir){
					count++
				}
			}
		}
	}
 


}


// position valid function 
func isValidPos(grid Grid, y, x int) bool {
    return y >= 0 && y < len(grid) && 
           x >= 0 && x < len(grid[0])
}
//check pattern from position in direction

func checkPattern(grid Grid, x,y int, direction Direction) bool{

}