package day2

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() (int, int) {
	file, err := os.Open("day2/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var score int
	var alternateScore int

	for scanner.Scan() {
		if scanner.Text() != "" {
			s := scanner.Text()
			var r1, r2 = calculateGameScore(s[0:1], s[len(s)-1:])
			score += r1
			alternateScore += r2
		}
	}

	return score, alternateScore
}
