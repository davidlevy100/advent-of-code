package main

import (
	"fmt"

	util "github.com/davidlevy100/advent-of-code/util"
)

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data[0]))
	fmt.Printf("Part 2 answer: %d\n", part2(data[0]))

}

func part1(data string) int {

	var result int

	for _, thisRune := range data {
		switch thisRune {
		case '(':
			result++
		case ')':
			result--
		}
	}

	return result

}

func part2(data string) int {

	var result int
	var floor int

	for index, thisRune := range data {
		switch thisRune {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			result = index + 1
			break
		}
	}

	return result

}
