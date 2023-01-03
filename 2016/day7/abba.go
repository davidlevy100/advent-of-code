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

	re := regexp.MustCompile(`\[(.*?)\]`)

	brackets := re.FindAllStringSubmatch(s, -1)

	// return false if any bracketed text
	// has the abba pattern
	for _, word := range brackets {
		if hasAbba(word[1]) {
			return false
		}
	}

	// return true if any remaining text
	// has the abba pattern
	noBrackets := re.ReplaceAllString(s, " ")
	words := strings.Fields(noBrackets)

	for _, word := range words {
		if hasAbba(word) {
			return true
		}
	}
	return false
}
