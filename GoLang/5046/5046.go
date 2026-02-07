package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, b, h, w, p, a int
	fmt.Fscanln(reader, &n, &b, &h, &w)

	minP := 2147483647
	for i := 0; i < h; i++ {
		fmt.Fscanln(reader, &p)

		for j := 0; j < w; j++ {
			if j == w-1 {
				fmt.Fscanln(reader, &a)
			} else {
				fmt.Fscan(reader, &a)
			}

			if a >= n && n*p <= b && minP > n*p {
				minP = n * p
			}
		}
	}

	if minP == 2147483647 {
		fmt.Println("stay home")
	} else {
		fmt.Println(minP)
	}
}
