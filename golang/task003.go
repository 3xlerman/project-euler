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


