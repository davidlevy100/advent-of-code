package main

import (
	"fmt"
	"strconv"

	util "github.com/davidlevy100/advent-of-code/util"
)

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func part1(data []string) int {

	var result int

	for _, l := range data {
		l2, _ := strconv.Unquote(l)
		result += len(l) - len(l2)
	}

	return result

}

func part2(data []string) int {

	var result int

	for _, l := range data {
		l2 := strconv.Quote(l)
		result += len(l2) - len(l)
	}

	return result

}
