package main

import (
	"regexp"
	"strings"
)

// isABBA tests whether a 4 character string
// follows the ABBA pattern
func isABBA(s string) bool {

	if len(s) != 4 {
		return false
	}

	if s[0] == s[3] && s[1] == s[2] && s[0] != s[1] {
		return true
	}

	return false

}

// hasAbba uses the sliding window algorithm
// and runs isAbba on each 4 character window
func hasAbba(s string) bool {

	if len(s) < 4 {
		return false
	}

	for i := 0; i < len(s)-3; i++ {
		if isABBA(s[i : i+4]) {
			return true
		}
	}

	return false

}

// abbaTest determines if the left or right has abba
// returns false if the middle has the abba pattern
func abbaTest(s string) bool {

	brackets, noBrackets := splitText(s)

	// return false if any bracketed text
	// has the abba pattern
	for _, word := range brackets {
		if hasAbba(word) {
			return false
		}
	}

	// return true if any remaining text
	// has the abba pattern
	for _, word := range noBrackets {
		if hasAbba(word) {
			return true
		}
	}
	return false
}

func splitText(s string) (brackets, noBrackets []string) {

	re := regexp.MustCompile(`\[(.*?)\]`)

	b := re.FindAllStringSubmatch(s, -1)

	for _, word := range b {
		brackets = append(brackets, word[1])
	}

	nb := re.ReplaceAllString(s, " ")
	noBrackets = strings.Fields(nb)

	return brackets, noBrackets
}

// isABA tests whether a 3 character string
// follows the ABA pattern
func isABA(s string) bool {

	if len(s) != 3 {
		return false
	}

	if s[0] == s[2] && s[0] != s[1] {
		return true
	}

	return false

}

// hasAba uses the sliding window algorithm
// and runs isAba on each 3 character window
func hasAba(s string) (string, bool) {

	if len(s) < 3 {
		return "", false
	}

	for i := 0; i < len(s)-2; i++ {
		word := s[i : i+3]
		if isABA(word) {
			newWord := word[1:] + string(word[0])
			return newWord, true
		}
	}

	return "", false

}

// abaTest determines if the left or right has abba
// returns false if the middle has the abba pattern
func abaTest(s string) bool {

	brackets, noBrackets := splitText(s)

	// find aba pattern
	pattern, ok := findAba(noBrackets)
	if !ok {
		return false
	}

	for _, word := range brackets {

		if findPattern(word, pattern) {
			return true
		}
	}

	return false

}

func findAba(ss []string) (string, bool) {

	for _, word := range ss {
		word, ok := hasAba(word)
		if ok {
			return word, true
		}
	}
	return "", false
}

func findPattern(s, p string) bool {

	for i := 0; i < len(s)-2; i++ {
		word := s[i : i+3]

		if word == p {
			return true
		}
	}

	return false

}
