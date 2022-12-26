package day4

//fullyContains() takes a slice of exactly 4 ints and returns true
//if one set of two completely fits within the other, i.e. if
//s[0]-s[1] is completely contained in s[2]-s[3] or vice versa.
func fullyContains(s []int) bool {
	if s[0] >= s[2] && s[1] <= s[3] { // s[0-1] starts within s[2-3] without overshooting
		return true
	} else if s[2] >= s[0] && s[3] <= s[1] { // s[2-3] starts within s[0-1] without overshooting
		return true
	}
	return false
}

//findOverlap() takes a slice of exactly 4 ints and returns true
//if one set of two *partly* overlaps with the other, i.e. if
//s[0]-s[1] goes into s[2]-s[3] in any end, or vice versa
func findOverlap(s []int) bool {
	if s[0] <= s[3] && s[0] >= s[2] { // 0-1 goes into 2-3 from right, but does not overshoot s[2]
		return true
	}
	if s[1] >= s[2] && s[1] <= s[3] { // 0-1 goes into 2-3 from left, but does not overshoot s[3]
		return true
	}
	if s[2] <= s[1] && s[2] >= s[0] { // 2-3 goes into 0-1 from right, but does not overshoot s[0]
		return true
	}
	if s[3] >= s[0] && s[3] <= s[1] { // 2-3 goes into 0-1 from left, but does not overshoot s[1]
		return true
	}
	return false
}
