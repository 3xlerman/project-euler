package main

import "fmt"

func main() {
	var max int
	for i := 2; i < 1000000; i++ {
		count := format(i)
		if count > max {
			max = count
			fmt.Println(i)
		}

	}
}

func format(num int) int {
	count := 0
	for ; num != 1; {
		if num%2 == 0 {
			num /= 2
		} else {
			num = 3*num + 1
		}
		count++
	}
	return count
}
