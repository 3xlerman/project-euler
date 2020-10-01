package main

import (
	"fmt"
	"project-euler/golang/numbers"
)

func main() {

	max := 0
	for i := 999; i > 100; i-- {
		for j := i - 1; j > 100; j-- {
			if numbers.IsPalindrome(i*j) && i*j > max {
				max = i * j
			}
		}
	}
	fmt.Println(max)
}
