package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() (int, int) {
	file, err := os.Open("day4/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var counter int
	var alternateCounter int

	for scanner.Scan() {
		var intSlice []int
		if scanner.Text() != "" {
			firstSplit := strings.Split(scanner.Text(), ",")
			for i := range firstSplit {
				secondSplit := strings.Split(firstSplit[i], "-")
				for j := range secondSplit {
					v, _ := strconv.Atoi(secondSplit[j])
					intSlice = append(intSlice, v)
				}
			}
			if fullyContains(intSlice) {
				counter++
			}
			if findOverlap(intSlice) {
				alternateCounter++
			}
		}
	}
	return counter, alternateCounter
}
