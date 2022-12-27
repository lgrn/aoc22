package day5

import (
	"strconv"
	"strings"
)

// parseMoveOrder() takes a string with syntax:
// "move N from N to N" -- it returns []int{N,N,N}
func parseMoveOrder(s string) []int {
	var returnSlice []int
	splitString := strings.Split(s, " ")
	move, _ := strconv.Atoi(splitString[1])
	from, _ := strconv.Atoi(splitString[3])
	to, _ := strconv.Atoi(splitString[5])
	returnSlice = append(returnSlice, move, from, to)
	return returnSlice
}

// executeMoveOrder() takes a 3-len []int and moves N crates from N to N.
// A "move" copies X[len(X)] to Y[len(Y)+1], then deletes X[len(X)]
func executeMoveOrder(o []int) bool {
	moveAmount, moveFrom, moveTo := o[0], o[1], o[2]
	for ; moveAmount > 0; moveAmount-- {
		stacks[moveTo][len(stacks[moveTo])+1] = stacks[moveFrom][len(stacks[moveFrom])]
		delete(stacks[moveFrom], len(stacks[moveFrom])) // delete last object in 'moveFrom'
	}
	return true
}

// executeAlternateMoveOrder() takes a 3-len []int and moves N crates from N to N.
// Unlike executeMoveOrder, this function moves all crates at once, keeping the internal order.
// It also works against 'stacks2' which is a copy of 'stacks'
func executeAlternateMoveOrder(o []int) bool {
	moveAmount, moveFrom, moveTo := o[0], o[1], o[2]
	for i := moveAmount; i > 0; i-- {
		stacks2[moveTo][len(stacks2[moveTo])+1] = stacks2[moveFrom][len(stacks2[moveFrom])+1-i]
	}
	for j := moveAmount; j > 0; j-- {
		delete(stacks2[moveFrom], len(stacks2[moveFrom]))
	}
	return true
}
