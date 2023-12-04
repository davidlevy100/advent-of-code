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

	screen := makeScreen()

	printScreen(screen)

	return result

}

func makeScreen() [][]string {
	var result [][]string = make([][]string, 6)

	for i := range result {
		result[i] = make([]string, 40)
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			result[i][j] = "."
		}
	}

	return result
}

func printScreen(screen [][]string) {
	for _, thisLine := range screen {
		fmt.Println(thisLine)
	}
}
