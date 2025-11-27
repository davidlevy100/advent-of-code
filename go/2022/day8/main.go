package main

import (
	"fmt"

	util "github.com/davidlevy100/advent-of-code/util"
)

type visible struct{}

type cell struct {
	x, y int
}

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func part1(data []string) int {

	var result int

	treeMatrix := loadMatrix(data)

	var visibleMap = make(map[cell]visible)
	var high int

	rows := len(treeMatrix)
	cols := len(treeMatrix[0])

	//test visibility left - right & right - left
	for y := 0; y < rows; y++ {
		high = -1
		for x := 0; x < cols; x++ {
			if treeMatrix[y][x] > high {
				visibleMap[cell{y, x}] = visible{}
				high = treeMatrix[y][x]
			}
		}

		high = -1
		for x := cols - 1; x >= 0; x-- {
			if treeMatrix[y][x] > high {
				visibleMap[cell{y, x}] = visible{}
				high = treeMatrix[y][x]
			}
		}
	}

	//test visibility top - bottom & bottom - top
	for x := 0; x < cols; x++ {
		high = -1
		for y := 0; y < rows; y++ {
			if treeMatrix[y][x] > high {
				visibleMap[cell{y, x}] = visible{}
				high = treeMatrix[y][x]
			}
		}

		high = -1
		for y := rows - 1; y >= 0; y-- {
			if treeMatrix[y][x] > high {
				visibleMap[cell{y, x}] = visible{}
				high = treeMatrix[y][x]
			}
		}
	}

	result = len(visibleMap)

	return result

}

func part2(data []string) int {

	var result int
	treeMatrix := loadMatrix(data)

	rows := len(treeMatrix)
	cols := len(treeMatrix[0])

	var highScore int

	for y := 1; y < rows-1; y++ {
		for x := 1; x < cols-1; x++ {
			var thisScore int

			up, down, left, right := y, y, x, x

			for up > 0 {
				up--
				if treeMatrix[up][x] >= treeMatrix[y][x] {
					break
				}
			}

			for down < rows-1 {
				down++
				if treeMatrix[down][x] >= treeMatrix[y][x] {
					break
				}
			}

			for left > 0 {
				left--
				if treeMatrix[y][left] >= treeMatrix[y][x] {
					break
				}
			}

			for right < cols-1 {
				right++
				if treeMatrix[y][right] >= treeMatrix[y][x] {
					break
				}
			}

			thisScore = (y - up) * (down - y) * (x - left) * (right - x)

			if thisScore > highScore {
				highScore = thisScore
			}

		}
	}

	result = highScore
	return result

}

func loadMatrix(data []string) [][]int {

	var result = make([][]int, len(data))
	columns := len(data[0])

	for i := range result {
		result[i] = make([]int, columns)
	}

	for i, thisRow := range data {
		for j, thisRune := range thisRow {
			result[i][j] = int(thisRune - '0')
		}
	}

	return result

}
