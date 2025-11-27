package main

import (
	"fmt"

	util "github.com/davidlevy100/advent-of-code/util"
)

var exists struct{}

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data[0]))
	fmt.Printf("Part 2 answer: %d\n", part2(data[0]))

}

func part1(data string) int {

	return findMarker(data, 4)

}

func part2(data string) int {

	return findMarker(data, 14)

}

func findMarker(data string, length int) int {

	var result int

	for i := length; i < len(data); i++ {

		var runeMap = make(map[byte]struct{})

		for j := i - length; j < i; j++ {
			runeMap[data[j]] = exists
		}

		if len(runeMap) == length {
			result = i
			break
		}

	}

	return result

}
