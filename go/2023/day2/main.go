package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/davidlevy100/advent-of-code/util"
)

type game struct {
	ID     int
	scores []map[string]int
}

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %d\n", part1(data))
	fmt.Printf("Part 2 answer: %d\n", part2(data))

}

func part1(data []string) int {

	var result int

	amounts := map[string]int{"red": 12, "green": 13, "blue": 14}

	for _, thisLine := range data {
		game := parseLine(thisLine)

		id, err := testGame(amounts, game)
		if err == nil {
			result += id
		}
	}

	return result

}

func part2(data []string) int {

	var result int

	for _, thisLine := range data {
		game := parseLine(thisLine)

		result += getPower(game)
	}

	return result

}

func parseLine(line string) game {

	var result game

	mainParts := strings.Split(line, ":")

	idParts := strings.Split(mainParts[0], " ")
	id, _ := strconv.Atoi(idParts[1])
	result.ID = id

	subGames := strings.Split(mainParts[1], ";")

	for _, thisSubGame := range subGames {

		var grabData = make(map[string]int)

		grabs := strings.Split(thisSubGame, ",")
		for _, thisGrab := range grabs {
			data := strings.Fields(thisGrab)
			val, _ := strconv.Atoi(data[0])
			grabData[data[1]] = val
		}
		result.scores = append(result.scores, grabData)
	}

	return result

}

func testGame(amounts map[string]int, g game) (int, error) {
	var result int = g.ID
	var err error

FOR:
	for _, thisScore := range g.scores {
		for key, val := range thisScore {
			if val > amounts[key] {
				err = fmt.Errorf("invalid game")
				break FOR
			}
		}
	}

	return result, err
}

func getPower(g game) int {

	var result int = 1

	var amounts = make(map[string]int)

	for _, thisScore := range g.scores {
		for key, val := range thisScore {
			if val > amounts[key] {
				amounts[key] = val
			}
		}
	}

	for _, val := range amounts {
		result *= val
	}

	return result

}
