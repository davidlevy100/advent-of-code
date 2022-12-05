package advent2022

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	advent2022 "github.com/davidlevy100/advent-of-code/advent2022/textreader"
)

func Day1() []string {

	path, _ := os.Getwd()
	fullPath := filepath.Join(path, "input.txt")

	lines, _ := advent2022.GetInput(fullPath)

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

	part1 := fmt.Sprintf("Part 1 answer: %d\n", scores[len(scores)-1])
	part2 := fmt.Sprintf("Part 2 answer: %d\n", total)

	return []string{part1, part2}

}
