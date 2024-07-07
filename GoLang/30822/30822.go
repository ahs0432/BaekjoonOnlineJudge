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

	table := make([]int, 5)
	for i := 0; i < n; i++ {
		if s[i] == 'u' {
			table[0]++
		} else if s[i] == 'o' {
			table[1]++
		} else if s[i] == 's' {
			table[2]++
		} else if s[i] == 'p' {
			table[3]++
		} else if s[i] == 'c' {
			table[4]++
		}
	}

	min := 1001
	for i := 0; i < 5; i++ {
		if table[i] < min {
			min = table[i]
		}
	}
	fmt.Println(min)
}
