package main

import (
	"fmt"
	"os"

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

		val, ok := checkCode(thisLine)
		if ok {
			result += val
		}

	}

	return result

}

func part2(data []string) int {

	var result int

	f, err := os.Create("results.txt")
	check(err)
	defer f.Close()

	for _, thisLine := range data {

		_, ok := checkCode(thisLine)
		if ok {
			sentence, val := shiftCypher(thisLine)
			fmt.Fprintf(f, "%d\t%s\n", val, sentence)
		}

	}

	return result

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
