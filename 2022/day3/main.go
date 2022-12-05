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

	var partOneResult int

	for _, thisWord := range lines {
		partOneResult += getScore(thisWord, letterScore)
	}

	fmt.Printf("\nPart 1 answer: %d\n", partOneResult)

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
