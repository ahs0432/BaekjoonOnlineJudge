package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, n, w int
	fmt.Fscanln(reader, &a, &b, &n, &w)

	g, s := -1, -1
	for i := 1; i < n; i++ {
		if i*a+(n-i)*b == w {
			if g != -1 || s != -1 {
				fmt.Println(-1)
				return
			}
			g, s = i, n-i
		}
	}

	if g == -1 && s == -1 {
		fmt.Println(-1)
		return
	}
	fmt.Println(g, s)
}
