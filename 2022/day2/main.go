package main

import (
	"fmt"
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

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func part1(data []string) int {

	var result int

	for _, thisLine := range data {
		plays := strings.Fields(thisLine)
		score := getScore(plays[0], plays[1])
		result += score
	}

	return result

}

func part2(data []string) int {
	var result int

	for _, thisLine := range data {
		plays := strings.Fields(thisLine)
		score := choosePlay(plays[0], plays[1])
		result += score
	}

	return result

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
