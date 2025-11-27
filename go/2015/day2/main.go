package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

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
		result += paperSize(thisLine)
	}

	return result

}

func part2(data []string) int {

	var result int

	for _, thisLine := range data {
		result += ribbonSize(thisLine)
	}

	return result

}

func paperSize(s string) int {
	var result int

	var vals []int
	valStrings := strings.Split(s, "x")

	for _, thisChar := range valStrings {
		val, _ := strconv.Atoi(thisChar)
		vals = append(vals, val)
	}

	sort.Ints(vals)
	result = vals[0]*vals[1] + 2*vals[0]*vals[1] + 2*vals[1]*vals[2] + 2*vals[0]*vals[2]

	return result
}

func ribbonSize(s string) int {
	var result int

	var vals []int
	valStrings := strings.Split(s, "x")

	for _, thisChar := range valStrings {
		val, _ := strconv.Atoi(thisChar)
		vals = append(vals, val)
	}

	sort.Ints(vals)
	result = 2*vals[0] + 2*vals[1] + vals[0]*vals[1]*vals[2]

	return result
}
