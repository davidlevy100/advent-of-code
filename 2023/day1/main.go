package main

import (
	"fmt"
	"strconv"
	"unicode"

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

	nums := []int{}

	for _, thisRune := range line {
		if unicode.IsDigit(thisRune) {
			val, _ := strconv.Atoi(string(thisRune))
			nums = append(nums, val)
		}
	}

	if len(nums) > 0 {
		result = nums[0]*10 + nums[len(nums)-1]
	}

	return result
}

func getDigits(line string) int {

	fmt.Println(line)
	var result int

	nums := []int{}

	start := 0
	var found bool

	for start < len(line) {

		if found {
			start++
			found = false
		}

		for key, val := range digits {
			end := start + len(key)
			if end < len(line) {
				word := line[start:end]
				fmt.Println(word, key)
				if word == key {
					nums = append(nums, val)
					start += len(key) -1
					found = true
					break
				}
			}
		}

	}

	if len(nums) > 0 {
		result = nums[0]*10 + nums[len(nums)-1]
	}
	fmt.Println(nums)
	return result

}
