package main

import "fmt"

func main() {
	var numbers [2000000]bool

	for i := 0; i < 2000000; i++ {
		numbers[i] = true
	}

	var sum int
	for i := 2; i < 2000000; i++ {
		if numbers[i] {
			for j := i * i; j < 2000000; j += i {
				numbers[j] = false
			}
			sum += i
		}
	}
	fmt.Println(sum)
}
