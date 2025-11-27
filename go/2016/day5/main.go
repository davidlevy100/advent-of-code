package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func main() {

	fmt.Printf("Part 1 answer: %s\n", part1("cxdnnyjw"))
	fmt.Printf("Part 2 answer: %s\n", part2("cxdnnyjw"))

}

func part1(data string) string {

	var result string

	for i := 0; i < MaxInt; i++ {
		try := fmt.Sprintf("%s%d", data, i)
		bTry := []byte(try)
		hex := fmt.Sprintf("%x", md5.Sum(bTry))

		if hex[:5] == "00000" {

			result = result + string(hex[5])

			if len(result) == 8 {
				break
			}

		}
	}

	return result

}

func part2(data string) string {

	var result string

	var chars = make([]string, 8)

	for i := 0; i < MaxInt; i++ {
		try := fmt.Sprintf("%s%d", data, i)
		bTry := []byte(try)
		hex := fmt.Sprintf("%x", md5.Sum(bTry))

		if hex[:5] == "00000" {

			index, err := strconv.Atoi(string(hex[5]))
			if err != nil || index >= 8 || chars[index] != "" {
				continue
			}

			val := string(hex[6])
			chars[index] = val

			if done(chars) {
				break
			}
		}
	}

	result = strings.Join(chars, "")

	return result

}

func done(ss []string) bool {

	for _, s := range ss {
		if s == "" {
			return false
		}
	}
	return true
}
