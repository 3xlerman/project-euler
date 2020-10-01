package numbers

import "strconv"

// Functions to work with palindromes

func IsPalindrome(num int) bool {
	numStr := strconv.Itoa(num)

	first := numStr[:len(numStr)/2]
	second := numStr[len(numStr)/2:]
	secondReverse := make([]byte, len(second))

	for i := 0; i < len(second); i++ {
		secondReverse[i] = second[len(second)-i-1]
	}

	if first == string(secondReverse) {
		return true
	} else {
		return false
	}
}
