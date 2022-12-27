package day6

// findMarker() takes string, int and returns an int
// representing the end position of an 'int' sized block
// where all runes contained are unique
func findMarker(s string, u int) int {
	for i := u; i <= len(s); i++ {
		if runesAreUnique(s[i-u : i]) {
			return i
		}
	}
	return 0
}

// runesAreUnique() takes a string and returns a bool indicating
// whether all runes in that string are unique
func runesAreUnique(s string) bool {
	occurrenceMap := make(map[rune]bool)
	for _, r := range s {
		_, occurred := occurrenceMap[r] // has the rune occurred?
		if occurred {                   // then it's not unique
			return false
		}
		occurrenceMap[r] = true // first occurrence of rune
	}
	return true
}
