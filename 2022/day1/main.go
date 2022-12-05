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

	lines, _ := util.GetInput(fullPath)

	currElf := 0
	maxElfScore := 0
	currScore := 0

	scores := []int{}

	for _, thisLine := range lines {

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

	sort.Ints(scores)

	total := 0

	for _, val := range scores[(len(scores) - 3):] {
		total += val
	}

	fmt.Printf("Part 1 answer: %d\n", scores[len(scores)-1])
	fmt.Printf("Part 2 answer: %d\n", total)

}
