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

	squares := make([]int, n+1)
	squares[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if squares[i] == 0 || squares[i-(j*j)]+1 < squares[i] {
				squares[i] = squares[i-(j*j)] + 1
			}
		}
	}

	fmt.Println(squares[n])
}
