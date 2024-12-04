package one

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/emday4prez/advent_of_code/utils"
)

func SolvePart1() int {
	lines, err := utils.ReadInput("one/input.txt")
	if err != nil {
		return 1
	}

	left := []string{}
	right := []string{}

	for _, line := range lines {
		inputs := strings.Split(line, "   ")
		left = append(left, inputs[0])
		right = append(right, inputs[1])

	}
	slices.Sort(left)
	slices.Sort(right)

	total := 0

	for i := range 1000 {
		a, err := strconv.Atoi(left[i])
		if err != nil {
			// ... handle error
			panic(err)
		}
		b, err := strconv.Atoi(right[i])
		if err != nil {
			// ... handle error
			panic(err)
		}

		if a > b {
			total += a - b
		} else {
			total += b - a
		}
	}
	fmt.Println(total)
	return int(total)
}

func SolvePart2() int {
	lines, err := utils.ReadInput("one/input.txt")
	if err != nil {
		return 1
	}

	left := []string{}
	right := []string{}

	for _, line := range lines {
		inputs := strings.Split(line, "   ")
		left = append(left, inputs[0])
		right = append(right, inputs[1])
	}

	keys := make(map[string]int)

	for _, entry := range left {
		keys[entry] = 0

		for _, rightEntry := range right {
			if entry == rightEntry {
				keys[entry]++
			}
		}
	}
	total := 0

	for k, v := range keys {
		i, err := strconv.Atoi(k)
		if err != nil {
			return 1
		}
		total += i * v
	}
	fmt.Println(total)
	return total
}
