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

	t := make([]int, 5)
	for _, c := range s {
		if c == 'H' {
			t[0]++
		} else if c == 'I' {
			t[1]++
		} else if c == 'A' {
			t[2]++
		} else if c == 'R' {
			t[3]++
		} else if c == 'C' {
			t[4]++
		}
	}

	min := len(s)
	for i := 0; i < len(t); i++ {
		if min > t[i] {
			min = t[i]
		}
	}
	fmt.Println(min)
}
