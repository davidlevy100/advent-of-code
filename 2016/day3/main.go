package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/davidlevy100/advent-of-code/util"
)

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func part1(data []string) int {

	var result int

	for _, thisLine := range data {

		inputs := strings.Fields(thisLine)

		if isTriangle(inputs[0], inputs[1], inputs[2]) {
			result++
		}
	}

	return result

}

func part2(data []string) int {

	var result int

	for i := 0; i < len(data)-2; i += 3 {

		var rows = make([][]string, 3)

		rows[0] = strings.Fields(data[i])
		rows[1] = strings.Fields(data[i+1])
		rows[2] = strings.Fields(data[i+2])

		transpose := [][]string{
			{rows[0][0], rows[1][0], rows[2][0]},
			{rows[0][1], rows[1][1], rows[2][1]},
			{rows[0][2], rows[1][2], rows[2][2]},
		}

		for _, inputs := range transpose {

			if isTriangle(inputs[0], inputs[1], inputs[2]) {
				result++
			}
		}
	}

	return result
}

func isTriangle(as, bs, cs string) bool {
	var result bool

	a, _ := strconv.Atoi(as)
	b, _ := strconv.Atoi(bs)
	c, _ := strconv.Atoi(cs)

	if a < b+c && b < a+c && c < a+b {
		result = true
	}

	return result
}
