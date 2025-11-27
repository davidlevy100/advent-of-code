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

	for _, thisLine := range data {
		if threeVowels(thisLine) &&
			hasDoubles(thisLine) &&
			!badStrings(thisLine) {
			result++
		}
	}

	return result

}

func part2(data []string) int {

	var result int

	for _, thisLine := range data {
		if hasPairs(thisLine) &&
			hasStraddles(thisLine) {
			result++
		}
	}

	return result

}

func threeVowels(s string) bool {

	var vowels int

	for _, thisRune := range s {
		switch thisRune {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
	}

	return vowels >= 3

}

func hasDoubles(s string) bool {

	var result bool

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			result = true
			break
		}
	}

	return result
}

func badStrings(s string) bool {

	baddies := map[string]bool{
		"ab": true,
		"cd": true,
		"pq": true,
		"xy": true,
	}

	var result bool

	for i := 0; i < len(s)-1; i++ {
		if _, ok := baddies[s[i:i+2]]; ok {
			result = true
			break
		}
	}

	return result
}

func hasPairs(s string) bool {

	var result bool

	var pairs = make(map[string]int)

	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if prev, ok := pairs[pair]; ok {
			if i-prev > 1 {
				result = true
				break
			}
		} else {
			pairs[pair] = i
		}
	}

	return result

}

func hasStraddles(s string) bool {

	var result bool

	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			result = true
			break
		}
	}

	return result
}
