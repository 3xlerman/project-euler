package main

import (
	"fmt"
	"strconv"
)

func main() {

	max := 0
	for i := 999; i > 100; i-- {
		for j := i - 1; j > 100; j-- {
			if isPalindrome(i*j) && i*j > max {
				max = i * j
			}
		}
	}
	fmt.Println(max)
}

func isPalindrome(num int) bool {
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
