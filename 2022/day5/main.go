package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	util "github.com/davidlevy100/advent-of-code/util"
)

type stack []string

// IsEmpty: check if stack is empty
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

func (s *stack) PushMany(st stack) {
	*s = append(*s, st...)
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *stack) PopMany(items int) (stack, bool) {
	if s.IsEmpty() {
		return stack{}, false
	} else {
		index := len(*s) - items // Get the index of the top most element.
		newStack := (*s)[index:] // Index into the slice and obtain the element.
		*s = (*s)[:index]        // Remove it from the stack by slicing it off.
		return newStack, true
	}
}

func (s *stack) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}
}

func main() {

	data, _ := util.GetInput("input.txt")

	fmt.Printf("Part 1 answer: %v\n", part1(data))
	fmt.Printf("Part 2 answer: %v\n", part2(data))

}

func part1(data []string) []string {

	var result []string

	stacks, commands := parseData(data)

	for _, thisCommand := range commands {
		for i := 0; i < thisCommand[0]; i++ {
			data, _ := stacks[thisCommand[1]-1].Pop()
			stacks[thisCommand[2]-1].Push(data)
		}
	}

	for _, thisStack := range stacks {
		data, _ := thisStack.Peek()
		result = append(result, data)
	}

	return result

}

func part2(data []string) []string {

	var result []string

	stacks, commands := parseData(data)

	for _, thisCommand := range commands {

		data, _ := stacks[thisCommand[1]-1].PopMany(thisCommand[0])
		stacks[thisCommand[2]-1].PushMany(data)

	}

	for _, thisStack := range stacks {
		data, _ := thisStack.Peek()
		result = append(result, data)
	}

	return result

}

func parseData(data []string) ([]stack, [][]int) {

	var commandIndex int

	for index, thisLine := range data {
		if thisLine == "" {
			commandIndex = index + 1
			break
		}
	}

	var columns int

	for _, thisRune := range data[commandIndex-2] {
		if unicode.IsDigit(thisRune) {
			columns = int(thisRune - '0')
		}
	}

	var stacks = make([]stack, columns)

	for i := commandIndex - 3; i >= 0; i-- {
		for j := 1; j < len(data[i]); j += 4 {

			if data[i][j] >= 65 && data[i][j] <= 90 {
				stacks[j/4].Push(string(data[i][j]))
			}

		}
	}

	var commands [][]int

	for _, thisLine := range data[commandIndex:] {
		data := strings.Fields(thisLine)
		amount, _ := strconv.Atoi(data[1])
		from, _ := strconv.Atoi(data[3])
		to, _ := strconv.Atoi(data[5])
		commands = append(commands, []int{amount, from, to})
	}

	return stacks, commands
}
