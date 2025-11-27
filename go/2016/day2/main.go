package main

import (
	"fmt"

	util "github.com/davidlevy100/advent-of-code/util"
)

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %s\n", part1(data))
	fmt.Printf("Part 2 answer: %s\n", part2(data))

}

var keyPad = [3][3]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
}

var keyPad2 = [5][5]string{
	{".", ".", "1", ".", "."},
	{".", "2", "3", "4", "."},
	{"5", "6", "7", "8", "9"},
	{".", "A", "B", "C", "."},
	{".", ".", "D", ".", "."},
}

func part1(data []string) string {

	var result string

	x, y := 1, 1

	for _, thisLine := range data {
		for _, thisRune := range thisLine {
			switch thisRune {
			case 'U':
				y--
			case 'D':
				y++
			case 'L':
				x--
			case 'R':
				x++
			}

			if y < 0 {
				y = 0
			}

			if y > 2 {
				y = 2
			}

			if x < 0 {
				x = 0
			}

			if x > 2 {
				x = 2
			}

		}
		result += keyPad[y][x]
	}

	return result

}

func part2(data []string) string {

	var result string

	x, y := 0, 2

	for _, thisLine := range data {
		for _, thisRune := range thisLine {

			switch thisRune {
			case 'U':
				if y > 0 && keyPad2[y-1][x] != "." {
					y--
				}
			case 'D':
				if y < 4 && keyPad2[y+1][x] != "." {
					y++
				}
			case 'L':
				if x > 0 && keyPad2[y][x-1] != "." {
					x--
				}
			case 'R':
				if x < 4 && keyPad2[y][x+1] != "." {
					x++
				}
			}
		}
		result += keyPad2[y][x]
	}

	return result

}
