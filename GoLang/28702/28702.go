package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calc(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return fmt.Sprint(n)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s1, s2, s3 string
	fmt.Fscanln(reader, &s1)
	fmt.Fscanln(reader, &s2)
	fmt.Fscanln(reader, &s3)

	var tmp int

	if s3 != "FizzBuzz" && s3 != "Fizz" && s3 != "Buzz" {
		tmp, _ = strconv.Atoi(s3)
		fmt.Println(calc(tmp + 1))
	} else if s2 != "FizzBuzz" && s2 != "Fizz" && s2 != "Buzz" {
		tmp, _ = strconv.Atoi(s2)
		fmt.Println(calc(tmp + 2))
	} else {
		tmp, _ = strconv.Atoi(s1)
		fmt.Println(calc(tmp + 3))
	}
}
