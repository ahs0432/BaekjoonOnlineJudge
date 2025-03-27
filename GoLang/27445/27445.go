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

	var x int
	r, c := 1, 1
	maxN, maxM := 50, 50
	for i := 1; i < n; i++ {
		fmt.Fscanln(reader, &x)
		if x < maxN {
			r = i
			maxN = x
		}
	}

	fmt.Fscan(reader, &x)
	if x < maxN {
		r = n
	}

	maxM = x
	for i := 2; i <= m; i++ {
		fmt.Fscan(reader, &x)
		if x < maxM {
			c = i
			maxM = x
		}
	}
	fmt.Println(r, c)
}
