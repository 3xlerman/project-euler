package main

import (
	"fmt"
	"math"
)

func main() {
	var bigNumber = 600851475143
	maxF := math.Sqrt(float64(bigNumber))
	n := bigNumber
	f := 3

	var lastF int

	for n > 1 && f <= int(maxF) {
		if n%f == 0 {
			n = n / f
			lastF = f
			for n%f == 0 {
				n = n / f
			}
		}
		f += 2
	}
	if n == 1 {
		fmt.Println(lastF)
	} else {
		fmt.Println(n)
	}

}

func IsPrime(num int) bool {
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
