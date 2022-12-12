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

var cyclemap = map[int]bool{
	20:  true,
	60:  true,
	100: true,
	140: true,
	180: true,
	220: true,
}

func part1(data []string) int {

	var result int
	var x int = 1
	var cycles int = 1
	var sum int

	for _, thisLine := range data {

		command := strings.Fields(thisLine)

		if command[0] == "noop" {
			cycles++

			if cyclemap[cycles] {
				sum += cycles * x
				fmt.Println(cycles, x)
			}

		} else {
			cycles++
			if cyclemap[cycles] {
				sum += cycles * x
				fmt.Println(cycles, x)
			}
			cycles++
			val, _ := strconv.Atoi(command[1])
			x += val
			if cyclemap[cycles] {
				sum += cycles * x
				fmt.Println(cycles, x)
			}
		}

	}

	result = sum

	return result

}

func part2(data []string) int {

	var result int

	return result

}
