package main

import (
	"fmt"
	"strconv"

	util "github.com/davidlevy100/advent-of-code/util"
)

var digits = map[string]int{
	"1":     1,
	"one":   1,
	"2":     2,
	"two":   2,
	"3":     3,
	"three": 3,
	"4":     4,
	"four":  4,
	"5":     5,
	"five":  5,
	"6":     6,
	"six":   6,
	"7":     7,
	"seven": 7,
	"8":     8,
	"eight": 8,
	"9":     9,
	"nine":  9,
}

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func part1(data []string) int {

	var result int

	vals := []int{}

	for _, thisLine := range data {
		vals = append(vals, getNum(thisLine))
	}

	for _, thisVal := range vals {
		result += thisVal
	}

	return result

}

func part2(data []string) int {

	var result int

	vals := []int{}

	for _, thisLine := range data {
		vals = append(vals, getDigits(thisLine))
	}

	for _, thisVal := range vals {
		result += thisVal
	}

	return result

}

func getNum(line string) int {
	var result int

	first, last := 0, 0

	for i := 0; i < len(line); i++ {
		if val, err := strconv.Atoi(string(line[i])); err == nil {
			first = val
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if val, err := strconv.Atoi(string(line[i])); err == nil {
			last = val
			break
		}
	}

	result = first*10 + last

	return result
}

func getDigits(line string) int {

	var result int = getFirstDigit(line)*10 + getLastDigit(line)
	return result
}

func getFirstDigit(line string) int {

	var result int

	start := 0

FOR:
	for ; start < len(line); start++ {

		for key, val := range digits {
			end := start + len(key)
			if end <= len(line) {
				word := line[start:end]
				if word == key {
					result = val
					break FOR
				}
			}
		}

	}

	return result

}

func getLastDigit(line string) int {

	var result int

	start := len(line) - 1

FOR:
	for ; start >= 0; start-- {

		for key, val := range digits {
			end := start + len(key)
			if end <= len(line) {
				word := line[start:end]
				if word == key {
					result = val
					break FOR
				}
			}
		}

	}

	return result

}
