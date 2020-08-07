package main

import "fmt"

func main() {
	fmt.Println(getPrimeNumberByIndex(10001))
}

func getPrimeNumberByIndex(n int) int {
	i := 2
	j := 0
	for {
		if isPrime(i) {
			j++
			if j == n {
				return i
			}
		}
		i++
	}
}

func isPrime(num int) bool {
	if num <= 0 {
		return false
	} else {
		for j := num - 1; j > 0; j-- {
			if num%j != 0 {
				continue
			} else {
				if j == 1 {
					return true
				} else {
					return false
				}
			}
		}
		return true
	}
}
