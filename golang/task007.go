package main

import "fmt"

func main() {
	fmt.Println(getPrimeNumberByIndex(10001))
}

func getPrimeNumberByIndex(n int) int {
	i := 2
	j := 0
	for {
		if IsPrime(i) {
			j++
			if j == n {
				return i
			}
		}
		i++
	}
}
