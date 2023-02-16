package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscanln(reader, &n, &m)

		total := n * m

		if n > m {
			n, m = m, n
		}

		for m != 0 {
			tmp := n % m
			n = m
			m = tmp
		}

		fmt.Println(total / n)
	}
}
