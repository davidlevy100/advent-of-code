package main

import (
	"fmt"
	"strconv"
	"unicode"

	util "github.com/davidlevy100/advent-of-code/util"
)

type coords struct {
	y, x int
}

type gears struct {
	adjacent, nonadjacent []int
	gearCoords            map[coords][]int
}

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func part1(data []string) int {

	var result int

	gearVals := parseData(data)

	for _, val := range gearVals.adjacent {
		result += val
	}

	return result

}

func part2(data []string) int {

	var result int

	return result

}

func parseData(data []string) gears {

	var result gears

	var maxY = len(data)
	var maxX = len(data[0])

	var checkAdjacent = func(y, x int) bool {

		var isAdjacent bool

		for i := y - 1; i <= y+1; i++ {

			if i < 0 || i >= maxY {
				continue
			}

			for j := x - 1; j <= x+1; j++ {
				if j < 0 || j >= maxX {
					continue
				}

				if data[i][j] == '.' || unicode.IsDigit(rune(data[i][j])) {
					continue
				} else {
					isAdjacent = true
					if data[i][j] == '*' {
						result.gearCoords[coords{i, j}] = append(result.gearCoords[coords{i, j}])
					}
				}
			}
		}

		return isAdjacent

	}

	var word string
	var adjacent bool

	for y := 0; y < maxY; y++ {

		word = ""
		adjacent = false

		for x := 0; x < maxX; x++ {
			if !unicode.IsDigit(rune(data[y][x])) {
				if word != "" {
					if value, err := strconv.Atoi(word); err == nil {
						if adjacent {
							result.adjacent = append(result.adjacent, value)
						} else {
							result.nonadjacent = append(result.nonadjacent, value)
						}
					}
				}
				word = ""
				adjacent = false
				continue
			} else {
				word += string(data[y][x])
				if !adjacent {
					adjacent = checkAdjacent(y, x)
				}
			}
		}

		if word != "" {
			if value, err := strconv.Atoi(word); err == nil {
				if adjacent {
					result.adjacent = append(result.adjacent, value)
				} else {
					result.nonadjacent = append(result.nonadjacent, value)
				}
			}
		}
	}

	return result

}
