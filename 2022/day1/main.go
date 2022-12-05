package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	currElf := 0
	//maxElf := 0
	maxElfScore := 0
	currScore := 0

	scores := []int{}

	for _, thisLine := range lines {

		if thisLine == "" {

			scores = append(scores, currScore)

			if currScore > maxElfScore {
				maxElfScore = currScore
				//maxElf = currElf
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

	fmt.Printf("Part A answer: %d\n", scores[len(scores)-1])
	fmt.Printf("Part B answer: %d\n", total)

}
