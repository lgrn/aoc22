package day5

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// One separate stack for each simulation.
// These are filled with values in a very ugly way in Solve(),
// but I can't think of a better way. At least most IDEs let you
// collapse 'stacks' and 'stacks2' when reading the code.
var stacks map[int]map[int]rune
var stacks2 map[int]map[int]rune

func Solve() (string, string) {

	stacks = map[int]map[int]rune{
		1: {
			1: 'B',
			2: 'Z',
			3: 'T',
		},
		2: {
			1: 'V',
			2: 'H',
			3: 'T',
			4: 'D',
			5: 'N',
		},
		3: {
			1: 'B',
			2: 'F',
			3: 'M',
			4: 'D',
		},
		4: {
			1: 'T',
			2: 'J',
			3: 'G',
			4: 'W',
			5: 'V',
			6: 'Q',
			7: 'L',
		},
		5: {
			1: 'W',
			2: 'D',
			3: 'G',
			4: 'P',
			5: 'V',
			6: 'F',
			7: 'Q',
			8: 'M',
		},
		6: {
			1: 'V',
			2: 'Z',
			3: 'Q',
			4: 'G',
			5: 'H',
			6: 'F',
			7: 'S',
		},
		7: {
			1: 'Z',
			2: 'S',
			3: 'N',
			4: 'R',
			5: 'L',
			6: 'T',
			7: 'C',
			8: 'W',
		},
		8: {
			1: 'Z',
			2: 'H',
			3: 'W',
			4: 'D',
			5: 'J',
			6: 'N',
			7: 'R',
			8: 'M',
		},
		9: {
			1: 'M',
			2: 'Q',
			3: 'L',
			4: 'F',
			5: 'D',
			6: 'S',
		},
	}
	stacks2 = map[int]map[int]rune{
		1: {
			1: 'B',
			2: 'Z',
			3: 'T',
		},
		2: {
			1: 'V',
			2: 'H',
			3: 'T',
			4: 'D',
			5: 'N',
		},
		3: {
			1: 'B',
			2: 'F',
			3: 'M',
			4: 'D',
		},
		4: {
			1: 'T',
			2: 'J',
			3: 'G',
			4: 'W',
			5: 'V',
			6: 'Q',
			7: 'L',
		},
		5: {
			1: 'W',
			2: 'D',
			3: 'G',
			4: 'P',
			5: 'V',
			6: 'F',
			7: 'Q',
			8: 'M',
		},
		6: {
			1: 'V',
			2: 'Z',
			3: 'Q',
			4: 'G',
			5: 'H',
			6: 'F',
			7: 'S',
		},
		7: {
			1: 'Z',
			2: 'S',
			3: 'N',
			4: 'R',
			5: 'L',
			6: 'T',
			7: 'C',
			8: 'W',
		},
		8: {
			1: 'Z',
			2: 'H',
			3: 'W',
			4: 'D',
			5: 'J',
			6: 'N',
			7: 'R',
			8: 'M',
		},
		9: {
			1: 'M',
			2: 'Q',
			3: 'L',
			4: 'F',
			5: 'D',
			6: 'S',
		},
	}

	file, err := os.Open("day5/input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "move") {
			moveOrder := parseMoveOrder(scanner.Text())
			executeMoveOrder(moveOrder)
			executeAlternateMoveOrder(moveOrder)
		}
	}

	var returnString string
	var alternateReturnString string

	for i := 1; i < len(stacks)+1; i++ {
		returnString = returnString + string(stacks[i][len(stacks[i])])
	}
	for i := 1; i < len(stacks2)+1; i++ {
		alternateReturnString = alternateReturnString + string(stacks2[i][len(stacks2[i])])
	}

	return returnString, alternateReturnString
}
