package main

import "fmt"

func main() {
	max := 0
	for i := 100; ; i++ {
		number := GetTriangleNum(i)
		n := make(map[int]int)
		multi := 1

		sameNum := number
		for i := 2; number != 1; i++ {
			if IsPrime(i) {
				if number%i == 0 {
					n[i]++
					number /= i
					i--
				}
			}
		}
		for _, val := range n {
			multi *= val + 1
		}
		if multi > max {
			max = multi
			fmt.Println(max)
		}
		if max > 500 {
			fmt.Println("number", sameNum, "has", max)
			break
		}
	}

}

func GetTriangleNum(n int) (sum int) {

	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum
}

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

func GetDividers(n, d int) (count int) {
	for n != 1 {
		if n%d == 0 {
			n /= d
			count++
		}
	}
	return count
}
