package numbers

// Useful functions from Project Euler tasks

// Here are special functions

func GetTriangleNum(n int) (sum int) {

	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum
}

func CountCollatzSeq(num int) int {
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
