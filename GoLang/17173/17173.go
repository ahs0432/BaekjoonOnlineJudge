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

	k := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &k[i])
	}

	sum := 0
	t := make([]int, n+1)

	for i := 0; i < m; i++ {
		for j := 1; j <= n; j++ {
			if j%k[i] == 0 {
				t[j] = j
			}
		}
	}

	for i := 1; i <= n; i++ {
		sum += t[i]
	}

	fmt.Println(sum)
}
