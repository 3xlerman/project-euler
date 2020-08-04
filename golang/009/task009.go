package main

import "fmt"

func main() {
	var xa, xb, xc, multi int

	for m := 1000; m > 0; m-- {
		for n := m - 1; n > 0; n-- {
			a := 2 * m * n
			b := m*m - n*n
			c := m*m + n*n
			if a+b+c == 1000 {
				xa, xb, xc = a, b, c
				multi = a * b * c
			}
		}
	}
	fmt.Printf("a: %d\nb: %d\nc: %d\nmulti: %d", xa, xb, xc, multi)
}
