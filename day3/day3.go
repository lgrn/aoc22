package day3

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() (int, int) {
	file, err := os.Open("day3/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var priority int
	var alternatePriority int
	var alternateRows []string

	for scanner.Scan() {
		if scanner.Text() != "" {
			// #1: calculate priority based on common rune in a split row
			commonCharacter := findCommonRune(splitStringInHalf(scanner.Text()))
			priority += calculatePriority(commonCharacter)

			// #2: calculate alternate priority over three unsplit lines
			alternateRows = append(alternateRows, scanner.Text())
			if len(alternateRows) == 3 {
				alternateCommonCharacter := findCommonRune(alternateRows)
				alternatePriority += calculatePriority(alternateCommonCharacter)
				alternateRows = []string{}
			}
		}
	}

	return priority, alternatePriority
}
