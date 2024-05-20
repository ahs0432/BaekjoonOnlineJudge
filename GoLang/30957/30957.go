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

	a := "BSA"
	t := make([]int, 3)
	max := 0

	for i := 0; i < n; i++ {
		if s[i] == 'B' {
			t[0]++

			if max < t[0] {
				max = t[0]
			}
		} else if s[i] == 'S' {
			t[1]++

			if max < t[1] {
				max = t[1]
			}
		} else {
			t[2]++

			if max < t[2] {
				max = t[2]
			}
		}
	}

	if t[0] == t[1] && t[1] == t[2] {
		fmt.Println("SCU")
	} else {
		for i := 0; i < 3; i++ {
			if t[i] == max {
				fmt.Print(string(a[i]))
			}
		}
		fmt.Println()
	}
}
