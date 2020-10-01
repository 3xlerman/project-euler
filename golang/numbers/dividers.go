package numbers

import "fmt"

// Functions to work with dividers

func PrintDividers(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i, "divides the number")
		}
	}
}

func CountDividers(n int) (count int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

func CountDividersFromEnd(n, d int) (count int) {
	for n != 1 {
		if n%d == 0 {
			n /= d
			count++
		}
	}
	return count
}
