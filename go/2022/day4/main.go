package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	util "github.com/davidlevy100/advent-of-code/util"
)

type Interval struct {
	start, end int
}

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func part1(data []string) int {

	var result int

	for _, thisLine := range data {
		intervals := parseLine(thisLine)
		if fullyOverlap(intervals) {
			result++
		}
	}

	return result

}

func part2(data []string) int {

	var result int

	for _, thisLine := range data {
		intervals := parseLine(thisLine)
		if partiallyOverlap(intervals) {
			result++
		}
	}

	return result

}

func parseLine(s string) []Interval {

	var result []Interval

	intervals := strings.Split(s, ",")

	for _, thisString := range intervals {
		times := strings.Split(thisString, "-")
		start, _ := strconv.Atoi(times[0])
		end, _ := strconv.Atoi(times[1])

		thisInterval := Interval{start, end}

		result = append(result, thisInterval)
	}
	return result
}

func fullyOverlap(si []Interval) bool {
	var result bool

	sort.SliceStable(si, func(i, j int) bool {
		return si[i].start < si[j].start
	})

	if si[0].end >= si[1].end {
		result = true
	}

	switch {
	case si[0].end >= si[1].end:
		result = true
	case si[0].start == si[1].start && si[1].end >= si[0].end:
		result = true
	}

	return result
}

func partiallyOverlap(si []Interval) bool {
	var result bool

	sort.SliceStable(si, func(i, j int) bool {
		return si[i].start < si[j].start
	})

	if si[0].end >= si[1].end {
		result = true
	}

	switch {
	case si[0].end >= si[1].end:
		result = true
	case si[0].start == si[1].start && si[1].end >= si[0].end:
		result = true
	case si[1].start <= si[0].end:
		result = true
	}

	return result
}
