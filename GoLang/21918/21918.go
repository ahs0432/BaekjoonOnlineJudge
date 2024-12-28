package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &s[i])
			break
		}
		fmt.Fscan(reader, &s[i])
	}

	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &a, &b, &c)

		if a == 1 {
			s[b-1] = c
		} else if a == 2 {
			for j := b - 1; j < c; j++ {
				if s[j] == 0 {
					s[j] = 1
				} else {
					s[j] = 0
				}
			}
		} else if a == 3 {
			for j := b - 1; j < c; j++ {
				s[j] = 0
			}
		} else if a == 4 {
			for j := b - 1; j < c; j++ {
				s[j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()
}
