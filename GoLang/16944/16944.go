package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	fmt.Fscanln(reader, &s)

	check := []int{1, 1, 1, 1}

	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			check[0] = 0
		} else if c >= 'a' && c <= 'z' {
			check[1] = 0
		} else if c >= '0' && c <= '9' {
			check[2] = 0
		} else if c == '!' || c == '@' || c == '#' || c == '$' || c == '%' || c == '^' || c == '&' || c == '*' || c == '(' || c == ')' || c == '-' || c == '+' {
			check[3] = 0
		}
	}

	sum := 0
	for i := 0; i < 4; i++ {
		sum += check[i]
	}

	n += sum
	if n > 6 {
		fmt.Println(sum)
	} else {
		fmt.Println(sum + 6 - n)
	}
}
