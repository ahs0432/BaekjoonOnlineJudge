package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k, p int
	fmt.Fscanln(reader, &n, &k, &p)

	b := make([]int, n*k)
	for i := 0; i < n*k; i++ {
		fmt.Fscan(reader, &b[i])
	}

	s := 0
	var nc int
	for i := 0; i < n; i++ {
		nc = 0
		for j := i * k; j < (i+1)*k; j++ {
			if b[j] == 0 {
				nc++
			}
		}

		if nc < p {
			s++
		}
	}

	fmt.Println(s)
}
