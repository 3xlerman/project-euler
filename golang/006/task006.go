package main

import (
	"fmt"
	"math"
)

func main() {
	sSquare := getSumOfSquares()
	sSum := getSquareOfSum()
	fmt.Printf("sSquare: %d; sSum: %d, difference: %f", sSquare, sSum, math.Abs(float64(sSquare-sSum)))
}

func getSumOfSquares() int {
	sum := 0
	for i := 1; i < 101; i++ {
		sum += i * i
	}
	return sum
}

func getSquareOfSum() int {
	sum := 0
	for i := 1; i < 101; i++ {
		sum += i
	}
	return sum * sum
}
