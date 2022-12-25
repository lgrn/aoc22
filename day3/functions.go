package day3

import (
	"log"
	"math"
	"strings"
	"unicode"
)

// splitStringInHalf() takes a string and returns a []string with two elements
// each composing half of the original string.
func splitStringInHalf(s string) []string {
	if math.Mod(float64(len(s)), 2) != 0 {
		log.Fatal("String '" + s + "' is not divisible by two!")
	}
	var returnSlice []string
	returnSlice = append(returnSlice, s[0:len(s)/2])
	returnSlice = append(returnSlice, s[len(s)/2:])
	return returnSlice
}

// findCommonRune() takes a []string (s[]) and looks for a common rune in all
// elements of that slice. It uses s[0] as reference where it iterates over each
// letter, checking it against s[N] in sequence. As soon as it detects a failure,
// it skips that letter in s[0] and moves on to the next.
// If no matches are found, it returns an empty rune.
func findCommonRune(s []string) rune {
	var emptyRune rune
	for _, ref := range s[0] { // iterate over every rune s[0]
	skip: // skip back to here if Contains fails, move to the next letter in s[0]
		for j := 1; j < len(s); j++ { // iterate over s[N], look for rune 'ref'
			if !strings.Contains(s[j], string(ref)) {
				break skip // failure detected, move to next ref
			}
			if j == len(s)-1 {
				return ref // found something that matches all elements in s[]
			}
		}

	}
	return emptyRune
}

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
// v = 118 should give 22 points = diff 96
// P = 80  should give 42 points = diff 38

// calculatePriority() calculates the priority score of a rune, returned as int.
func calculatePriority(c rune) int {
	if unicode.IsUpper(c) {
		return int(c) - 38
	} else {
		return int(c) - 96
	}
}
