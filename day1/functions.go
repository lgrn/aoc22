package day1

import "sort"

// functions

// findLargestNumber() takes a slice of ints and returns the largest one in the set.
// The assumption is that there is no way to find the largest number in
// an unsorted set without looking at all of them once.
func findLargestNumber(s []int) int {
	var max int
	for _, n := range s {
		if n > max {
			max = n
		}
	}
	return max
}

// findNLargestNumbers() takes a slice of ints and an amount, and returns the 'amount'
// largest numbers in that slice.
func findNLargestNumbers(s []int, a int) []int {
	// sort incoming slice in-place
	sort.Ints(s)
	var returnSlice []int
	for ; a > 0; a-- {
		returnSlice = append(returnSlice, s[len(s)-a])
	}
	return returnSlice
}

// sumSlice takes a slice of ints and returns their sum
func sumSlice(s []int) int {
	var sum int
	for _, n := range s {
		sum += n
	}
	return sum
}
