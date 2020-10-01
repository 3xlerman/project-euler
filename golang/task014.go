package main

import (
	"fmt"
	"project-euler/golang/numbers"
)

func main() {
	var max, maxnum int
	for i := 2; i < 1000000; i++ {
		count := numbers.CountCollatzSeq(i)
		if count > max {
			max = count
			maxnum = i
		}
	}
	fmt.Println(maxnum)
}
