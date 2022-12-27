package day6

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() (int, int) {
	file, err := os.Open("day6/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	return findMarker(line, 4), findMarker(line, 14)
}
