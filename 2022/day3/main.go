package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	lines, err := getInput()
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
