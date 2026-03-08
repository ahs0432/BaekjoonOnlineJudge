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

	s := make([]string, n)
	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &s[i])
		} else {
			fmt.Fscan(reader, &s[i])
		}
	}

	var m, k int
	fmt.Fscan(reader, &m, &k)

	var p, w, x int
	for i := 0; i < m; i++ {
		p = 0
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &x)
			if s[x-1] == "P" {
				p++
			}
		}
		if p == 0 {
			w++
		}
	}

	if w != 0 {
		fmt.Println("W")
	} else {
		fmt.Println("P")
	}
}
