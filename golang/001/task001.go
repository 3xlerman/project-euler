package main

import "fmt"

func main() {
	var sum int
	N := 1000

	for i := 0; i < N; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
