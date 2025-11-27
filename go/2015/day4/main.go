package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {

	fmt.Printf("Part 1 answer: %d\n", part1("bgvyzdsv"))
	fmt.Printf("Part 2 answer: %d\n", part2("bgvyzdsv"))

}

func part1(data string) int {

	var result int

	hashString := "111111"

	for hashString[:5] != "00000" {
		result++
		h := md5.New()
		thisString := fmt.Sprintf("%s%d", data, result)
		io.WriteString(h, thisString)
		hashString = fmt.Sprintf("%x", h.Sum(nil))
	}

	return result

}

func part2(data string) int {

	var result int

	hashString := "111111"

	for hashString[:6] != "000000" {
		result++
		h := md5.New()
		thisString := fmt.Sprintf("%s%d", data, result)
		io.WriteString(h, thisString)
		hashString = fmt.Sprintf("%x", h.Sum(nil))
	}

	return result

}
