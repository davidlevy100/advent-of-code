package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	util "github.com/davidlevy100/advent-of-code/util"
)

func main() {

	path, _ := os.Getwd()
	fullPath := filepath.Join(path, "input.txt")

	data, _ := util.GetInput(fullPath)

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func compile(data []string) []int {

	currElf := 0
	maxElfScore := 0
	currScore := 0

	scores := []int{}

	for _, thisLine := range data {

		if thisLine == "" {

			scores = append(scores, currScore)

			if currScore > maxElfScore {
				maxElfScore = currScore
			}
			currElf++
			currScore = 0
		} else {
			val, _ := strconv.Atoi(thisLine)
			currScore += val
		}

	}

	scores = append(scores, currScore)

	return scores

}

func part1(data []string) int {

	var result int

	scores := compile(data)

	for _, val := range scores {
		if val > result {
			result = val
		}
	}

	return result
}

func part2(data []string) int {

	var result int

	scores := compile(data)

	sort.Ints(scores)

	for _, val := range scores[(len(scores) - 3):] {
		result += val
	}

	return result
}
