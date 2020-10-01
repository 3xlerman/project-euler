package main

import (
	"fmt"
	"project-euler/golang/numbers"
)

func main() {
	max := 0
	for i := 100; ; i++ {
		number := numbers.GetTriangleNum(i)
		n := make(map[int]int)
		multi := 1

		sameNum := number
		for i := 2; number != 1; i++ {
			if numbers.IsPrime(i) {
				if number%i == 0 {
					n[i]++
					number /= i
					i--
				}
			}
		}
		for _, val := range n {
			multi *= val + 1
		}
		if multi > max {
			max = multi
			fmt.Println(max)
		}
		if max > 500 {
			fmt.Println("number", sameNum, "has", max)
			break
		}
	}

}








