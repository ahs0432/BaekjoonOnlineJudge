package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscanln(reader, &n, &m)

	f := make([]int, 100)
	b := make([]int, 100)
	res := make([]int, 100)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &f[i], &b[i])
		res[i] = f[i]
	}

	for i := m; i > 0; i-- {
		fmt.Fscanln(reader, &k)
		for j := 0; j < n; j++ {
			if res[j] <= k {
				if res[j] == f[j] {
					res[j] = b[j]
				} else {
					res[j] = f[j]
				}
			}
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += res[i]
	}
	fmt.Println(sum)
}
