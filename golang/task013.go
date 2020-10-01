package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)


const evilNumber = 8231174977792365466257246923322810917141914302881971032885978066697608929386382850253334033441306557801612781592181500556186883646842009047023053081172816430487623791969842487255036638784583114876969321549028104240201383351244621814417734706378329949063625966649858761822122522551248676453367720186971698544312419572409913959008952310058822

func main() {
	numbers := getAllNumbers()
	sum := SumNumbers(numbers)
	s := strconv.FormatFloat(sum, 'f', -1, 64)
	fmt.Printf("sum is %s", s[:10] )

}

func getAllNumbers() []byte {

	file, err := os.Open("golang/task013.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	numbers := make([]byte, 16000)

	for {
		_, err := file.Read(numbers)
		if err == io.EOF {
			break
		}

	}

	return numbers
}

func SumNumbers(numbers []byte) (sum float64) {

	s := string(numbers)
	c := 0
	for i := 0; ; i++ {
		if string(s[i]) == "\n" {
			s = s[:i] + s[i+1:]
			c++
			fmt.Println(c, "--", s[:i])
		}
		if c == 50 {
			break
		}
	}

	for i := 0; i < 50; i++ {
		temp := s[i*50 : i*50+50]
		num, err := strconv.ParseFloat(temp, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sum += num
	}

	return sum
}
