package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	util "github.com/davidlevy100/advent-of-code/util"
)

func main() {
	path, _ := os.Getwd()
	fullPath := filepath.Join(path, "input.txt")

	lines, err := util.GetInput(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nPart 1 answer: %d\n", part1(lines))
	fmt.Printf("\nPart 2 answer: %d\n", part2(lines))
}

func part1(data []string) int {

	var result int
	var letterScore = make(map[rune]int)

	val := 1
	for thisRune := 'a'; thisRune <= 'z'; thisRune++ {
		letterScore[thisRune] = val
		val++
	}
	for thisRune := 'A'; thisRune <= 'Z'; thisRune++ {
		letterScore[thisRune] = val
		val++
	}

	for _, thisWord := range data {
		result += getScore(thisWord, letterScore)
	}

	return result

}

func part2(data []string) int {

	var result int

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
