package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	util "github.com/davidlevy100/advent-of-code/util"
)

var winner = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var loser = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

var scores = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func main() {
	path, _ := os.Getwd()
	fullPath := filepath.Join(path, "input.txt")

	lines, err := util.GetInput(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	var partOneResult int

	for _, thisLine := range lines {
		plays := strings.Fields(thisLine)
		score := getScore(plays[0], plays[1])
		partOneResult += score
	}

	fmt.Printf("\nPart 1 answer: %d\n", partOneResult)

	var partTwoResult int

	for _, thisLine := range lines {
		plays := strings.Fields(thisLine)
		score := choosePlay(plays[0], plays[1])
		partTwoResult += score
	}

	fmt.Printf("\nPart 2 answer: %d\n", partTwoResult)
}

func choosePlay(a, b string) int {
	var result int

	switch b {
	case "X":
		result += getScore(a, loser[a])
	case "Y":
		result += getScore(a, a)
	case "Z":
		result += getScore(a, winner[a])
	}

	return result
}

func getScore(a, b string) int {
	var result int

	result += scores[b]

	switch {
	case scores[a] == scores[b]:
		result += 3
	case b == winner[a]:
		result += 6
	}

	return result
}
