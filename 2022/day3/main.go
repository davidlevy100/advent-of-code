package main

import (
	"fmt"

	util "github.com/davidlevy100/advent-of-code/util"
)

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func part1(data []string) int {

	var result int

	letterScores := scoreMap()

	for _, thisWord := range data {
		result += getScore(thisWord, letterScores)
	}

	return result

}

func part2(data []string) int {

	var result int

	letterScores := scoreMap()

	for i := 0; i < len(data)-2; i += 3 {
		result += badges(data[i:i+3], letterScores)
	}

	return result

}

func getScore(s string, m map[rune]int) int {
	var result int

	var runemap = make(map[rune]bool)

	for _, thisRune := range s[:len(s)/2] {
		runemap[thisRune] = true
	}

	for _, thisRune := range s[(len(s) / 2):] {
		if runemap[thisRune] {
			result += m[thisRune]
			break
		}
	}

	return result
}

func scoreMap() map[rune]int {

	var result = make(map[rune]int)

	val := 1
	for thisRune := 'a'; thisRune <= 'z'; thisRune++ {
		result[thisRune] = val
		val++
	}
	for thisRune := 'A'; thisRune <= 'Z'; thisRune++ {
		result[thisRune] = val
		val++
	}

	return result

}

func badges(data []string, scores map[rune]int) int {
	var result int

	var exists = struct{}{}

	var sets = make([]map[rune]struct{}, 1)

	for _, thisLine := range data {
		thisSet := make(map[rune]struct{})

		for _, thisRune := range thisLine {
			thisSet[thisRune] = exists
		}

		sets = append(sets, thisSet)

	}

	var counts = make(map[rune]int)

	for _, thisSet := range sets {
		for key := range thisSet {
			counts[key]++
		}
	}

	for key, val := range counts {
		if val == len(data) {
			result += scores[key]
		}
	}

	return result
}
