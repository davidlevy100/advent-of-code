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
	var grid [1000][1000]bool

	for _, thisLine := range data {
		command(thisLine, &grid)
	}

	result = count(&grid)

	return result

}

func part2(data []string) int {

	var result int
	var grid [1000][1000]int

	for _, thisLine := range data {
		command2(thisLine, &grid)
	}

	result = count2(&grid)

	return result

}

func command(s string, grid *[1000][1000]bool) {

	thisSlice := strings.Fields(s)

	switch {
	case thisSlice[0] == "toggle":
		x1, y1 := getCoords(thisSlice[1])
		x2, y2 := getCoords(thisSlice[3])
		apply(grid, thisSlice[0], x1, y1, x2, y2)
	case thisSlice[1] == "on":
		x1, y1 := getCoords(thisSlice[2])
		x2, y2 := getCoords(thisSlice[4])
		apply(grid, thisSlice[1], x1, y1, x2, y2)
	case thisSlice[1] == "off":
		x1, y1 := getCoords(thisSlice[2])
		x2, y2 := getCoords(thisSlice[4])
		apply(grid, thisSlice[1], x1, y1, x2, y2)
	}

}

func apply(grid *[1000][1000]bool, command string, x1, y1, x2, y2 int) {

	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			switch command {
			case "on":
				grid[y][x] = true
			case "off":
				grid[y][x] = false
			case "toggle":
				grid[y][x] = !grid[y][x]
			}
		}
	}

}

func getCoords(s string) (int, int) {
	data := strings.Split(s, ",")

	x, _ := strconv.Atoi(data[0])
	y, _ := strconv.Atoi(data[1])

	return x, y

}

func count(grid *[1000][1000]bool) int {

	var result int

	for _, thisRow := range grid {
		for _, thisColumn := range thisRow {
			if thisColumn {
				result++
			}
		}
	}

	return result

}

func command2(s string, grid *[1000][1000]int) {

	thisSlice := strings.Fields(s)

	switch {
	case thisSlice[0] == "toggle":
		x1, y1 := getCoords(thisSlice[1])
		x2, y2 := getCoords(thisSlice[3])
		apply2(grid, thisSlice[0], x1, y1, x2, y2)
	case thisSlice[1] == "on":
		x1, y1 := getCoords(thisSlice[2])
		x2, y2 := getCoords(thisSlice[4])
		apply2(grid, thisSlice[1], x1, y1, x2, y2)
	case thisSlice[1] == "off":
		x1, y1 := getCoords(thisSlice[2])
		x2, y2 := getCoords(thisSlice[4])
		apply2(grid, thisSlice[1], x1, y1, x2, y2)
	}

}

func apply2(grid *[1000][1000]int, command string, x1, y1, x2, y2 int) {

	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			switch command {
			case "on":
				grid[y][x]++
			case "off":
				grid[y][x]--
				if grid[y][x] < 0 {
					grid[y][x] = 0
				}
			case "toggle":
				grid[y][x] += 2
			}
		}
	}
}

func count2(grid *[1000][1000]int) int {

	var result int

	for _, thisRow := range grid {
		for _, thisColumn := range thisRow {
			result += thisColumn
		}
	}

	return result

}
