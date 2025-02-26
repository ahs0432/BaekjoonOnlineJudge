package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m uint64
	fmt.Fscanln(reader, &n)

	i := 1
	for ; i <= 64; i++ {
		m = 1
		for j := 1; j <= i; j++ {
			m *= 2
		}
		m -= 1

		for j := i + 1; j <= 64; j++ {
			m *= 2
		}

		if m == n {
			break
		}
	}
	fmt.Println(i)
}
