package main

import (
	"fmt"

	util "github.com/davidlevy100/advent-of-code/util"
)

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %s\n", part1(data))
	fmt.Printf("Part 2 answer: %s\n", part2(data))

}

func part1(data []string) string {

	var result string

	var counts []map[rune]int

	for i := 0; i < len(data[0]); i++ {
		counts = append(counts, make(map[rune]int))
	}

	for _, thisLine := range data {
		for i, thisRune := range thisLine {
			counts[i][thisRune]++
		}
	}

	for _, thisMap := range counts {
		var maxCount int
		var maxRune rune
		for key, value := range thisMap {
			if value > maxCount {
				maxCount = value
				maxRune = key
			}
		}

		result = result + string(maxRune)
	}

	return result

}

func part2(data []string) string {

	var result string

	var counts []map[rune]int

	for i := 0; i < len(data[0]); i++ {
		counts = append(counts, make(map[rune]int))
	}

	for _, thisLine := range data {
		for i, thisRune := range thisLine {
			counts[i][thisRune]++
		}
	}

	for _, thisMap := range counts {
		var minCount int = 100
		var minRune rune
		for key, value := range thisMap {
			if value < minCount {
				minCount = value
				minRune = key
			}
		}

		result = result + string(minRune)
	}

	return result

}
