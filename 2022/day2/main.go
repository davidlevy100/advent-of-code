package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

var scoresRev = map[int]string{
	1: "A",
	2: "B",
	3: "C",
}

var names = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
}

func main() {
	lines, err := getInput()
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

func getInput() ([]string, error) {

	f, err := os.Open("input.txt")

	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
