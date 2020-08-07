package main

import "fmt"

func main() {
	num1 := 1
	num2 := 1
	sum := 0

	for {
		num3 := num1 + num2
		num1 = num2
		num2 = num3
		if num3%2 == 0 {
			if num3 > 4000000 {
				break
			} else {
				sum += num3
			}
		}
	}
	fmt.Println(sum)
}

