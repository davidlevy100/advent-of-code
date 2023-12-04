package main

import (
	"fmt"

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
		if abbaTest(thisLine) {
			result++
		}
	}

	return result

}

func part2(data []string) int {

	var result int

	return result

}
