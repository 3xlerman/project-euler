package numbers

// Functions to work with prime numbers

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

func GetPrimeNumberByIndex(n int) int {
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
